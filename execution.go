package stella

type SubmissionInput struct {
	SourceCode      string `json:"source_code,omitempty"`
	LanguageId      int    `json:"language_id,omitempty"`
	AdditionalFiles string `json:"additional_files,omitempty"`
}

type TestSubmissionInput struct {
	SourceCode      string `json:"source_code,omitempty"`
	TestSourceCode  string `json:"test_source_code,omitempty"`
	LanguageId      int    `json:"language_id,omitempty"`
	AdditionalFiles string `json:"additional_files,omitempty"`
	FrameworkId     int    `json:"framework_id,omitempty"`
}

type SubmissionOutput struct {
	Stdout   string  `json:"stdout,omitempty"`
	Stderr   string  `json:"stderr,omitempty"`
	ExitCode int     `json:"exit_code,omitempty"`
	Token    string  `json:"token,omitempty"`
	Memory   float32 `json:"memory,omitempty"`
	Executed bool    `json:"executed,omitempty"`
	Time     float32 `json:"time,string,omitempty"`
}

type SubmissionLanguage struct {
	Id   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type ExecutionService interface {
	CreateSubmission(input SubmissionInput) (*SubmissionOutput, error)
	CreateTestSubmission(input TestSubmissionInput, base64_encoded bool, wait bool) (*SubmissionOutput, error)
	GetSubmission(token string, base64_encoded bool, fields []string) (*SubmissionOutput, error)
	GetLanguages() ([]SubmissionLanguage, error)
}
