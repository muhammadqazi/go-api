package utils

import "time"

func GetCurrentSemester() string {

	_, month, _ := time.Now().Date()
	var semester string
	switch {
	case month >= time.September && month <= time.January:
		semester = "Fall"
	case month >= time.February && month <= time.June:
		semester = "Spring"
	default:
		semester = "Summer"
	}

	return semester
}
