package test

import (
	"encoding/json"
	"errors"

	"github.com/khai93/stella"
)

type TestService struct{}

func (t TestService) ParseTestOutput(stdout string, framework stella.TestFramework) (*stella.TestParseOutput, error) {
	var passedTests []stella.Test
	var failedTests []stella.Test

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
				} else {
					failedTests = append(passedTests, stella.Test{
						Description: a.Title,
						Passed:      false,
						RawOutput:   a.FailureMessages[0],
						Duration:    a.Duration,
					})
				}
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
