package stella

import "encoding/json"

type SubmissionInput struct {
	Token           string `json:"token,omitempty"`
	SourceCode      string `json:"source_code,omitempty"`
	LanguageId      int    `json:"language_id,omitempty"`
	AdditionalFiles string `json:"additional_files,omitempty"`
	ExpectedOutput  string `json:"expected_output,omitempty"`
	StdIn           string `json:"std_in,omitempty"`
}

type TestSubmissionInput struct {
	SourceCode      string `json:"source_code,omitempty"`
	TestSourceCode  string `json:"test_source_code,omitempty"`
	LanguageId      int    `json:"language_id,omitempty"`
	AdditionalFiles string `json:"additional_files,omitempty"`
	FrameworkId     int    `json:"framework_id,omitempty"`
}

type SubmissionOutput struct {
	Stdout        string  `json:"stdout,omitempty"`
	Stderr        string  `json:"stderr,omitempty"`
	ExitCode      int     `json:"exit_code,omitempty"`
	Token         string  `json:"token"`
	Memory        float32 `json:"memory,omitempty"`
	Executed      bool    `json:"executed"`
	OutputMatched bool    `json:"output_matched"`
	Time          float32 `json:"time,string,omitempty"`
}

func (s SubmissionOutput) MarshalBinary() ([]byte, error) {
	return json.Marshal(s)
}

func (s SubmissionInput) MarshalBinary() ([]byte, error) {
	return json.Marshal(s)
}

type SubmissionLanguage struct {
	Id   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type SubmissionService interface {
	CreateSubmission(input SubmissionInput) (*SubmissionOutput, error)
	GetSubmission(token string) (*SubmissionOutput, error)
}
