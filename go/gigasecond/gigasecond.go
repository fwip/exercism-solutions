package gigasecond

import (
	"time"
)

const (
	TestVersion = 2
)

func AddGigasecond(t time.Time) time.Time {
	oneGigaSecond, _ := time.ParseDuration("1000000000s")

	return t.Add(oneGigaSecond)
}

var Birthday, _ = time.Parse("01-02-2006", "06-17-1987")
