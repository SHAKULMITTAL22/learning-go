package main

import (
  "fmt"
  "testing"
	"time"

	"github.com/robfig/cron/v3"
)

// checkIfRan is a wrapper function for checking if a certain cron job ran
func checkIfRan(schedule string, job func(), delay time.Duration) (bool, error) {
  ran := make(chan bool, 1)

  c := cron.New(cron.WithSeconds())

	// Add function on cron schedule and flip ran to true when executed
	c.AddFunc(schedule, func() { job(); ran <- true })
	
  go c.Start()

  select {
  // Return after the delay
  case <-time.After(delay):

  // If cron job doesn't run within given delay, return error
  case <-ran:
  	return true, nil
  }

  c.Stop()
  return false, fmt.Errorf("cron did not run")
}

// Testing the main function
func TestMain(t *testing.T) {
  // Creating a mock job function for the test
  mockJob := func() {
    fmt.Println("Testing cron job")
  }

  // Test cases
  tests := []struct {
    schedule string
    job      func()
    delay    time.Duration
  }{
    {"@every 1s", mockJob, 2 * time.Second},
  }

  // Run test cases
  for _, test := range tests {

    // Check if each cron job ran within its schedule
    ran, err := checkIfRan(test.schedule, test.job, test.delay)

    if err != nil {
      t.Fatalf("Test failed for schedule: %s, Error: %s", test.schedule, err.Error())
    }

    if !ran {
      t.Errorf("Test failed for schedule: %s", test.schedule)
    } else {
      t.Logf("Test passed for schedule: %s", test.schedule)
    }
  }
}
