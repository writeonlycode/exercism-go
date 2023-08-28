// Package clock defines a Clock type, with functions to display and manipulate
// time.
package clock

import "fmt"

// Clock represents a a time in hours and minutes in 24-hours format.
type Clock struct {
	h, m int
}

// New creates a new clock. It handles positive and negative values correctly.
func New(h, m int) Clock {
	totalMinutes := h*60 + m

	hours := totalMinutes / 60
	hours = hours % 24

	minutes := totalMinutes % 60

	for minutes < 0 {
		hours -= 1
		minutes += 60
	}

	for hours < 0 {
		hours += 24
	}

	return Clock{hours, minutes}
}

// Add minutes to the current time in the clock.
func (c Clock) Add(m int) Clock {
	return New(c.h, c.m+m)
}

// Subtracts minutes to the current time in the clock.
func (c Clock) Subtract(m int) Clock {
	return New(c.h, c.m-m)
}

// Converts to string.
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.h, c.m)
}
