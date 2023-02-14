package handlers

import (
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/khai93/stella"
)

type SubmissionHandler struct {
	SubmissionService stella.SubmissionService
}

// @Description Creates a Submission to SubmissionService and returns the response
// @Param request body stella.SubmissionInput true "Submission Input"
// @Success 201 {object} stella.SubmissionOutput
// @Failure 500 {object} httputil.HttpError
// @Router /submissions/create [post]
func (h SubmissionHandler) CreateSubmission(c *gin.Context) {
	var body stella.SubmissionInput
	bodyErr := c.Bind(&body)
	if bodyErr != nil {
		c.Error(bodyErr)
		return
	}

	if body.LanguageId == 0 || body.LanguageId > len(stella.Languages) {
		c.Error(errors.New("Language id '" + fmt.Sprint(body.LanguageId) + "' does not exist"))
		return
	}

	output, err := h.SubmissionService.CreateSubmission(body)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(201, output)
}

// @Description Get a submission from the SubsmissionService and return the response
// @Param token path string true "Submission Token"
// @Success 200 {object} stella.SubmissionOutput
// @Failure 500 {object} httputil.HttpError
// @Router /submissions/{token} [get]
func (h SubmissionHandler) GetSubmission(c *gin.Context) {
	token := c.Param("token")

	output, err := h.SubmissionService.GetSubmission(token)
	if output == nil {
		c.Status(404)
		return
	}
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, output)
}
