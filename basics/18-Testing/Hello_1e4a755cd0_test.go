// Test generated by RoostGPT for test go-open-test using AI Type Vertex AI and AI Model code-bison

package main

import (
	"testing"
)

func TestHello_1e4a755cd0(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{"Bob", "Hello Bob"},
		{"Alice", "Hello Alice"},
		{"", "Hello "},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Hello(tt.name); got != tt.want {
				t.Errorf("Hello() = %q, want %q", got, tt.want)
			}
		})
	}
}
