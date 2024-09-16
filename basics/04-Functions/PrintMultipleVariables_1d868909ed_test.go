package main

import (
	"fmt"
	"reflect"
	"testing"
)

func printMultipleVariables(i ...interface{}) {
	for _, v := range i {
		fmt.Println(v, "--", reflect.ValueOf(v).Kind())
	}
}

func TestPrintMultipleVariables_1d868909ed(t *testing.T) {
	// Test case 1: Print multiple variables of different types
	printMultipleVariables(1, "Hello", 3.14, true)

	// Expected output:
	// 1 -- int
	// Hello -- string
	// 3.14 -- float64
	// true -- bool

	// Test case 2: Print an empty slice
	var slice []int
	printMultipleVariables(slice)

	// Expected output:
	// [] -- slice

	// Test case 3: Print a nil pointer
	var ptr *int
	printMultipleVariables(ptr)

	// Expected output:
	// <nil> -- ptr

	// Test case 4: Print a struct
	type person struct {
		name string
		age  int
	}
	p := person{"John", 30}
	printMultipleVariables(p)

	// Expected output:
	// {John 30} -- struct

	// Test case 5: Print a map
	m := map[string]int{"apple": 1, "banana": 2}
	printMultipleVariables(m)

	// Expected output:
	// map[apple:1 banana:2] -- map

	// Test case 6: Print a function
	f := func() {
		fmt.Println("Hello")
	}
	printMultipleVariables(f)

	// Expected output:
	// func() -- func
}
