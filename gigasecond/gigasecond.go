// Package gigasecond provides a function to to add a gigasecond to a given
// time.
package gigasecond

import "time"

// AddGigasecond adds one gigasecond to the given time and returns the
// resulting time.
func AddGigasecond(t time.Time) time.Time {
	d, err := time.ParseDuration("1000000000s")

	if err != nil {
		panic(err)
	}

	return t.Add(d)
}
