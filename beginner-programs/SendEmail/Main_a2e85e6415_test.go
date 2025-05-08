package main

import (
	"net/smtp"
	"testing"
)

func TestMain_a2e85e6415(t *testing.T) {
	t.Log("Given the main function to send an email")
	{
		// Choose auth method and set it up
		auth := smtp.PlainAuth("", "email", "password", "smtp-client")

		// Email details
		from := "email"
		to := "email"
		subject := "Hello there!"
		body := "Hello there my friend!"

		msg := "From: " + from + "\n" +
			"To: " + to + "\n" +
			"Subject: " + subject + "\n\n" +
			body

		t.Log("When the email is sent with valid credentials")
		{
			err := smtp.SendMail("email-client", auth, "email", []string{to}, []byte(msg))
			if err != nil {
				t.Errorf("smtp.SendMail failed: %v", err)
			}
			t.Log("Then the email is sent successfully")
		}

		t.Log("When the email is sent with invalid credentials")
		{
			err := smtp.SendMail("email-client", auth, "email", []string{to}, []byte(msg))
			if err == nil {
				t.Error("smtp.SendMail should have failed with invalid credentials")
			}
			t.Log("Then the email is not sent")
		}
	}
}
