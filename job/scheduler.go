package job

import (
	"github.com/jasonlvhit/gocron"
)

type Scheduler interface {
	Schedule() (*gocron.Scheduler, error)
}

type greetingScheduler struct {
	time    string
	greeter Greeter
}

func NewScheduler(time string, g Greeter) Scheduler {
	return &greetingScheduler{time: time, greeter: g}
}

func (gs *greetingScheduler) Schedule() (*gocron.Scheduler, error) {
	s := gocron.NewScheduler()
	s.Every(1).Day().At(gs.time).Do(gs.greeter.Greet)
	return s, nil
}
