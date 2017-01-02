package job

import (
	"errors"
	"fmt"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"

	"encoding/json"
	"github.com/vikashvverma/greeter/config"
	"github.com/vikashvverma/greeter/person"
	"io/ioutil"
)

type Greeter interface {
	Greet() []error
}

type greetingMailer struct {
	config *config.Config
}

const (
	endpoint = "/v3/mail/send"
	host     = "https://api.sendgrid.com"
	method   = "POST"
)

func NewGreeter(c *config.Config) Greeter {
	return &greetingMailer{config: c}
}

func (gm *greetingMailer) Greet() []error {
	p := gm.people()
	request := sendgrid.GetRequest(gm.config.APIKey, endpoint, host)
	var errs []error
	for _, person := range p {
		if person.IsToday() == false {
			continue
		}
		email := person.Email()
		m := mail.NewV3MailInit(gm.config.From, gm.config.Subject, email, gm.config.Content)
		request.Method = method
		request.Body = mail.GetRequestBody(m)
		response, err := sendgrid.API(request)
		if err != nil {
			fmt.Printf("Mail: could not send greeting to %s(%s): err:%s ", email.Name, email.Address, err)
			errs = append(errs, fmt.Errorf("Mail: could not send greeting to %s(%s): err:%s ", email.Name, email.Address, err))
			continue
		}
		if response.StatusCode >= 400 && response.StatusCode <= 500 {
			fmt.Printf("Mail: greeting not sent: %#v", response)
			errs = append(errs, errors.New("Mail: greeting not sent"))
			continue
		}
		fmt.Printf("Greet: greeting sent successfully to: %s", person.Email)
	}
	return errs
}

func (gm *greetingMailer) people() []person.Person {
	var people []person.Person
	dobs, err := ioutil.ReadFile("dob.json")
	if err != nil {
		fmt.Println("people: could not read people's data!")
		return nil
	}
	err = json.Unmarshal(dobs, &people)
	if err != nil {
		fmt.Printf("people: could unmarshal people's data : err: %s", err)
		return nil
	}
	return people
}
