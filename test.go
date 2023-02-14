package stella

type TestFramework string

const (
	JestTestFramework TestFramework = "jest"
	GoTestFramework   TestFramework = "gotest"
)

type Test struct {
	Description string `json:"description"`
	Passed      bool   `json:"passed"`
	RawOutput   string `json:"raw_output"`
	Duration    int    `json:"duration"` // In Milliseconds
}

type TestParseOutput struct {
	Passed []Test
	Failed []Test
}

// handles parsing of test framework's json output
type TestService interface {
	ParseTestOutput(stdout string, framework TestFramework) (*TestParseOutput, error)
}
