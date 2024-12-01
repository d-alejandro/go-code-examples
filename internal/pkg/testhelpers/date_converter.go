package testhelpers

import (
	"time"

	"github.com/d-alejandro/go-code-examples/internal/config"
)

func ConvertDateWithSystemLocal(date time.Time) time.Time {
	return time.Date(
		date.Year(),
		date.Month(),
		date.Day(),
		date.Hour(),
		date.Minute(),
		date.Second(),
		config.DateNSECMin,
		time.Local,
	)
}
