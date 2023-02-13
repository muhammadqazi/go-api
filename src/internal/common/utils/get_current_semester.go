package utils

import (
	"strconv"
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

func GetCurrentYear() string {
	year := time.Now().Year()
	return strconv.Itoa(year)
}
