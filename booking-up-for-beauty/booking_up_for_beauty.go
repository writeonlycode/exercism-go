package booking

import (
	"fmt"
	"time"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	time, _ := time.Parse("1/2/2006 15:04:05", date)
	return time
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	t, _ := time.Parse("January 1, 2006 15:04:05", date)
	return t.Before(time.Now())
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	t, _ := time.Parse("Monday, January 1, 2006 15:04:05", date)
	return t.Hour() > 12 || (t.Hour() == 12 && (t.Minute() > 0 || t.Second() > 0 || t.Nanosecond() > 0))
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	return Schedule(date).Format("You have an appointment on Monday, January 2, 2006, at 15:04.")
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	t, _ := time.Parse("2006-1-2", fmt.Sprintf("%d-09-15", time.Now().Year()))
	return t
}
