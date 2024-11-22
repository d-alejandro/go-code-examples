package helpers

import "time"

func ConvertDate[Date time.Time | *time.Time, DateFormatted string | *string](date Date, layout string) DateFormatted {
	var dateFormatted DateFormatted

	convertFunc := func(text string) DateFormatted {
		return any(text).(DateFormatted)
	}

	if _, ok := any(dateFormatted).(string); !ok {
		convertFunc = func(text string) DateFormatted {
			return any(&text).(DateFormatted)
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
