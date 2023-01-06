package utils

import (
	"fmt"
	"strconv"
	"time"
)

func GenerateStudentNumber(universityCode int, semester string, lastSID uint) uint {
	currentYear, _, _ := time.Now().Date()
	lastTwoDigitsOfYear := currentYear % 100

	var semesterCode int
	switch semester {
	case "Fall":
		semesterCode = 0
	case "Spring":
		semesterCode = 1
	case "Summer":
		semesterCode = 3
	default:
		semesterCode = 0
	}

	lastFourDigits := lastSID + 1

	studentNumber, _ := strconv.ParseUint(fmt.Sprintf("%d%02d%d%04d", universityCode, lastTwoDigitsOfYear, semesterCode, lastFourDigits), 10, 32)
	return uint(studentNumber)
}
