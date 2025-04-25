package TelloDrone

import (
	fmt "fmt"
	debug "runtime/debug"
	testing "testing"
	assert "github.com/stretchr/testify/assert"
	mock "github.com/stretchr/testify/mock"
	tello "gobot.io/x/gobot/platforms/dji/tello"
	keyboard "gobot.io/x/gobot/platforms/keyboard"
)



var intensity = 50

type MockTelloDriver struct {
	mock.Mock
}


/*
ROOST_METHOD_HASH=handleKeyboardInput_fa50d499f0
ROOST_METHOD_SIG_HASH=handleKeyboardInput_fa50d499f0

FUNCTION_DEF=func handleKeyboardInput(drone *tello.Driver) func(interface) 

*/
func (m *MockTelloDriver) Backward(val int) error {
	args := m.Called(val)
	return args.Error(0)
}

func (m *MockTelloDriver) Clockwise(val int) error {
	args := m.Called(val)
	return args.Error(0)
}

func (m *MockTelloDriver) CounterClockwise(val int) error {
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

func (m *MockTelloDriver) Land() error {
	args := m.Called()
	return args.Error(0)
}

func (m *MockTelloDriver) Left(val int) error {
	args := m.Called(val)
	return args.Error(0)
}

func (m *MockTelloDriver) Right(val int) error {
	args := m.Called(val)
	return args.Error(0)
}

func (m *MockTelloDriver) TakeOff() error {
	args := m.Called()
	return args.Error(0)
}

func (m *MockTelloDriver) Up(val int) error {
	args := m.Called(val)
	return args.Error(0)
}

func TestHandleKeyboardInput(t *testing.T) {

	testCases := []struct {
		name              string
		inputKey          keyboard.Key
		inputData         interface{}
		expectedMethod    string
		expectedIntensity int
		expectReset       bool
		expectPanic       bool
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
			inputKey:    keyboard.X,
			inputData:   keyboard.KeyEvent{Key: keyboard.X},
			expectReset: true,
			expectPanic: false,
		},
		{
			name:        "Scenario 12: Handle Invalid Input Data Type (Panic Scenario)",
			inputData:   "this is not a KeyEvent",
			expectReset: false,
			expectPanic: true,
		},
	}

	allMethods := []string{
		"Left", "Right", "Up", "Down", "Forward", "Backward",
		"CounterClockwise", "Clockwise", "Land", "TakeOff",
	}

	resetMethods := []string{
		"Forward", "Backward", "Up", "Down", "Left", "Right", "Clockwise",
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			mockDriver := new(MockTelloDriver)
			handler := handleKeyboardInput(mockDriver)
			panicked := false

			defer func() {
				r := recover()
				if r != nil {
					panicked = true
					if !tc.expectPanic {

						t.Errorf("Test panicked unexpectedly: %v\n%s", r, string(debug.Stack()))
					} else {

						t.Logf("Caught expected panic: %v", r)
					}
				} else if tc.expectPanic {

					t.Errorf("Expected test to panic, but it did not")
				}

				if !panicked && !tc.expectPanic {
					if tc.expectedMethod != "" {

						switch tc.expectedMethod {
						case "Land", "TakeOff":
							mockDriver.AssertCalled(t, tc.expectedMethod)
							t.Logf("Successfully asserted call to %s()", tc.expectedMethod)
						default:
							mockDriver.AssertCalled(t, tc.expectedMethod, tc.expectedIntensity)
							t.Logf("Successfully asserted call to %s(%d)", tc.expectedMethod, tc.expectedIntensity)
						}

						for _, method := range allMethods {
							if method != tc.expectedMethod {
								mockDriver.AssertNotCalled(t, method)
							}
						}

						if !tc.expectReset {
							for _, method := range resetMethods {

								mockDriver.AssertNotCalled(t, method, 0)
							}
						}
						t.Logf("Successfully asserted other methods were not called for %s", tc.expectedMethod)

					} else if tc.expectReset {

						mockDriver.AssertCalled(t, "Forward", 0)
						mockDriver.AssertCalled(t, "Backward", 0)
						mockDriver.AssertCalled(t, "Up", 0)
						mockDriver.AssertCalled(t, "Down", 0)
						mockDriver.AssertCalled(t, "Left", 0)
						mockDriver.AssertCalled(t, "Right", 0)
						mockDriver.AssertCalled(t, "Clockwise", 0)
						mockDriver.AssertNotCalled(t, "CounterClockwise")
						mockDriver.AssertNotCalled(t, "Land")
						mockDriver.AssertNotCalled(t, "TakeOff")
						t.Logf("Successfully asserted calls for resetDronePostion")

					} else {

						t.Logf("No specific method call or reset expected and none asserted.")
					}
				} else if panicked && tc.expectPanic {

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

			}()

			t.Logf("Invoking handler with input data: %#v", tc.inputData)
			handler(tc.inputData)

			if !tc.expectPanic && !panicked {
				t.Logf("Handler executed successfully without panic.")
			}

		})
	}
}

