package cmd

import (
	fmt "fmt"
	io "io"
	os "os"
	strings "strings"
	sync "sync"
	testing "testing"
	bytes "bytes"
	filepath "path/filepath"
	debug "runtime/debug"
	viper "github.com/spf13/viper"
)



var exitCalled bool
var exitCode int = -1
var exitMu sync.Mutex
var osExit = os.Exit




/*
ROOST_METHOD_HASH=er_6b05c3a223
ROOST_METHOD_SIG_HASH=er_6b05c3a223

FUNCTION_DEF=func er(msg interface) 

*/
func TestEr(t *testing.T) {

	title := "Test Title"

	testStruct := &createBlogCmdParams{Title: &title}

	expectedStructOutput := fmt.Sprintf("Error: %v\n", testStruct)

	testCases := []struct {
		name         string
		input        interface{}
		expectedOut  string
		expectedCode int
	}{
		{
			name:         "Scenario 1: Basic String Message",
			input:        "Something went wrong",
			expectedOut:  "Error: Something went wrong\n",
			expectedCode: 1,
		},
		{
			name:         "Scenario 2: Error Type Message",
			input:        fmt.Errorf("database connection failed: %w", io.ErrUnexpectedEOF),
			expectedOut:  "Error: database connection failed: unexpected EOF\n",
			expectedCode: 1,
		},
		{
			name:         "Scenario 3: Nil Input",
			input:        nil,
			expectedOut:  "Error: <nil>\n",
			expectedCode: 1,
		},
		{
			name:         "Scenario 4: Integer Input",
			input:        500,
			expectedOut:  "Error: 500\n",
			expectedCode: 1,
		},
		{
			name:         "Scenario 5: Struct Pointer Input",
			input:        testStruct,
			expectedOut:  expectedStructOutput,
			expectedCode: 1,
		},
		{
			name:         "Scenario 6: Empty String Input",
			input:        "",
			expectedOut:  "Error: \n",
			expectedCode: 1,
		},
	}

	originalStdout := os.Stdout
	originalOsExit := osExit

	t.Cleanup(func() {
		os.Stdout = originalStdout
		osExit = originalOsExit

		exitMu.Lock()
		exitCode = -1
		exitCalled = false
		exitMu.Unlock()
	})

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Logf("Running test case: %s", tc.name)
			t.Logf("Input: %#v", tc.input)

			exitMu.Lock()
			exitCode = -1
			exitCalled = false
			exitMu.Unlock()
			osExit = mockExit

			r, w, pipeErr := os.Pipe()
			if pipeErr != nil {
				t.Fatalf("os.Pipe failed: %v", pipeErr)
			}
			os.Stdout = w

			var recoveredPanic interface{}
			var capturedOutput string

			func() {

				defer func() {

					os.Stdout = originalStdout

					w.Close()

					outBytes, err := io.ReadAll(r)
					if err != nil {
						t.Errorf("Failed to read captured output: %v", err)
					}
					capturedOutput = string(outBytes)

					recoveredPanic = recover()
				}()

				er(tc.input)

			}()

			expectedPanicMsg := fmt.Sprintf("mock os.Exit called with code %d", tc.expectedCode)
			if recoveredPanic == nil {
				t.Errorf("Expected panic with message '%s', but function did not panic", expectedPanicMsg)
			} else {
				panicMsg, ok := recoveredPanic.(string)
				if !ok || panicMsg != expectedPanicMsg {
					t.Errorf("Expected panic with message '%s', got %T: %v", expectedPanicMsg, recoveredPanic, recoveredPanic)

				}
			}

			exitMu.Lock()
			wasCalled := exitCalled
			finalCode := exitCode
			exitMu.Unlock()

			if !wasCalled {
				t.Errorf("Expected os.Exit (mocked) to be called, but it wasn't")
			}

			if finalCode != tc.expectedCode {
				t.Errorf("Expected exit code %d, got %d", tc.expectedCode, finalCode)
			}

			if capturedOutput != tc.expectedOut {
				t.Errorf("Expected output:\n%q\nGot:\n%q", tc.expectedOut, capturedOutput)
			}

			t.Logf("Captured output:\n%s", capturedOutput)
			t.Logf("Captured exit code: %d", finalCode)

		})
	}
}

func compareMultilineString(t *testing.T, expected, actual string) {
	t.Helper()
	if strings.TrimSpace(expected) != strings.TrimSpace(actual) {
		t.Errorf("Mismatch (-expected +got):\n-%s\n+%s", expected, actual)
	}
}

func mockExit(code int) {
	exitMu.Lock()
	exitCode = code
	exitCalled = true
	exitMu.Unlock()

	panic(fmt.Sprintf("mock os.Exit called with code %d", code))
}

func safeString(v interface{}) string {
	if v == nil {
		return "<nil>"
	}

	return fmt.Sprintf("%v", v)
}


/*
ROOST_METHOD_HASH=initConfig_b4ae76b127
ROOST_METHOD_SIG_HASH=initConfig_25f2d0dcb4

FUNCTION_DEF=func initConfig() 

*/
func TestInitConfig(t *testing.T) {

	type testCase struct {
		name   string
		setup  func(t *testing.T) (cleanup func())
		assert func(t *testing.T, output string)

	}

