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

    // Add a function to be run every hour on the half hour
    c.AddFunc("0 30 * * * *", func() { fmt.Println("Every hour on the half hour") })

    // Add a function to be run at 04:30 New York time every day
    c.AddFunc("TZ=America/New_York 30 04 * * * *", func() { fmt.Println("Runs at 04:30 New York time every day") })

    // Add a function to be run every hour
    c.AddFunc("@hourly", func() { fmt.Println("Runs every hour") })

    // Add a function to be run every 0 hours, 0 minutes, and 1 second
    c.AddFunc("@every 0h0m1s", func() { sayHelloTo("Someone!") })

    // Start the cron instance
    c.Start()

    // Add a function to be run every day
    c.AddFunc("@daily", func() { fmt.Println("Every day") })

    // Sleep for 10 seconds to allow the cron jobs to run
    time.Sleep(10 * time.Second)

    // Stop the cron instance
    c.Stop()

    // Test that the cron jobs were run
    if !c.Entry("0 30 * * * *").Next.IsZero() {
        t.Errorf("Expected cron job to be run every hour on the half hour, but it was not")
    }
    if !c.Entry("TZ=America/New_York 30 04 * * * *").Next.IsZero() {
        t.Errorf("Expected cron job to be run at 04:30 New York time every day, but it was not")
    }
    if !c.Entry("@hourly").Next.IsZero() {
        t.Errorf("Expected cron job to be run every hour, but it was not")
    }
    if !c.Entry("@every 0h0m1s").Next.IsZero() {
        t.Errorf("Expected cron job to be run every 0 hours, 0 minutes, and 1 second, but it was not")
    }
    if !c.Entry("@daily").Next.IsZero() {
        t.Errorf("Expected cron job to be run every day, but it was not")
    }
}
