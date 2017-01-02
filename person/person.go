package person

import (
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"time"
)

type Person struct {
	Address string `json:"email"`

	Month time.Month `json:"month"`
	Day   int        `json:"day"`
	Name  string     `json:"name"`
}

func (p *Person) Email() *mail.Email {
	return mail.NewEmail(p.Name, p.Address)
}

func (p *Person) IsToday() bool {
	t := time.Now()
	_, m, d := t.Date()
	if m == p.Month && d == p.Day {
		return true
	}
	return false
}
