package TelloDrone

import (
	"fmt"
	"runtime/debug"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	// Expecting tello and keyboard packages to be imported from the main code or project dependencies
	// "gobot.io/x/gobot" // Not directly used in this test, but likely needed by tello/keyboard
	"gobot.io/x/gobot/platforms/dji/tello"
	"gobot.io/x/gobot/platforms/keyboard"
)

// Assume 'intensity' is defined in the TelloDrone package scope.
// If it's not, define it here for testing purposes or adjust the tests.
// // TODO: Adjust this value if the actual 'intensity' used in the package is different.
var intensity = 50

// MockTelloDriver is a mock implementation of the tello.Driver for testing.
// It includes all methods used by the handleKeyboardInput function.
// Note: This mock assumes the methods return nil error for simplicity,
// matching the provided reference implementations which don't show error returns being checked.
type MockTelloDriver struct {
	mock.Mock
	// Embedding tello.Driver to satisfy potential interface requirements
	// or implicit expectations, though we primarily rely on the mock methods.
	// This might be unnecessary if tello.Driver is a struct and not an interface
	// being passed around, but it's safer for broader compatibility.
	// If tello.Driver is large or complex, consider mocking only the required interface.
	// tello.Driver // Commented out as Driver is a struct, not interface here. We mock specific methods.
}

// --- Mock Methods Implementation ---

func (m *MockTelloDriver) Left(val int) error {
	args := m.Called(val)
	return args.Error(0)
}

func (m *MockTelloDriver) Right(val int) error {
	args := m.Called(val)
	return args.Error(0)
}

func (m *MockTelloDriver) Up(val int) error {
	args := m.Called(val)
	return args.Error(0)
}

func (m *MockTelloDriver) Down(val int) error {
	args := m.Called(val)
	return args.Error(0)
}

func (m *MockTelloDriver) Forward(val int) error {
	args := m.Called(val)
	return args.Error(0)
}

func (m *MockTelloDriver) Backward(val int) error {
	args := m.Called(val)
	return args.Error(0)
}

func (m *MockTelloDriver) CounterClockwise(val int) error {
	args := m.Called(val)
	return args.Error(0)
}

func (m *MockTelloDriver) Clockwise(val int) error {
	args := m.Called(val)
	return args.Error(0)
}

func (m *MockTelloDriver) Land() error {
	args := m.Called()
	return args.Error(0)
}

func (m *MockTelloDriver) TakeOff() error {
	args := m.Called()
	return args.Error(0)
}

// Halt is needed to satisfy gobot.Device interface if Driver implicitly requires it
// Add other methods from tello.Driver if they are needed to satisfy interfaces
// or if the code under test might call them indirectly.
// For this specific test, only the methods directly called are strictly necessary.
// func (m *MockTelloDriver) Halt() error {
// 	args := m.Called()
// 	return args.Error(0)
// }
// func (m *MockTelloDriver) Name() string { return "mock" }
// func (m *MockTelloDriver) SetName(n string) {}
// func (m *MockTelloDriver) Start() error { return nil }
// func (m *MockTelloDriver) Connection() gobot.Connection { return nil }

