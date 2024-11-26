package helpers

import "time"

func ConvertDate[Date time.Time | *time.Time, FormattedDate string | *string](date Date, layout string) FormattedDate {
	var (
		formattedDate       FormattedDate
		formattedDateString string
	)

	switch dateTime := any(date).(type) {
	case time.Time:
		formattedDateString = dateTime.Format(layout)
	case *time.Time:
		if dateTime == nil {
			return formattedDate
		}
		formattedDateString = dateTime.Format(layout)
	}

	if _, isString := any(formattedDate).(string); isString {
		if value, ok := any(formattedDateString).(FormattedDate); ok {
			return value
		}
		return formattedDate
	}

	if value, ok := any(&formattedDateString).(FormattedDate); ok {
		return value
	}

	return formattedDate
}
