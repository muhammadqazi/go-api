package utils

import (
	"time"
)

func GetCurrentSemester() string {

	_, month, _ := time.Now().Date()
	var semester string
	switch {
	case month >= time.September && month <= time.January:
		semester = "fall"
	case month >= time.February && month <= time.June:
		semester = "spring"
	default:
		semester = "summer"
	}

	return semester
}

func GetCurrentYear() int {
	return time.Now().Year()
}
