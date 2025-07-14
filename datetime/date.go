package datetime

import "time"

func EqualsDate(d time.Time, d0 time.Time) bool {
	return d.Year() == d0.Year() && d.Month() == d0.Month() && d.Day() == d0.Day()
}
