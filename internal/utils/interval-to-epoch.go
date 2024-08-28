package utils

import (
	"daily-quote-api/internal/enums"
	"time"
)

func getStartOfWeek(date time.Time) time.Time {
	weekday := int(date.Weekday())
	startOfWeek := date.AddDate(0, 0, -((weekday + 6) % 7)) // Calculate the start of this week
	return startOfWeek
}

func getStartOfFortnight(date time.Time) time.Time {
	startOfWeek := getStartOfWeek(date)

	_, weekNumber := startOfWeek.ISOWeek()

	// If the week is even, subtract 1 week to get the start of the latest odd week
	if weekNumber%2 == 0 {
		startOfWeek = startOfWeek.AddDate(0, 0, -7)
	}

	return startOfWeek
}

func IntervalToEpoch(interval enums.TimeInterval) (int64, string) {
	var dateTime time.Time
	var unit string

	now := time.Now()

	switch interval {
	case enums.SECONDLY:
		dateTime = time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), 0, time.UTC)
		unit = "second"
	case enums.MINUTELY:
		dateTime = time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), 0, 0, time.UTC)
		unit = "minute"
	case enums.HOURLY:
		dateTime = time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), 0, 0, 0, time.UTC)
		unit = "hour"
	case enums.DAILY:
		dateTime = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)
		unit = "day"
	case enums.WEEKLY:
		startOfWeek := getStartOfWeek(now)
		dateTime = time.Date(startOfWeek.Year(), startOfWeek.Month(), startOfWeek.Day(), 0, 0, 0, 0, time.UTC)
		unit = "week"
	case enums.FORTNIGHTLY:
		startOfFortnight := getStartOfFortnight(now)
		dateTime = time.Date(startOfFortnight.Year(), startOfFortnight.Month(), startOfFortnight.Day(), 0, 0, 0, 0, time.UTC)
		unit = "fortnight"
	case enums.MONTHLY:
		dateTime = time.Date(now.Year(), now.Month(), 0, 0, 0, 0, 0, time.UTC)
		unit = "month"
	case enums.YEARLY:
		dateTime = time.Date(now.Year(), 0, 0, 0, 0, 0, 0, time.UTC)
		unit = "year"
	}

	return dateTime.Unix(), unit

}
