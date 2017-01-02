package config

import (
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"io/ioutil"
	"encoding/json"
	"fmt"
)

type Config struct {
	Time string `json:"time"`

	APIKey string `json:"apiKey"`

	Subject string `json:"subject"`
	From    *mail.Email
	Content *mail.Content
}

func New(contents []byte) *Config {
	f := mail.NewEmail("Promice", "vikash@programminggeek.in")
	c := mail.NewContent("text/html", "Many happy returns of the day!")
	var config Config
	err := json.Unmarshal(contents, &config)
	if err != nil {
		return nil
	}
	config.Content = c
	config.From = f
	return &config
}

func ReadConfig(path string) *Config {
	contents, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Printf("could not read config file: %s", err)
		return nil
	}
	return New(contents)
}
