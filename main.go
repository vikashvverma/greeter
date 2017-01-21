package main

import (
	"log"

	"github.com/vikashvverma/greeter/config"
	"github.com/vikashvverma/greeter/job"
)

func main() {
	c, err := config.ReadConfig("./config.json")
	if err != nil {
		log.Fatal("Could not read config file!")
	}
	g := job.NewGreeter(c)
	s := job.NewScheduler(c.Time, g)
	gocron, err := s.Schedule()
	if err != nil {
		log.Fatalf("main: %s", err)
	}
	<-gocron.Start()
}
