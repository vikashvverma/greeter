package main

import (
	"log"

	"github.com/vikashvverma/greeter/config"
	"github.com/vikashvverma/greeter/job"
)

func main() {
	c := config.New()
	g := job.NewGreeter(c)
	s := job.NewScheduler(c.Time, g)
	gocron, err := s.Schedule()
	if err != nil {
		log.Fatalf("ListenAndServe: %s", err)
	}
	<-gocron.Start()
}
