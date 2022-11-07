package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/khai93/stella"
	"github.com/khai93/stella/config"
	"github.com/khai93/stella/handlers"
	"github.com/khai93/stella/middlewares"
)

type routes struct {
	router *gin.Engine
}

// initializes routes with their handlers and rreturns the routes
func InitRoutes(config config.Configuration, execService stella.ExecutionService, subService stella.SubmissionService) routes {
	r := routes{
		router: gin.Default(),
	}

	v1 := r.router.Group("/v1")
	{
		lh := handlers.LanguageHandler{
			ExecService: execService,
		}

		sh := handlers.SubmissionHandler{
			SubmissionService: subService,
		}

		v1.GET("/languages", middlewares.ErrorHandler(), lh.GetLanguages)
		v1.POST("/submissions/create", middlewares.ErrorHandler(), sh.CreateSubmission)
		v1.GET("/submissions/:token", middlewares.ErrorHandler(), sh.GetSubmission)
		//v1.POST("/test-submissions", middlewares.ErrorHandler(), sh.CreateTestSubmission)
	}

	return r
}

func (r routes) Run(addr ...string) error {
	return r.router.Run(addr...)
}
