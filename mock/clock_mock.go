package mock

import "time"

type ClockMock struct {
	NowMock time.Time
	AddMock time.Time
}

func (c *ClockMock) Now() time.Time {
	return c.NowMock
}

func (c *ClockMock) Add(t time.Time, d time.Duration) time.Time {
	return c.AddMock
}
