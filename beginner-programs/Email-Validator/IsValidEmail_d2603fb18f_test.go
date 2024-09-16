package Validator

import (
	"testing"
)

func TestIsValidEmail_d2603fb18f(t *testing.T) {
	type args struct {
		email string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "valid email",
			args: args{email: "example@example.com"},
			want: true,
		},
		{
			name: "invalid email",
			args: args{email: "example"},
			want: false,
		},
		{
			name: "empty email",
			args: args{email: ""},
			want: false,
		},
		{
			name: "email with spaces",
			args: args{email: "example@ example.com"},
			want: false,
		},
		{
			name: "email with invalid characters",
			args: args{email: "example@example.com!"},
			want: false,
		},
		{
			name: "email with too many characters",
			args: args{email: "example@example.com" + "a" * 255},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValidEmail(tt.args.email); got != tt.want {
				t.Errorf("IsValidEmail() = %v, want %v", got, tt.want)
			}
		})
	}
}
