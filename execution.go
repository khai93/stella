package stella

// ExecutionService handles execution of submissions
type ExecutionService interface {
	ExecuteSubmission(input SubmissionInput) (*SubmissionOutput, error)
}
