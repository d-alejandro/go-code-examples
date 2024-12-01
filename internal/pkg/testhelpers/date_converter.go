package testhelpers

import "time"

func ConvertDateWithSystemLocal(date time.Time) time.Time {
	return time.Date(
		date.Year(),
		date.Month(),
		date.Day(),
		date.Hour(),
		date.Minute(),
		date.Second(),
		0,
		time.Local,
	)
}
