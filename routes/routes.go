package routes

import (
	"github.com/docker/docker/client"
	"github.com/gin-gonic/gin"
	"github.com/khai93/stella/config"
	"github.com/khai93/stella/handlers"
	"github.com/khai93/stella/middlewares"
	"github.com/khai93/stella/pkg/docker"
)

type routes struct {
	router *gin.Engine
}

// initializes routes with their handlers and rreturns the routes
func InitRoutes(config config.Configuration) routes {
	r := routes{
		router: gin.Default(),
	}

	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	// Services
	execService := docker.ExecutionService{
		DockerClient: cli,
	}

	v1 := r.router.Group("/v1")
	{
		lh := handlers.LanguageHandler{
			ExecService: execService,
		}

		sh := handlers.SubmissionHandler{
			ExecService: execService,
		}

		v1.GET("/languages", middlewares.ErrorHandler(), lh.GetLanguages)
		v1.POST("/submissions", middlewares.ErrorHandler(), sh.CreateSubmission)
		//v1.GET("/submissions/:token", middlewares.ErrorHandler(), sh.GetSubmission)
		//v1.POST("/test-submissions", middlewares.ErrorHandler(), sh.CreateTestSubmission)
	}

	return r
}

func (r routes) Run(addr ...string) error {
	return r.router.Run(addr...)
}
