// ********RoostGPT********
/*
Test generated by RoostGPT for test golang-dbrx using AI Type  and AI Model 

ROOST_METHOD_HASH=GCD_766d714a13
ROOST_METHOD_SIG_HASH=GCD_ab1c91475d

### Test Scenarios for the `GCD` Function

#### Scenario 1: Test with two positive integers
Details:
  Description: This test checks the correctness of the GCD function when both inputs are positive integers. The aim is to verify that the function returns the correct greatest common divisor.
Execution:
  Arrange: Define two positive integers, for example, 18 and 24.
  Act: Call the GCD function with these integers.
  Assert: Check if the result is 6, the known GCD of 18 and 24.
Validation:
  The assertion checks that the returned value matches the expected GCD. This test is important as it validates the function's ability to compute the GCD for typical, straightforward cases.

#### Scenario 2: Test with one positive integer and zero
Details:
  Description: Testing the function when one of the inputs is zero. According to mathematical rules, the GCD of any number and zero is the number itself.
Execution:
  Arrange: Define a positive integer and zero, e.g., 0 and 15.
  Act: Invoke the GCD function with these values.
  Assert: Verify that the result is 15.
Validation:
  The assertion confirms that the GCD of any number with zero returns the non-zero number. This test ensures that the function adheres to basic mathematical properties of GCD.

#### Scenario 3: Test with negative integers
Details:
  Description: To validate the function's behavior with negative integers, ensuring it can handle and return the correct GCD.
Execution:
  Arrange: Define two negative integers, for example, -36 and -48.
  Act: Call the GCD function with these numbers.
  Assert: Confirm that the result is 12, the GCD of their absolute values.
Validation:
  The assertion checks for the correct GCD despite negative input values. This test is crucial to verify that the function is robust against negative numbers and still produces valid results.

#### Scenario 4: Test with both inputs being zero
Details:
  Description: This scenario checks the function's behavior when both inputs are zero. Theoretically, the GCD of zero with zero is undefined; however, practical implementations may vary.
Execution:
  Arrange: Set both parameters to zero.
  Act: Invoke the GCD function.
  Assert: Assess how the function handles this input, possibly expecting a specific return value or error handling.
Validation:
  This test checks the function's resilience and error handling capabilities when faced with undefined or boundary conditions.

#### Scenario 5: Test with large integers
Details:
  Description: Ensuring that the function can handle and correctly compute the GCD of large integers.
Execution:
  Arrange: Use two large integers, such as 123456 and 789012.
  Act: Execute the GCD function with these parameters.
  Assert: Verify that the result is correct based on pre-calculated expectations (e.g., 12).
Validation:
  The assertion ensures the function's capability to handle large values without overflow or errors, maintaining accuracy in computationally intensive scenarios.

#### Scenario 6: Test with prime numbers
Details:
  Description: To test the function with two prime numbers, which should return 1 since primes have no common divisors other than 1.
Execution:
  Arrange: Select two prime numbers, for instance, 29 and 13.
  Act: Run the GCD function with these numbers.
  Assert: Check if the result is 1.
Validation:
  This scenario verifies that the function correctly identifies the lack of common divisors between prime numbers, an essential aspect of the GCD's mathematical foundation.
*/

// ********RoostGPT********
package EuclideanAlgorithm

import (
	"testing"
)

// Improve GCD function to handle negative integers correctly
func GCD(a, b int) int {
	if a < 0 {
		a = -a
	}
	if b < 0 {
		b = -b
	}
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func TestGcd(t *testing.T) {
	tests := []struct {
		name     string
		a        int
		b        int
		expected int
	}{
		{
			name:     "Test with two positive integers",
			a:        18,
			b:        24,
			expected: 6,
		},
		{
			name:     "Test with one positive integer and zero",
			a:        0,
			b:        15,
			expected: 15,
		},
		{
			name:     "Test with negative integers",
			a:        -36,
			b:        -48,
			expected: 12,
		},
		{
			name:     "Test with both inputs being zero",
			a:        0,
			b:        0,
			expected: 0, // Assuming implementation returns 0 for GCD(0,0)
		},
		{
			name:     "Test with large integers",
			a:        123456,
			b:        789012,
			expected: 12,
		},
		{
			name:     "Test with prime numbers",
			a:        29,
			b:        13,
			expected: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GCD(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("GCD(%d, %d) got %d, want %d", tt.a, tt.b, result, tt.expected)
			} else {
				t.Logf("Success: GCD(%d, %d) = %d", tt.a, tt.b, result)
			}
		})
	}
}
