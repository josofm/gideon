package clock

import "time"

type TimeClock struct{}

func (c *TimeClock) Now() time.Time {
	return time.Now()
}

func (c *TimeClock) Add(t time.Time, d time.Duration) time.Time {
	return t.Add(d)
}
