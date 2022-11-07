package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/khai93/stella"
)

type LanguageHandler struct {
	ExecService stella.ExecutionService
}

/*
Gets the languages from the Execution Service and sends it as response
*/
func (h LanguageHandler) GetLanguages(c *gin.Context) {
	languages := "todo"

	c.JSON(200, languages)
}
