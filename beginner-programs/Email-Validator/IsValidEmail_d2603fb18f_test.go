package Validator

import (
	"testing"
)

func TestIsValidEmail_d2603fb18f(t *testing.T) {
	cases := []struct {
		email string
		want  bool
	}{
		{"correct_email@domain.com", true},
		{"incorrect_email", false},
		{"exceeding_length_email@domain.com.com.com.com.com.com.com.com.com.com.com.com.com.com.com.com.com.com.com.com.com.com.com.com.com.com.com.com.com.com.com.com.com.com.com.com.com.com.com.com.com.com.com.com.com.com.com.com.com.com.com.com.com.com.com.com.com.com", false},
	}

	for _, testcase := range cases {
		got := IsValidEmail(testcase.email)

		if got != testcase.want {
			t.Errorf("IsValidEmail(%q) = %v; want %v", testcase.email, got, testcase.want)
		} else {
			t.Logf("Success: IsValidEmail(%q) = %v", testcase.email, got)
		}
	}
}
