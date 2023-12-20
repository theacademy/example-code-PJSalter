package time_travel

import (
	"fmt"
	"time"
)

func GetTodaysDate() time.Time {
	return time.Now()
}

func GetLaterDateByDays(date time.Time, x int) time.Time {
	return date.Add(time.Hour * 24 * time.Duration(x))
}

func GetPreviousDateByWeeks(date time.Time, x int) time.Time {
	return date.AddDate(0, 0, -7*x)
}

func GetTimeDifference(date1, date2 time.Time) string {
	yearDiff := date2.Year() - date1.Year()
	monthDiff := int(date2.Month()) - int(date1.Month())
	dayDiff := date2.Day() - date1.Day()

	if dayDiff < 0 {
		borrow := date1.AddDate(0, 1, 0).Day()
		dayDiff = (date2.Day() + borrow) - date1.Day()
		monthDiff--
	}

	if monthDiff < 0 {
		monthDiff += 12
		yearDiff--
	}

	timeDiff := fmt.Sprintf("Years-%d:Months-%d:Days-%d", yearDiff, monthDiff, dayDiff)
	return timeDiff
}
