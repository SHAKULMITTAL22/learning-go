package main

import (
    "fmt"
    "testing"
)

func TestPrint_f973b185d0(t *testing.T) {
    tests := []struct {
        name string
        args []interface{}
        want int
        wantErr bool
    }{
        {"success", []interface{}{"Hello, world!"}, 13, false},
        {"error", []interface{}{fmt.Errorf("error")}, 0, true},
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got, err := Print(tt.args...)
            if (err != nil) != tt.wantErr {
                t.Errorf("Print() error = %v, wantErr %v", err, tt.wantErr)
                return
            }
            if got != tt.want {
                t.Errorf("Print() = %v, want %v", got, tt.want)
            }
        })
    }
}
