package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type Config struct {
	Time string `json:"time"`

	APIKey string `json:"apiKey"`

	Subject string        `json:"subject"`
	From    *mail.Email   `json:"from"`
	TOs     []*mail.Email `json:"tos"`
	Ccs     []*mail.Email `json:"ccs"`
	Bcc     []*mail.Email `json:"bccs"`
	Content *mail.Content
}

func New(contents []byte) *Config {
	var config Config
	err := json.Unmarshal(contents, &config)
	if err != nil {
		return nil
	}
	return &config
}

func ReadConfig(path string) (*Config, error) {
	contents, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Printf("could not read config file: %s", err)
		return nil, fmt.Errorf("could not read config file: %s", err)
	}
	return New(contents), nil
}
