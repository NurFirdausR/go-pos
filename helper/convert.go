package helper

import (
	"time"
)

func ConvertTime(timeDate string) string {

	parsedTime, _ := time.Parse(time.RFC3339, timeDate)

	// Format time.Time to obtain only the date portion
	formattedDate := parsedTime.Format("2006-01-02")
	// fmt.Println("Formatted date:", formattedDate)
	return formattedDate
}
