package reader

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/georgexx009/eagle/test"
)

const DTO_PATH = "./DTOs/"
const REQUESTS = "requests"
const RESPONSES = "responses"

type tests struct {
	Tests []test.Test `json:"tests"`
}

func ReadDTO() ([]test.Test, error) {
	testsPath := fmt.Sprintf("./eagle-tests/tests/tests.json")
	byteArr, err := readFile(testsPath)
	if err != nil {
		return nil, err
	}

	var dtoTests tests

	err = json.Unmarshal(byteArr, &dtoTests)
	if err != nil {
		return nil, err
	}

	return dtoTests.Tests, nil
}

func readFile(path string) ([]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	return io.ReadAll(file)
}
