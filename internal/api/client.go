package api

import (
	"fmt"
	"net/smtp"
)

type Client struct {
	SenderEmail    string
	SenderPassword string
	TargetEmail    string
}

func New(senderEmail, senderPassword, targetEmail string) *Client {
	return &Client{
		SenderEmail:    senderEmail,
		SenderPassword: senderPassword,
		TargetEmail:    targetEmail,
	}
}

func (c *Client) CreateEntry(text string, tags []string) error {
	// Day One supports parsing tags if they are hashtags in the body.
	// We append them to the bottom of the email.
	var tagString string
	if len(tags) > 0 {
		tagString = "\n\n"
		for _, t := range tags {
			tagString += fmt.Sprintf("#%s ", t)
		}
	}

	body := text + tagString
	
	// Setup headers
	// Subject is usually used as the entry title by Day One, 
	// or the first line if subject is empty. We'll leave subject generic or use first few words.
	subject := "New Journal Entry" 
	
	msg := []byte("To: " + c.TargetEmail + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"\r\n" +
		body + "\r\n")

	// Gmail SMTP server configuration.
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// Authentication.
	auth := smtp.PlainAuth("", c.SenderEmail, c.SenderPassword, smtpHost)

	// Sending email.
	fmt.Println("Sending via Gmail...")
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, c.SenderEmail, []string{c.TargetEmail}, msg)
	if err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}
	
	fmt.Println("Entry sent successfully!")
	return nil
}
