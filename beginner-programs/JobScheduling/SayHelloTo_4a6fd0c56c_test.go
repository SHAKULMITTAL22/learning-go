package main

import (
	"testing"
)

func TestSayHelloTo_4a6fd0c56c(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"Valid name", args{"John"}, false},
		{"Empty name", args{""}, true},
		{"Name with numbers", args{"John123"}, false},
		{"Name with special characters", args{"John!"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := sayHelloTo(tt.args.name); (err != nil) != tt.wantErr {
				t.Errorf("sayHelloTo() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
