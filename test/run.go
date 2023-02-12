package test

import (
	"encoding/json"
	"fmt"
	"io"
	"reflect"

	"github.com/wizeline/integration-tests-wizepace-api/pkg/apiclient"
	"github.com/wizeline/integration-tests-wizepace-api/tests/setup"
)

func (test *Test) Run() (TestResult, error) {
	setUpRunner := setup.GetSetUp(test.SetUp)
	if setUpRunner == nil {
		return TestResult{
			Passed:       false,
			FailedReason: "set up does not exists",
		}, nil
	}

	err := setUpRunner.Run()
	if err != nil {
		return TestResult{
			Passed:       false,
			FailedReason: "set up error",
		}, err
	}

	resp, err := apiclient.ConsumeApi(test.EndpointPath, test.HttpVerb, test.RequestBody)
	if err != nil {
		return TestResult{
			Passed:       false,
			FailedReason: "error",
		}, err
	}

	// check http status
	if resp.StatusCode != test.ExpectedHttpStatus {
		reason := fmt.Sprintf("status code do not match, received: %d, expected: %d", resp.StatusCode, test.ExpectedHttpStatus)
		return TestResult{
			Passed:       false,
			FailedReason: reason,
		}, nil
	}

	// check response body
	defer resp.Body.Close()
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return TestResult{
			Passed:       false,
			FailedReason: "error",
		}, err
	}

	if len(bodyBytes) == 0 && test.ExpectedResponseBody != nil {
		return TestResult{
			Passed:       false,
			FailedReason: "expected a response body but it is empty",
		}, nil
	}

	if test.ExpectedResponseBody != nil {
		var body interface{}
		if err := json.Unmarshal(bodyBytes, &body); err != nil {
			return TestResult{
				Passed:       false,
				FailedReason: "error",
			}, err
		}

		if !reflect.DeepEqual(test.ExpectedResponseBody, body) {
			return TestResult{
				Passed:       false,
				FailedReason: "response body do not match",
			}, nil
		}
	}

	err = setUpRunner.CleanUpDb()
	if err != nil {
		return TestResult{
			Passed:       false,
			FailedReason: "set up clean up error",
		}, err
	}

	return TestResult{
		Passed: true,
	}, nil
}
