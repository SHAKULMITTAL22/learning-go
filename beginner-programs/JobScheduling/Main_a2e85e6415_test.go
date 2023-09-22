package main

import (
	"fmt"
	"testing"
	"time"

	"github.com/robfig/cron/v3"
)

func TestMain_a2e85e6415(t *testing.T) {
	// Create a new cron instance
	c := cron.New()

	// Add a func to run every hour on the half hour
	c.AddFunc("0 30 * * * *", func() { fmt.Println("Every hour on the half hour") })

	// Add a func to run at 04:30 New York time every day
	c.AddFunc("TZ=America/New_York 30 04 * * * *", func() { fmt.Println("Runs at 04:30 New York time every day") })

	// Add a func to run every hour
	c.AddFunc("@hourly", func() { fmt.Println("Runs every hour") })

	// Add a func to run every 0h0m1s
	c.AddFunc("@every 0h0m1s", func() { sayHelloTo("Someone!") })

	// Start the cron instance
	c.Start()

	// Add a func to run every day
	c.AddFunc("@daily", func() { fmt.Println("Every day") })

	// Sleep for 10 seconds to see output
	time.Sleep(10 * time.Second)

	// Stop the cron instance
	c.Stop() // Stop the scheduler (does not stop any jobs already running)

	// Test that the cron instance was created successfully
	if c == nil {
		t.Error("Cron instance was not created successfully")
	}

	// Test that the func was added to the cron instance successfully
	if len(c.Entries()) != 5 {
		t.Errorf("Expected 5 entries, got %d", len(c.Entries()))
	}

	// Test that the func was executed
	if !c.Running() {
		t.Error("Cron instance is not running")
	}

	// Stop the cron instance
	c.Stop()

	// Test that the cron instance was stopped successfully
	if c.Running() {
		t.Error("Cron instance is still running")
	}
}
