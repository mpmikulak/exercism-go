package logs

// Application identifies the application emitting the given log.
func Application(log string) string {
	for _, c := range log {
		switch {
		case c == '\u2757':
			return "recommendation"
		case c == '🔍':
			return "search"
		case c == '\u2600':
			return "weather"
		}
	}
	return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	var newLog string
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
	var count = 0
	for _, _ = range log {
		count++
	}
	return count <= limit
}
