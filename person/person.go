package person

import "time"

type Person struct {
	Month time.Month
	Day   int
	Email string
	Name  string
}

func (p *Person) IsToday() bool {
	t := time.Now()
	_, m, d := t.Date()
	if m == p.Month && d == p.Day {
		return true
	}
	return false
}
