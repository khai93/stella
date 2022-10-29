package docker

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"os"
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

func (j ExecutionService) CreateSubmission(input stella.SubmissionInput) (*stella.SubmissionOutput, error) {
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

	reader, err := j.DockerClient.ImagePull(ctx, "docker.io/library/"+langauge.Image, types.ImagePullOptions{})
	if err != nil {
		panic(err)
	}

	defer reader.Close()
	io.Copy(os.Stdout, reader)

	resp, err := j.DockerClient.ContainerCreate(ctx, &container.Config{
		Image:           langauge.Image,
		Cmd:             strings.Split(langauge.Cmd, " "),
		Tty:             false,
		AttachStdout:    true,
		NetworkDisabled: true,
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
		panic(err)
	}

	if err := j.DockerClient.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		panic(err)
	}

	statusCh, errCh := j.DockerClient.ContainerWait(ctx, resp.ID, container.WaitConditionNotRunning)
	select {
	case err := <-errCh:
		if err != nil {
			panic(err)
		}
	case <-statusCh:
	}

	reader, err2 := j.DockerClient.ContainerLogs(ctx, resp.ID, types.ContainerLogsOptions{ShowStdout: true, ShowStderr: true, Timestamps: false, Follow: true})
	if err2 != nil {
		panic(err)
	}
	defer reader.Close()

	stdoutput := &bytes.Buffer{}
	stderror := &bytes.Buffer{}
	data, err := j.DockerClient.ContainerInspect(ctx, resp.ID)
	if err != nil {
		panic(err)
	}

	startedAt, err := time.Parse(time.RFC3339Nano, data.State.StartedAt)
	if err != nil {
		panic(err)
	}

	endedAt, err := time.Parse(time.RFC3339Nano, data.State.FinishedAt)
	if err != nil {
		panic(err)
	}

	stdcopy.StdCopy(stdoutput, stderror, reader)

	var output stella.SubmissionOutput = stella.SubmissionOutput{
		Stdout:   stdoutput.String(),
		Stderr:   stderror.String(),
		Executed: true,
		ExitCode: data.State.ExitCode,
		Memory:   float32(data.Node.Memory),
		Time:     float32(endedAt.Sub(startedAt).Seconds()),
	}

	return &output, nil
}

// TODO
func (j ExecutionService) CreateTestSubmission(input stella.TestSubmissionInput, base64_encoded bool, wait bool) (*stella.SubmissionOutput, error) {
	return nil, errors.New("Not Implemented")
}

func (j ExecutionService) GetSubmission(token string, base64_encoded bool, fields []string) (*stella.SubmissionOutput, error) {
	return nil, errors.New("Not Implemented")
}

func (j ExecutionService) GetLanguages() ([]stella.SubmissionLanguage, error) {
	return nil, errors.New("Not Implemented")
}
