package config

import "github.com/sendgrid/sendgrid-go/helpers/mail"

type Config struct {
	APIKey  string
	From    *mail.Email
	Subject string
	To      *mail.Email
	Content *mail.Content
}

func New() Config {
	f := mail.NewEmail("Promice", "vikash@programminggeek.in")
	t := mail.NewEmail("Vikash", "vermavikash014@gmail.com")
	c := mail.NewContent("text/plain", "Many happy returns of the day!")
	return Config{
		APIKey: "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",
		From:   f,
		Subject:"Happy B'day!",
		To:     t,
		Content:c,
	}
}
