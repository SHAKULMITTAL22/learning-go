```go
package cmd

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime/debug"
	"strings"
	"testing"

	// Assuming homedir is available from the package context as per instructions.
	// "github.com/mitchellh/go-homedir"

	// Assuming cobra is available if needed, though not directly used in initConfig.
	// "github.com/spf13/cobra"

	"github.com/spf13/viper"
)

// Assume cfgFile is a global variable in the cmd package being tested.
// It is expected to be declared in the actual cmd package code.
// var cfgFile string

// Assume er is a function in the cmd package that handles errors, potentially by exiting.
// It is expected to be declared in the actual cmd package code.
// func er(msg interface{}) { ... }

// Helper function to capture standard output
func captureOutput(f func()) string {
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f() // Execute the function whose output we want to capture

	w.Close()
	os.Stdout = oldStdout // Restore stdout
	var buf bytes.Buffer
	io.Copy(&buf, r)
	return buf.String()
}

// TestInitConfig provides unit tests for the initConfig function.
func TestInitConfig(t *testing.T) {
	// --- Test Data Structures --- (Defined within the test function)
	type testCase struct {
		name   string
		setup  func(t *testing.T) (cleanup func()) // Returns a cleanup function
		assert func(t *testing.T, output string)
		// expectPanic is not used as the critical error path calls os.Exit via er()
	}

	// --- Test Cases ---
	testCases := []testCase{
		// Scenario 1: Configuration loaded successfully using cfgFile
		{
			name: "Scenario 1: Success with cfgFile",
			setup: func(t *testing.T) func() {
				viper.Reset() // Ensure viper is clean
				// Create a temporary config file
				tmpDir := t.TempDir()
				tmpFile, err := os.CreateTemp(tmpDir, "test_config_*.yaml")
				if err != nil {
					t.Fatalf("Scenario 1 Setup: Failed to create temp config file: %v", err)
				}
				configContent := "setting1: value1\notherkey: value2"
				if _, err := tmpFile.WriteString(configContent); err != nil {
					tmpFile.Close()
					t.Fatalf("Scenario 1 Setup: Failed to write to temp config file: %v", err)
				}
				tmpFile.Close()
				cfgFile = tmpFile.Name() // Set the global cfgFile variable for initConfig

				// Return cleanup function
				return func() {
					cfgFile = "" // Clean up global state
					viper.Reset()
					// Temp file and dir are cleaned up automatically by t.TempDir()
				}
			},
			assert: func(t *testing.T, output string) {
				t.Log("Scenario 1: Verifying config loaded from specified cfgFile")
				expectedConfigFile := cfgFile // Capture the value used in setup
				if viper