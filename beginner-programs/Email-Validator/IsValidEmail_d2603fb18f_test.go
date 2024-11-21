package Validator

import (
    "regexp"
    "testing"
)

var emailRegexp = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

func TestIsValidEmail_d2603fb18f(t *testing.T) {
    tests := []struct {
        name  string
        email string
        want  bool
    }{
        {"Test with valid email", "example@gmail.com", true},
        {"Test with invalid email", "examplegmail.com", false},
        {"Test with empty string", "", false},
        {"Test with long email", "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa@gmail.com", false},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := IsValidEmail(tt.email); got != tt.want {
                t.Errorf("IsValidEmail() = %v, want %v, email: %v", got, tt.want, tt.email)
            } else {
                t.Logf("Success: IsValidEmail() = %v, want %v, email: %v", got, tt.want, tt.email)
            }

        })
    }
}
