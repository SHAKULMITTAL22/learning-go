package Validator

import (
	"testing"
)

// TestIsValidEmail_d2603fb18f checks the validity of IsValidEmail function
func TestIsValidEmail_d2603fb18f(t *testing.T) {
	// Defining test cases
	var tests = []struct {
		input    string
		expected bool
	}{
		{"test@example.com", true},
		{"example.com", false},
		{"example@.com", false},
		{"#", false},
		{"", false},
		{"testemail@gmail.com", true},
		{"invalidemail@invalid", false},
		{"test@com", false},
        {"test@invalid..com", false}, // new test case: multiple periods are not allowed after @ 
        {"test.@invalid.com", false}, // new test case: period is not allowed after local part
        {"test..test@invalid.com", false}, // new test case: multiple periods are not allowed in local part 
        {"test.test@invalid.com", true}, // new test case: single period is allowed in local part
        {"test.test@invalid.co.in", true}, // new test case: mulitple period allowed after @ 
	}

	// Running test cases
	for _, test := range tests {
		if output := IsValidEmail(test.input); output != test.expected {
			t.Errorf("Test failed: %v inputted, %v expected, recieved: %v", test.input, test.expected, output)
		}else{
			t.Logf("Test success: %v inputted, %v expected, recieved: %v", test.input, test.expected, output)
        }
	}
}
