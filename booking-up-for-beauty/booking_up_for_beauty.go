package booking

import (
	"fmt"
	"time"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	t, err := time.Parse("1/02/2006 15:04:05", date)
	if err != nil {
		fmt.Println(err)
	}
	return t
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	var timeNow = time.Now()
	scheduleDate, err := time.Parse("January 2, 2006 15:04:05", date)
	if err != nil {
		fmt.Println(err)
	}
	if timeNow.After(scheduleDate) {
		return true
	}
	return false
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	var appTime, err = time.Parse("Monday, January 2, 2006 15:04:05 ", date)
	if err != nil {
		fmt.Println(err)
	}
	if appTime.Hour() >= 12 && appTime.Hour() < 18 {
		return true
	}
	return false
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	appTime, err := time.Parse("1/2/2006 15:04:05", date)
	if err != nil {
		fmt.Println(err)
	}
	return fmt.Sprintf("You have an appointment on %s.", appTime.Format("Monday, January 2, 2006, at 15:04"))
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	date, err := time.Parse("2006-01-02 15:04:05", "2025-09-15 00:00:00")
	if err != nil {
		fmt.Println(err)
	}
	return date
}
