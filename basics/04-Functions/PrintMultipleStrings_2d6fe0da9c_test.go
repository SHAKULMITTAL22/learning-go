package main

import (
    "fmt"
    "reflect"
    "testing"
)

func printMultipleStrings(s ...string) {
    for i := 0; i < len(s); i++ {
        fmt.Println(s[i])
    }
}

func TestPrintMultipleStrings_2d6fe0da9c(t *testing.T) {
    type args struct {
        s []string
    }
    tests := []struct {
        name string
        args args
        want []string
    }{
        {
            name: "test_case_1",
            args: args{s: []string{"Hello", "World"}},
            want: []string{"Hello", "World"},
        },
        {
            name: "test_case_2",
            args: args{s: []string{"Hello", "World", "!"}},
            want: []string{"Hello", "World", "!"},
        },
        {
            name: "test_case_3",
            args: args{s: []string{}},
            want: []string{},
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := printMultipleStrings(tt.args.s...); !reflect.DeepEqual(got, tt.want) {
                t.Errorf("printMultipleStrings() = %v, want %v", got, tt.want)
            }
        })
    }
}
