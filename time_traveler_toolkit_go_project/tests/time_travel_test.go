package tests

import (
	"testing"
	"time"
	"time_traveler_toolkit/internal/time_travel"
)

func TestGetTodaysDate(t *testing.T) {
	expected := time.Now().Truncate(24 * time.Hour) // Truncate to the day for comparison
	actual := time_travel.GetTodaysDate().Truncate(24 * time.Hour)

	if !expected.Equal(actual) {
		t.Errorf("Expected date: %v, but got: %v", expected, actual)
	}
}


func TestGetLaterDateByDays(t *testing.T) {
	laterDays := 10
	today := time.Now()
	actual := time_travel.GetLaterDateByDays(today, laterDays)
	expected := today.Add(time.Hour * 24 * time.Duration(laterDays))
	if !expected.Equal(actual) {
		t.Errorf("Expected: %v, but got: %v", expected, actual)
	}
}

func TestGetPreviousDateByWeeks(t *testing.T) {
	prevWeeks := 15
	today := time.Now()
	actual := time_travel.GetPreviousDateByWeeks(today, prevWeeks)
	expected := today.AddDate(0, 0, -7*prevWeeks)
	if !expected.Equal(actual) {
		t.Errorf("Expected: %v, but got: %v", expected, actual)
	}
}

func TestGetTimeDifference(t *testing.T) {
	date1 := time.Date(2002, 10, 3, 0, 0, 0, 0, time.UTC)
	date2 := time.Date(2012, 10, 7, 0, 0, 0, 0, time.UTC)
	actual := time_travel.GetTimeDifference(date1, date2)
	expected := "Years-10:Months-0:Days-4"
	if actual != expected {
		t.Errorf("Expected: %s, but got: %s", expected, actual)
	}
}
