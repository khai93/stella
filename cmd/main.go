package main

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/alitto/pond"
	"github.com/docker/docker/client"
	"github.com/go-redis/redis/v9"
	"github.com/khai93/stella"
	"github.com/khai93/stella/config"
	"github.com/khai93/stella/pkg/docker"
	stella_redis "github.com/khai93/stella/pkg/redis"
	"github.com/khai93/stella/routes"
)

// @title Stella API
// @version 1.0
// @Description Code execution API
// @termsOfService https://github.com/khai93/stella/blob/main/LICENSE

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:4000
// @BasePath /v1

func main() {
	c, err := config.Get()
	if err != nil {
		panic(err)
	}

	// Docker
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	// Redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     c.Redis.Address,
		Password: c.Redis.Password,
		DB:       c.Redis.DB,
	})

	// Services
	execService := docker.ExecutionService{
		DockerClient: cli,
	}

	subService := stella_redis.SubmissionService{
		Client: rdb,
	}

	// Start workers
	go func() {
		pond := pond.New(c.Workers, 0)
		ctx := context.Background()

		sub := rdb.Subscribe(ctx, "submissions")
		defer sub.Close()
		for {
			msg, err := sub.ReceiveMessage(ctx)
			if err != nil {
				panic(err)
			}

			fmt.Println(msg)

			pond.Submit(func() {
				job := stella.SubmissionInput{}
				json.Unmarshal([]byte(msg.Payload), &job)
				output, err := execService.ExecuteSubmission(job)
				if err != nil {
					panic(err)
				}

				setErr := rdb.Set(ctx, "submission:"+output.Token, output, 0).Err()
				if setErr != nil {
					panic(setErr)
				}
			})
		}
	}()

	r := routes.InitRoutes(*c, execService, subService)

	r.Run(":" + strconv.Itoa(c.Port))
}
