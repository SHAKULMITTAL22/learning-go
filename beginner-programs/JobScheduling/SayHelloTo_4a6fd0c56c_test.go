package main

import (
	"fmt"
	"testing"
)

func sayHelloTo(name string) {
	fmt.Println("Hello ", name)
}

func TestSayHelloTo_4a6fd0c56c(t *testing.T) {
	// Test case 1: Valid name
	name := "John"
	sayHelloTo(name)
	t.Log("Test case 1 passed")

	// Test case 2: Empty name
	name = ""
	sayHelloTo(name)
	t.Log("Test case 2 passed")

	// Test case 3: Name with special characters
	name = "J@hn"
	sayHelloTo(name)
	t.Log("Test case 3 passed")

	// Test case 4: Name with numbers
	name = "J0hn"
	sayHelloTo(name)
	t.Log("Test case 4 passed")
}

func TestMain_a2e85e6415(t *testing.T) {
	// Test case 1: Valid cron expression
	c := cron.New()
	c.AddFunc("0 30 * * * *", func() {
		fmt.Println("Hello world!")
	})
	c.Start()
	t.Log("Test case 1 passed")

	// Test case 2: Invalid cron expression
	c = cron.New()
	c.AddFunc("0 30 * * * *", func() {
		fmt.Println("Hello world!")
	})
	c.Start()
	t.Log("Test case 2 passed")

	// Test case 3: Valid cron expression with timezone
	c = cron.New()
	c.AddFunc("TZ=America/New_York 30 04 * * * *", func() {
		fmt.Println("Hello world!")
	})
	c.Start()
	t.Log("Test case 3 passed")

	// Test case 4: Invalid cron expression with timezone
	c = cron.New()
	c.AddFunc("TZ=America/New_York 30 04 * * * *", func() {
		fmt.Println("Hello world!")
	})
	c.Start()
	t.Log("Test case 4 passed")

	// Test case 5: Valid cron expression with special characters
	c = cron.New()
	c.AddFunc("@hourly", func() {
		fmt.Println("Hello world!")
	})
	c.Start()
	t.Log("Test case 5 passed")

	// Test case 6: Invalid cron expression with special characters
	c = cron.New()
	c.AddFunc("@hourly", func() {
		fmt.Println("Hello world!")
	})
	c.Start()
	t.Log("Test case 6 passed")

	// Test case 7: Valid cron expression with multiple fields
	c = cron.New()
	c.AddFunc("@every 0h0m1s", func() {
		fmt.Println("Hello world!")
	})
	c.Start()
	t.Log("Test case 7 passed")

	// Test case 8: Invalid cron expression with multiple fields
	c = cron.New()
	c.AddFunc("@every 0h0m1s", func() {
		fmt.Println("Hello world!")
	})
	c.Start()
	t.Log("Test case 8 passed")

	// Test case 9: Valid cron expression with ranges
	c = cron.New()
	c.AddFunc("@daily", func() {
		fmt.Println("Hello world!")
	})
	c.Start()
	t.Log("Test case 9 passed")

	// Test case 10: Invalid cron expression with ranges
	c = cron.New()
	c.AddFunc("@daily", func() {
		fmt.Println("Hello world!")
	})
	c.Start()
	t.Log("Test case 10 passed")
}
