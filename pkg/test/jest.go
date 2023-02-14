package test

type JestOutput struct {
	TestResults []JestOutputTestResult `json:"testResults"`
}

type JestOutputTestResult struct {
	AssertionResults []JestOutputAssertionResult `json:"assertionResults"`
	EndTime          int64                       `json:"endTime"`
	StartTime        int64                       `json:"startTime"`
	Name             string                      `json:"name"`
	Status           string                      `json:"status"`
}

type JestOutputAssertionResult struct {
	Status          string   `json:"status"`
	Title           string   `json:"title"`
	Duration        int      `json:"duration"`
	FullName        string   `json:"fullName"`
	FailureMessages []string `json:"failureMessages"`
}
