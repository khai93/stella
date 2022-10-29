package handlers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/khai93/stella"
)

type SubmissionHandler struct {
	ExecService stella.ExecutionService
}

// Creates a Submission to ExecutionService and returns the response
func (h SubmissionHandler) CreateSubmission(c *gin.Context) {
	var body stella.SubmissionInput
	bodyErr := c.Bind(&body)
	if bodyErr != nil {
		c.Error(bodyErr)
		return
	}

	output, err := h.ExecService.CreateSubmission(body)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(201, output)
}

// Get a submission from the ExecutionService and return the response
func (h SubmissionHandler) GetSubmission(c *gin.Context) {
	token := c.Param("token")

	b64e, err := strconv.ParseBool(c.Request.URL.Query().Get("base64_encoded"))
	if err != nil {
		c.Error(err)
		return
	}

	fields := []string{}

	output, err := h.ExecService.GetSubmission(token, b64e, fields)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, output)
}

func (h SubmissionHandler) CreateTestSubmission(c *gin.Context) {
}
