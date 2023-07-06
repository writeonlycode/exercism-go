package logs

// Application identifies the application emitting the given log.
func Application(log string) string {
	for _, c := range log {
		switch c {
		case '❗':
			return "recommendation"
		case '🔍':
			return "search"
		case '☀':
			return "weather"
		}
	}

	return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	newLog := ""

	for _, c := range log {
		if c == oldRune {
			newLog += string(newRune)
		} else {
			newLog += string(c)
		}
	}

	return newLog
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	l := 0

	for range log {
		l++
	}

	return l <= limit
}
