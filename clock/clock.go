package clock

import (
	"fmt"
)

type Clock struct {
	hour, minute int
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hour, c.minute)
}

func FormatClock(hour, minute int) Clock {
	if minute > 59 || minute < -59 {
		hour = hour + (minute / 60)
		minute = minute % 60
	}

	if minute < 0 && minute >= -59 {
		hour--
		minute = (60 + minute) % 60
	}

	if hour > 23 || hour < -23 {
		hour = hour % 24
	}

	if hour < 0 {
		hour = (24 + hour) % 24
	}

	return Clock{hour, minute}
}

func New(hour, minute int) Clock {
	return FormatClock(hour, minute)
}

func (c Clock) Add(add int) Clock {
	return FormatClock(c.hour, c.minute+add)
}

func (c Clock) Subtract(subs int) Clock {
	return FormatClock(c.hour, c.minute-subs)
}

func Compare(a, b Clock) bool {
	result := false
	clock1 := New(a.hour, a.minute)
	clock2 := New(b.hour, b.minute)

	if clock1 == clock2 {
		result = true
	}
	return result
}
