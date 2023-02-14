package test

import (
	"encoding/json"
	"errors"
	"strings"

	"github.com/khai93/stella"
)

type TestService struct{}

func (t TestService) ParseTestOutput(stdout string, framework stella.TestFramework) (*stella.TestParseOutput, error) {
	passedTests := []stella.Test{}
	failedTests := []stella.Test{}

	switch framework {
	case stella.JestTestFramework:
		var parsed JestOutput
		json.Unmarshal([]byte(stdout), &parsed)

		for _, t := range parsed.TestResults {
			for _, a := range t.AssertionResults {
				if a.Status == "passed" {
					passedTests = append(passedTests, stella.Test{
						Description: a.Title,
						Passed:      true,
						RawOutput:   a.FullName,
						Duration:    a.Duration,
					})
				} else if a.Status == "failed" {
					failedTests = append(passedTests, stella.Test{
						Description: a.Title,
						Passed:      false,
						RawOutput:   a.FailureMessages[0],
						Duration:    a.Duration,
					})
				}
			}
		}
	case stella.GoTestFramework:
		lines := strings.Split(stdout, "\n")

		for i := 0; i < len(lines); i++ {
			var parsed GoTestOutput
			json.Unmarshal([]byte(lines[i]), &parsed)
			if parsed.Action == "pass" {
				passedTests = append(passedTests, stella.Test{
					Description: parsed.Test,
					Passed:      true,
					RawOutput:   parsed.Package,
					Duration:    int(parsed.Elapsed * 1000),
				})
			} else if parsed.Action == "fail" {
				passedTests = append(passedTests, stella.Test{
					Description: parsed.Test,
					Passed:      true,
					RawOutput:   parsed.Package,
					Duration:    int(parsed.Elapsed * 1000),
				})
			}
		}
	default:
		return nil, errors.New("unknown testing framework provided")
	}

	return &stella.TestParseOutput{
		Passed: passedTests,
		Failed: failedTests,
	}, nil
}
