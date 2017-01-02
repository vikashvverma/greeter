package job

import (
	"errors"
	"fmt"

	"github.com/sendgrid/sendgrid-go"
	"github.com/vikashvverma/greeter/config"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type Greeter interface {
	Greet() error
}

type greetingMailer struct {
	config config.Config
}

func NewGreeter(c config.Config) Greeter {
	return &greetingMailer{config: c}
}

func (gm *greetingMailer) Greet() error {
	m := mail.NewV3MailInit(gm.config.From, gm.config.Subject, gm.config.To, gm.config.Content)
	request := sendgrid.GetRequest(gm.config.APIKey, "/v3/mail/send", "https://api.sendgrid.com")
	request.Method = "POST"
	request.Body = mail.GetRequestBody(m)
	response, err := sendgrid.API(request)
	if err != nil {
		return fmt.Errorf("Mail: could not send greeting to %s(%s): err:%s ", gm.config.To.Name, gm.config.To.Address, err)
	}

	if response.StatusCode >= 400 && response.StatusCode <= 500 {
		return errors.New("Mail: greeting not sent")
	}

	fmt.Printf("Greet: greeting sent successfully!: %#v", response)
	return nil
}
