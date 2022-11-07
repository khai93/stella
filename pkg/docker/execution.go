package docker

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/stdcopy"
	"github.com/khai93/stella"
	"github.com/khai93/stella/config"
	filelib "github.com/khai93/stella/lib/file"
	"golang.org/x/exp/slices"
)

type ExecutionService struct {
	DockerClient *client.Client
}

func (j ExecutionService) ExecuteSubmission(input stella.SubmissionInput) (*stella.SubmissionOutput, error) {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}
	j.DockerClient = cli

	config, err := config.Get()
	if err != nil {
		return nil, err
	}

	languageIndex := slices.IndexFunc(stella.Langauges, func(l stella.Language) bool { return l.Id == input.LanguageId })
	if languageIndex == -1 {
		return nil, errors.New("Language Id '" + fmt.Sprint(input.LanguageId) + "' does not exist.")
	}
	langauge := stella.Langauges[languageIndex]

	file, err := filelib.CreateFileBuffer(input.SourceCode, langauge.EntryFileName)
	if err != nil {
		return nil, err
	}

	ctx := context.Background()

	// inst
	containerTimeout := 0

	resp, err := j.DockerClient.ContainerCreate(ctx, &container.Config{
		Image:           "stella-compilers",
		Cmd:             strings.Split(langauge.Cmd, " "),
		Tty:             false,
		OpenStdin:       true,
		AttachStdin:     true,
		AttachStdout:    true,
		NetworkDisabled: true,
		StopTimeout:     &containerTimeout,
	},
		&container.HostConfig{
			Resources: container.Resources{
				Memory: int64(config.MemoryLimits),
			},
		}, nil, nil, "")
	if err != nil {
		return nil, err
	}

	defer j.DockerClient.ContainerRemove(ctx, resp.ID, types.ContainerRemoveOptions{})

	if err := j.DockerClient.CopyToContainer(ctx, resp.ID, "/", file, types.CopyToContainerOptions{AllowOverwriteDirWithFile: true}); err != nil {
		return nil, err
	}

	if err := j.DockerClient.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		return nil, err
	}

	waiter, err := cli.ContainerAttach(ctx, resp.ID, types.ContainerAttachOptions{
		Stdin:  true,
		Stream: true,
	})
	if err != nil {
		return nil, err
	}

	// Write StdIn
	_, writeErr := waiter.Conn.Write([]byte(input.StdIn))
	if writeErr != nil {
		return nil, writeErr
	}
	statusCh, errCh := j.DockerClient.ContainerWait(ctx, resp.ID, container.WaitConditionNotRunning)
	select {
	case err := <-errCh:
		if err != nil {
			return nil, err
		}
	case <-time.After(time.Duration(config.Timeout * int(time.Second))):
		exitErr := j.DockerClient.ContainerStop(ctx, resp.ID, nil)
		if exitErr != nil {
			return nil, exitErr
		}
		exitOutput := stella.SubmissionOutput{
			Executed: true,
			ExitCode: 124,
			Token:    input.Token,
			Time:     float32(config.Timeout),
		}

		return &exitOutput, nil
	case <-statusCh:
	}

	reader, err2 := j.DockerClient.ContainerLogs(ctx, resp.ID, types.ContainerLogsOptions{ShowStdout: true, ShowStderr: true, Timestamps: false, Follow: true})
	if err2 != nil {
		return nil, err
	}
	defer reader.Close()

	stdoutput := &bytes.Buffer{}
	stderror := &bytes.Buffer{}
	data, err := j.DockerClient.ContainerInspect(ctx, resp.ID)
	if err != nil {
		return nil, err
	}

	startedAt, err := time.Parse(time.RFC3339Nano, data.State.StartedAt)
	if err != nil {
		return nil, err
	}

	endedAt, err := time.Parse(time.RFC3339Nano, data.State.FinishedAt)
	if err != nil {
		return nil, err
	}

	stdcopy.StdCopy(stdoutput, stderror, reader)

	var output stella.SubmissionOutput = stella.SubmissionOutput{
		Stdout:   stdoutput.String(),
		Stderr:   stderror.String(),
		Executed: true,
		ExitCode: data.State.ExitCode,
		Token:    input.Token,
		Time:     float32(endedAt.Sub(startedAt).Seconds()),
	}

	return &output, nil
}

func (j ExecutionService) GetSubmission(token string, base64_encoded bool, fields []string) (*stella.SubmissionOutput, error) {
	return nil, errors.New("Not Implemented")
}

func (j ExecutionService) GetLanguages() ([]stella.SubmissionLanguage, error) {
	return nil, errors.New("Not Implemented")
}
