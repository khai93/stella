package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/khai93/stella"
)

type LanguageHandler struct {
	ExecService stella.ExecutionService
}

type LanguageView struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Version string `json:"version"`
}

// @Description Gets the languages from the Execution Service and sends it as response
// @Success 200 {object} handlers.LanguageView
// @Failure 500 {object} httputil.HttpError
// @Router /languages [get]
func (h LanguageHandler) GetLanguages(c *gin.Context) {
	var output = []LanguageView{}

	for _, l := range stella.Langauges {
		var base = LanguageView{
			Id:      l.Id,
			Name:    l.Name,
			Version: l.Version,
		}
		output = append(output, base)
	}

	c.JSON(200, output)
}
