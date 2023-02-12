package test

type Test struct {
	Summary string `json:"summary"`
	HttpVerb             string      `json:"http_verb"`
	EndpointPath         string      `json:"endpoint_path"`
	RequestBody          interface{} `json:"request_body"`
	ExpectedResponseBody interface{} `json:"expected_response_body"`
	ExpectedHttpStatus   int         `json:"expected_status_code"`
	SetUp string `json:"set_up"`
}

type TestResult struct {
	Passed bool
	FailedReason string
}
