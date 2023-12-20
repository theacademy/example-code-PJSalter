package main

import (
	"fmt"
	"time"
	"time_traveler_toolkit/internal/time_travel"
)

func main() {
	fmt.Println("Today's date is:", time_travel.GetTodaysDate())

	today := time_travel.GetTodaysDate()
	fmt.Println("14 days from today is:", time_travel.GetLaterDateByDays(today, 14))

	fmt.Println("5 weeks prior to today was:", time_travel.GetPreviousDateByWeeks(today, 5))

	date1 := time.Date(2002, 10, 3, 0, 0, 0, 0, time.UTC)
	date2 := time.Date(2012, 10, 7, 0, 0, 0, 0, time.UTC)
	fmt.Println("The difference between date1 and date2 is:", time_travel.GetTimeDifference(date1, date2))
}
