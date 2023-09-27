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

	// Sleep for 10 seconds to allow the cron jobs to run
	time.Sleep(10 * time.Second)

	// Stop the cron instance
	c.Stop()

	// Assert that the cron jobs were run
	if c.EntryID("0 30 * * * *").Next.IsZero() {
		t.Errorf("Expected cron job to run every hour on the half hour, but it did not")
	}
	if c.EntryID("TZ=America/New_York 30 04 * * * *").Next.IsZero() {
		t.Errorf("Expected cron job to run at 04:30 New York time every day, but it did not")
	}
	if c.EntryID("@hourly").Next.IsZero() {
		t.Errorf("Expected cron job to run every hour, but it did not")
	}
	if c.EntryID("@every 0h0m1s").Next.IsZero() {
		t.Errorf("Expected cron job to run every 0h0m1s, but it did not")
	}
	if c.EntryID("@daily").Next.IsZero() {
		t.Errorf("Expected cron job to run every day, but it did not")
	}

	// Log that the test was successful
	t.Log("Test successful")
}
