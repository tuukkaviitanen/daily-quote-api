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

func UnitOfTimeToEpoch(interval enums.UnitOfTime) int64 {
	var dateTime time.Time

	now := time.Now()

	switch interval {
	case enums.SECOND:
		dateTime = time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), 0, time.UTC)
	case enums.MINUTE:
		dateTime = time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), 0, 0, time.UTC)
	case enums.HOUR:
		dateTime = time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), 0, 0, 0, time.UTC)
	case enums.DAY:
		dateTime = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)
	case enums.WEEK:
		startOfWeek := getStartOfWeek(now)
		dateTime = time.Date(startOfWeek.Year(), startOfWeek.Month(), startOfWeek.Day(), 0, 0, 0, 0, time.UTC)
	case enums.FORTNIGHT:
		startOfFortnight := getStartOfFortnight(now)
		dateTime = time.Date(startOfFortnight.Year(), startOfFortnight.Month(), startOfFortnight.Day(), 0, 0, 0, 0, time.UTC)
	case enums.MONTH:
		dateTime = time.Date(now.Year(), now.Month(), 0, 0, 0, 0, 0, time.UTC)
	case enums.YEAR:
		dateTime = time.Date(now.Year(), 0, 0, 0, 0, 0, 0, time.UTC)
	}

	return dateTime.Unix()

}
