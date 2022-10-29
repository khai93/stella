package main

import (
	"archive/tar"
	"bytes"
	"context"
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/stdcopy"
	v1 "github.com/opencontainers/image-spec/specs-go/v1"
)

func s() {
	ctx := context.Background()

	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	// PULL AN IMAGE BASED ON THE LANGUAGE
	reader, err := cli.ImagePull(ctx, "docker.io/library/golang", types.ImagePullOptions{})
	if err != nil {
		panic(err)
	}

	defer reader.Close()
	io.Copy(os.Stdout, reader)

	content := "package main;\n\nimport \"fmt\";\n\nfunc FibonacciRecursion(n int) int {\n    if n <= 1 {\n        return n;\n    }\n    return FibonacciRecursion(n-1) + FibonacciRecursion(n-2);\n}\n\nfunc main() {\n\tfmt.Println(FibonacciRecursion(50));\n}\n"
	var buf bytes.Buffer
	tw := tar.NewWriter(&buf)
	err = tw.WriteHeader(&tar.Header{
		Name: "main.go",           // filename
		Mode: 0777,                // permissions
		Size: int64(len(content)), // filesize
	})
	if err != nil {
		panic(err)
	}
	tw.Write([]byte(content))
	tw.Close()

	resp, err := cli.ContainerCreate(ctx, &container.Config{
		Image:           "golang",
		Cmd:             []string{"go", "run", "/main.go"},
		Tty:             false,
		AttachStdout:    true,
		NetworkDisabled: true,
	}, &container.HostConfig{
		Resources: container.Resources{
			Memory: 1e+7,
		},
	}, &network.NetworkingConfig{}, &v1.Platform{}, "")
	if err != nil {
		panic(err)
	}

	if err := cli.CopyToContainer(context.Background(), resp.ID, "/", &buf, types.CopyToContainerOptions{}); err != nil {
		panic(err)
	}

	if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		panic(err)
	}

	statusCh, errCh := cli.ContainerWait(ctx, resp.ID, container.WaitConditionNotRunning)
	select {
	case err := <-errCh:
		if err != nil {
			panic(err)
		}
	case <-statusCh:
	}

	reader, err2 := cli.ContainerLogs(ctx, resp.ID, types.ContainerLogsOptions{ShowStdout: true, ShowStderr: true, Timestamps: true, Follow: true})
	if err2 != nil {
		panic(err)
	}
	defer reader.Close()

	stdoutput := &bytes.Buffer{}
	stderror := &bytes.Buffer{}
	data, err := cli.ContainerInspect(ctx, resp.ID)
	if err != nil {
		panic(err)
	}

	startedAt, err := time.Parse(time.RFC3339Nano, data.State.StartedAt)
	if err != nil {
		panic(err)
	}

	stdcopy.StdCopy(stdoutput, stderror, reader)
	split := strings.Split(stdoutput.String(), " ")
	if len(strings.Trim(stdoutput.String(), " ")) > 0 {
		date, err := time.Parse(time.RFC3339Nano, split[0])
		if err != nil {
			panic(err)
		}
		fmt.Println(stdoutput.String())
		fmt.Printf("Time Elapsed: %vs", date.Sub(startedAt).Seconds())
	}

	fmt.Println(stderror.String())

	defer cli.ContainerRemove(ctx, resp.ID, types.ContainerRemoveOptions{})
}
