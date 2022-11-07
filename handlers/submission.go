package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/khai93/stella"
)

type SubmissionHandler struct {
	SubmissionService stella.SubmissionService
}

// Creates a Submission to SubmissionService and returns the response
func (h SubmissionHandler) CreateSubmission(c *gin.Context) {
	var body stella.SubmissionInput
	bodyErr := c.Bind(&body)
	if bodyErr != nil {
		c.Error(bodyErr)
		return
	}

	output, err := h.SubmissionService.CreateSubmission(body)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(201, output)
}

// Get a submission from the SubsmissionService and return the response
func (h SubmissionHandler) GetSubmission(c *gin.Context) {
	token := c.Param("token")

	output, err := h.SubmissionService.GetSubmission(token)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, output)
}
