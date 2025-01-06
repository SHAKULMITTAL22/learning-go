package cmd

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"testing"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"your/full/path/to/cmd"
)

type T struct {
	common
	isEnvSet bool
	context  *testContext // For running tests and subtests.
}
func Tester(t *testing.T) {

	testCases := []struct {
		name          string
		errorMessage  interface{}
		expectedError string
	}{
		{
			name:          "Normal Operation of the Error Function",
			errorMessage:  "Normal error message",
			expectedError: "Error: Normal error message",
		},
		{
			name:          "Passing an Empty Error Message to the Function",
			errorMessage:  "",
			expectedError: "Error: ",
		},
		{
			name:          "Passing a Non-String Error Message to the Function",
			errorMessage:  123,
			expectedError: "Error: 123",
		},
	}

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in Tester", r)
		}
	}()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			var buf bytes.Buffer
			fmt.SetOutput(&buf)

			er(tc.errorMessage)

			if gotError := strings.TrimSpace(buf.String()); gotError != tc.expectedError {
				t.Errorf("Er() = %v, want %v", gotError, tc.expectedError)
			}
		})
	}
}

