package clock

import (
	"fmt"
)

const (
	TestVersion   = 2 // This needs to be set?
	minutesInHour = 60
	hoursInDay    = 24
)

type Clock struct {
	hour   int
	minute int
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hour, c.minute)
}

func (c *Clock) wrapAround() {
	// Handle wrap-around in a naive sort of way
	for c.minute >= minutesInHour {
		c.minute -= minutesInHour
		c.hour++
	}
	for c.minute < 0 {
		c.minute += minutesInHour
		c.hour--
	}
	for c.hour >= hoursInDay {
		c.hour -= hoursInDay
	}
	for c.hour < 0 {
		c.hour += hoursInDay
	}
}

func (c Clock) Add(minutes int) Clock {
	c.minute += minutes
	c.wrapAround()
	return c
}

func Time(hour int, minute int) Clock {
	c := Clock{hour, minute}
	c.wrapAround()
	return c
}
