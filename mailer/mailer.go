package mailer

import (
	"fmt"
	"errors"

	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"github.com/sendgrid/sendgrid-go"

	"github.com/vikashvverma/greeter/config"
)

type Mailer interface {
	Greet() error
}

type greetingMailer struct {
	config config.Config
}

func New(c config.Config) Mailer {
	return &greetingMailer{config:c}
}

func (gm *greetingMailer) Greet() error {
	m := mail.NewV3MailInit(gm.config.From, gm.config.Subject, gm.config.To, gm.config.Content)
	request := sendgrid.GetRequest(gm.config.APIKey, "/v3/mail/send", "https://api.sendgrid.com")
	request.Method = "POST"
	request.Body = mail.GetRequestBody(m)
	response, err := sendgrid.API(request)
	if err != nil {
		return fmt.Errorf("Mail: could not send greeting to %s(%s): err: ", gm.config.To.Name, gm.config.To.Address, err)
	}

	if response.StatusCode >= 400 && response.StatusCode <= 500 {
		return errors.New("Mail: greeting not sent")
	}
	
	fmt.Printf("Greet: greeting sent successfully!: %#v", response)
	return nil
}
