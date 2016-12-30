package main

import (
	"github.com/vikashvverma/greeter/config"
	"github.com/vikashvverma/greeter/mailer"
)

func main() {
	c := config.New()
	m := mailer.New(c)
	m.Greet()
}
