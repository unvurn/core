package datetime

import "time"

func Between(start time.Time, end *time.Time, t time.Time) bool {
	if start.After(t) {
		return false
	}

	if end == nil {
		return true
	}

	return end.After(t)
}

func BetweenOrEqual(start time.Time, end *time.Time, t time.Time) bool {
	return start.Equal(t) || (end != nil && end.Equal(t)) || Between(start, end, t)
}
