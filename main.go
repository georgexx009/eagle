package main

import (
	"github.com/georgexx009/eagle/pkg/logger"
	"github.com/georgexx009/eagle/pkg/reader"
)

func main() {
	logger.Log("Running tests...")
	logger.Log("Reading tests...")
	tests, err := reader.ReadDTO()
	if err != nil {
		logger.Log(err.Error())
	}
	logger.Log(tests)
}