// TestHandleKeyboardInput tests the handleKeyboardInput function using table-driven tests.
func TestHandleKeyboardInput(t *testing.T) {
	// Define test cases based on the scenarios
	testCases := []struct {
		name              string
		inputKey          keyboard.Key // The key being pressed
		inputData         interface{}  // Data passed to the handler
		expectedMethod    string       // Method expected to be called on the driver
		expectedIntensity int          // Intensity value for movement methods
		expectReset       bool         // True if resetDronePostion effects are expected
		expectPanic       bool         // True if the handler call is expected to panic
	}{
		{
			name:              "Scenario 1: Handle 'A' Key for Left Movement",
			inputKey:          keyboard.A,
			inputData:         keyboard.KeyEvent{Key: keyboard.A},
			expectedMethod:    "Left",
			expectedIntensity: intensity,
			expectReset:       false,
			expectPanic:       false,
		},
		{
			name:              "Scenario 2: Handle 'D' Key for Right Movement",
			inputKey:          keyboard.D,
			inputData:         keyboard.KeyEvent{Key: keyboard.D},
			expectedMethod:    "Right",
			expectedIntensity: intensity,
			expectReset:       false,
			expectPanic:       false,
		},
		{
			name:              "Scenario 3: Handle 'W' Key for Up Movement",
			inputKey:          keyboard.W,
			inputData:         keyboard.KeyEvent{Key: keyboard.W},
			expectedMethod:    "Up",
			expectedIntensity: intensity,
			expectReset:       false,
			expectPanic:       false,
		},
		{
			name:              "Scenario 4: Handle 'S' Key for Down Movement",
			inputKey:          keyboard.S,
			inputData:         keyboard.KeyEvent{Key: keyboard.S},
			expectedMethod:    "Down",
			expectedIntensity: intensity,
			expectReset:       false,
			expectPanic:       false,
		},
		{
			name:              "Scenario 5: Handle 'U' Key for Forward Movement",
			inputKey:          keyboard.U,
			inputData:         keyboard.KeyEvent{Key: keyboard.U},
			expectedMethod:    "Forward",
			expectedIntensity: intensity,
			expectReset:       false,
			expectPanic:       false,
		},
		{
			name:              "Scenario 6: Handle 'J' Key for Backward Movement",
			inputKey:          keyboard.J,
			inputData:         keyboard.KeyEvent{Key: keyboard.J},
			expectedMethod:    "Backward",
			expectedIntensity: intensity,
			expectReset:       false,
			expectPanic:       false,
		},
		{
			name:              "Scenario 7: Handle 'K' Key for Counter-Clockwise Rotation",
			inputKey:          keyboard.K,
			inputData:         keyboard.KeyEvent{Key: keyboard.K},
			expectedMethod:    "CounterClockwise",
			expectedIntensity: intensity,
			expectReset:       false,
			expectPanic:       false,
		},
		{
			name:              "Scenario 8: Handle 'H' Key for Clockwise Rotation",
			inputKey:          keyboard.H,
			inputData:         keyboard.KeyEvent{Key: keyboard.H},
			expectedMethod:    "Clockwise",
			expectedIntensity: intensity,
			expectReset:       false,
			expectPanic:       false,
		},
		{
			name:           "Scenario 9: Handle 'L' Key for Landing",
			inputKey:       keyboard.L,
			inputData:      keyboard.KeyEvent{Key: keyboard.L},
			expectedMethod: "Land",
			expectReset:    false,
			expectPanic:    false,
		},
		{
			name:           "Scenario 10: Handle 'T' Key for Take Off",
			inputKey:       keyboard.T,
			inputData:      keyboard.KeyEvent{Key: keyboard.T},
			expectedMethod: "TakeOff",
			expectReset:    false,
			expectPanic:    false,
		},
		{
			name:        "Scenario 11: Handle Unrecognized Key (Default Case)",
			inputKey:    keyboard.X, // Using 'X' as an example unrecognized key
			inputData:   keyboard.KeyEvent{Key: keyboard.X},
			expectReset: true, // Expect resetDronePostion effects
			expectPanic: false,
		},
		{
			name:        "Scenario 12: Handle Invalid Input Data Type (Panic Scenario)",
			inputData:   "this is not a KeyEvent", // Incorrect data type
			expectReset: false,
			expectPanic: true,
		},
	}

	// List of all methods that might be called, used for AssertNotCalled checks
	allMethods := []string{
		"Left", "Right", "Up", "Down", "Forward", "Backward",
		"CounterClockwise", "Clockwise", "Land", "TakeOff",
	}
	// Methods called by resetDronePostion
	resetMethods := []string{
		"Forward", "Backward", "Up", "Down", "Left", "Right", "Clockwise",
	}

	for _, tc := range testCases {
		tc := tc // Capture range variable
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel() // Run tests in parallel

			// Arrange
			mockDriver := new(MockTelloDriver)
			handler := handleKeyboardInput(mockDriver) // Get the handler function
			panicked := false

			// Defer panic recovery
			defer func() {
				r := recover()
				if r != nil {
					panicked = true
					if !tc.expectPanic {
						// Log unexpected panic
						t.Errorf("Test panicked unexpectedly: %v\n%s", r, string(debug.Stack()))
					} else {
						// Log expected panic
						t.Logf("Caught expected panic: %v", r)
					}
				} else if tc.expectPanic {
					// Log failure if panic was expected but didn't happen
					t.Errorf("Expected test to panic, but it did not")
				}

				// Assert: Verify mock calls after potential panic recovery
				if !panicked && !tc.expectPanic {
					if tc.expectedMethod != "" {
						// Assert specific method call
						switch tc.expectedMethod {
						case "Land", "TakeOff":
							mockDriver.AssertCalled(t, tc.expectedMethod)
							t.Logf("Successfully asserted call to %s()", tc.expectedMethod)
						default: // Movement methods with intensity
							mockDriver.AssertCalled(t, tc.expectedMethod, tc.expectedIntensity)
							t.Logf("Successfully asserted call to %s(%d)", tc.expectedMethod, tc.expectedIntensity)
						}
						// Assert other methods were NOT called
						for _, method := range allMethods {
							if method != tc.expectedMethod {
								mockDriver.AssertNotCalled(t, method)
							}
						}
						// Also ensure reset methods weren't called if not expected
						if !tc.expectReset {
							for _, method := range resetMethods {
								// Need to check with intensity 0 for reset
								mockDriver.AssertNotCalled(t, method, 0)
							}
						}
						t.Logf("Successfully asserted other methods were not called for %s", tc.expectedMethod)

					} else if tc.expectReset {
						// Assert resetDronePostion effects (calls with intensity 0)
						// Note: resetDronePostion calls Clockwise(0), not CounterClockwise(0)
						mockDriver.AssertCalled(t, "Forward", 0)
						mockDriver.AssertCalled(t, "Backward", 0)
						mockDriver.AssertCalled(t, "Up", 0)
						mockDriver.AssertCalled(t, "Down", 0)
						mockDriver.AssertCalled(t, "Left", 0)
						mockDriver.AssertCalled(t, "Right", 0)
						mockDriver.AssertCalled(t, "Clockwise", 0)
						mockDriver.AssertNotCalled(t, "CounterClockwise") // Ensure CCW wasn't called
						mockDriver.AssertNotCalled(t, "Land")
						mockDriver.AssertNotCalled(t, "TakeOff")
						t.Logf("Successfully asserted calls for resetDronePostion")

					} else {
						// Should not happen based on test cases, but safety check
						t.Logf("No specific method call or reset expected and none asserted.")
					}
				} else if panicked && tc.expectPanic {
					// If panic was expected and occurred, ensure no driver methods were called
					mockDriver.AssertNumberOfCalls(t, "Left", 0)
					mockDriver.AssertNumberOfCalls(t, "Right", 0)
					mockDriver.AssertNumberOfCalls(t, "Up", 0)
					mockDriver.AssertNumberOfCalls(t, "Down", 0)
					mockDriver.AssertNumberOfCalls(t, "Forward", 0)
					mockDriver.AssertNumberOfCalls(t, "Backward", 0)
					mockDriver.AssertNumberOfCalls(t, "CounterClockwise", 0)
					mockDriver.AssertNumberOfCalls(t, "Clockwise", 0)
					mockDriver.AssertNumberOfCalls(t, "Land", 0)
					mockDriver.AssertNumberOfCalls(t, "TakeOff", 0)
					t.Logf("Successfully asserted no driver methods were called during panic scenario.")
				}

				// Final check for mock expectations (optional if using AssertCalled/NotCalled extensively)
				// mockDriver.AssertExpectations(t) // Can be redundant but ensures all preset expectations met
			}()

			// Setup mock expectations *before* Act if return values are needed or strict order matters.
			// Here, we use AssertCalled/AssertNotCalled *after* Act, so setup is minimal.
			// Example setup (if needed):
			// if !tc.expectPanic && tc.expectedMethod != "" {
			// 	switch tc.expectedMethod {
			// 	case "Land", "TakeOff":
			// 		mockDriver.On(tc.expectedMethod).Return(nil)
			// 	default:
			// 		mockDriver.On(tc.expectedMethod, tc.expectedIntensity).Return(nil)
			// 	}
			// }
			// if !tc.expectPanic && tc.expectReset {
			// 	mockDriver.On("Forward", 0).Return(nil)
			// 	// ... setup for other reset methods
			// }

			// Act
			t.Logf("Invoking handler with input data: %#v", tc.inputData)
			handler(tc.inputData) // Execute the handler function

			// Assertions are handled in the deferred function to ensure they run even if Act panics.
			if !tc.expectPanic && !panicked {
				t.Logf("Handler executed successfully without panic.")
			}

		})
	}
}

// Helper function to check if a string is in a slice of strings
// (Not strictly needed with testify's AssertNotCalled, but can be useful)
// func contains(slice []string, item string) bool {
// 	for _, s := range slice {
// 		if s == item {
// 			return true
// 		}
// 	}
// 	return false
// }

