package helpers

import "time"

func ConvertDate[Date time.Time | *time.Time, DateFormatted string | *string](date Date, layout string) DateFormatted {
	var (
		dateFormatted DateFormatted
		convertFunc   func(text string) DateFormatted
	)

	if _, isOk := any(dateFormatted).(string); isOk {
		convertFunc = func(text string) DateFormatted {
			if value, ok := any(text).(DateFormatted); ok {
				return value
			}
			return dateFormatted
		}
	} else {
		convertFunc = func(text string) DateFormatted {
			if value, ok := any(&text).(DateFormatted); ok {
				return value
			}
			return dateFormatted
		}
	}

	switch dateTime := any(date).(type) {
	case time.Time:
		return convertFunc(dateTime.Format(layout))
	case *time.Time:
		if dateTime != nil {
			return convertFunc(dateTime.Format(layout))
		}
	}

	return dateFormatted
}
