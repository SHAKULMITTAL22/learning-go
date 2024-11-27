// ********RoostGPT********
/*
Test generated by RoostGPT for test golang-1 using AI Type Claude AI and AI Model claude-3-5-sonnet-20240620

ROOST_METHOD_HASH=IsValidEmail_9ebd397bff
ROOST_METHOD_SIG_HASH=IsValidEmail_d2603fb18f

Existing Test Information:
These test cases are already implemented and not included for test generation scenario:
File: learning-go/beginner-programs/Email-Validator/IsValidEmail_test.go
Test Cases:
    [TestIsValidEmail]

Based on the provided function and existing tests, here are additional test scenarios for the IsValidEmail function:

Scenario 1: Email with multiple dots in domain

Details:
  Description: Test if an email address with multiple dots in the domain part is considered valid.
Execution:
  Arrange: Prepare an email address with multiple dots in the domain.
  Act: Call IsValidEmail with the prepared email address.
  Assert: Check if the function returns true.
Validation:
  This test ensures that the function correctly handles valid email addresses with multiple dots in the domain, which are allowed according to email standards.

Scenario 2: Email with special characters in local part

Details:
  Description: Verify that an email address with allowed special characters in the local part is considered valid.
Execution:
  Arrange: Create an email address with special characters like !#$%&'*+-/=?^_`{|}~ in the local part.
  Act: Pass the email address to IsValidEmail.
  Assert: Confirm that the function returns true.
Validation:
  This test checks if the function correctly handles valid special characters in the local part of the email address, as per email specifications.

Scenario 3: Email with IP address as domain

Details:
  Description: Test if an email address using an IP address as the domain is considered valid.
Execution:
  Arrange: Prepare an email address with an IP address as the domain (e.g., user@[192.168.0.1]).
  Act: Call IsValidEmail with the prepared email address.
  Assert: Verify that the function returns true.
Validation:
  This scenario tests the function's ability to recognize valid email addresses that use IP addresses instead of domain names, which is allowed in email standards.

Scenario 4: Email with quoted local part

Details:
  Description: Check if an email address with a quoted local part is considered valid.
Execution:
  Arrange: Create an email address with a quoted local part (e.g., "John Doe"@example.com).
  Act: Pass the email address to IsValidEmail.
  Assert: Confirm that the function returns true.
Validation:
  This test ensures that the function correctly handles quoted local parts in email addresses, which are valid according to email specifications.

Scenario 5: Email with very short local part and domain

Details:
  Description: Verify that an email address with minimum length local part and domain is considered valid.
Execution:
  Arrange: Prepare an email address with a single character local part and domain (e.g., a@b.c).
  Act: Call IsValidEmail with the prepared email address.
  Assert: Check if the function returns true.
Validation:
  This scenario tests the function's ability to handle valid email addresses with the shortest possible local part and domain, ensuring it doesn't incorrectly reject them.

Scenario 6: Email with non-ASCII characters

Details:
  Description: Test if an email address containing non-ASCII characters is considered invalid.
Execution:
  Arrange: Create an email address with non-ASCII characters (e.g., üser@exämple.com).
  Act: Pass the email address to IsValidEmail.
  Assert: Verify that the function returns false.
Validation:
  This test checks if the function correctly rejects email addresses with non-ASCII characters, which are not allowed in standard email addresses without proper encoding.

Scenario 7: Email with valid subdomain

Details:
  Description: Verify that an email address with a valid subdomain is considered valid.
  Execution:
    Arrange: Prepare an email address with a subdomain (e.g., user@subdomain.example.com).
    Act: Call IsValidEmail with the prepared email address.
    Assert: Confirm that the function returns true.
  Validation:
    This scenario ensures that the function correctly handles email addresses with subdomains, which are common and valid in email addresses.

These additional test scenarios cover various aspects of email validation that weren't explicitly covered in the existing tests, including special cases and edge conditions that are important for comprehensive email validation.
*/

// ********RoostGPT********
package Validator

import (
	"testing"
)

func TestIsValidEmail(t *testing.T) {
	tests := []struct {
		name  string
		email string
		want  bool
	}{
		{
			name:  "Email with multiple dots in domain",
			email: "user@example.co.uk",
			want:  true,
		},
		{
			name:  "Email with special characters in local part",
			email: "user!#$%&'*+-/=?^_`{|}~@example.com",
			want:  true,
		},
		{
			name:  "Email with IP address as domain",
			email: "user@[192.168.0.1]",
			want:  true,
		},
		{
			name:  "Email with quoted local part",
			email: "\"John Doe\"@example.com",
			want:  true,
		},
		{
			name:  "Email with very short local part and domain",
			email: "a@b.c",
			want:  true,
		},
		{
			name:  "Email with non-ASCII characters",
			email: "üser@exämple.com",
			want:  false,
		},
		{
			name:  "Email with valid subdomain",
			email: "user@subdomain.example.com",
			want:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValidEmail(tt.email); got != tt.want {
				t.Errorf("IsValidEmail() = %v, want %v", got, tt.want)
			}
		})
	}
}
