package main

import (
    "fmt"
    "time"
)

func main() {
    currentTime := time.Now()																				// Curr date and time
	formattedTime := currentTime.Format("2006-01-02 15:04:05")												// Formatted
	futureTime := currentTime.Add(24 * time.Hour + 60 * time.Minute + 15 * time.Second)						// Add Time
	pastTime := currentTime.Add(-24 * time.Hour)															// Subtract Time
    dateStr := "2023-09-13T10:00:00Z"																		// Date string, need to be parsed in to date-time format
	parsedTime, _ := time.Parse(time.RFC3339, dateStr)														// time.Parse(format, string) method
	nyTime, _ := time.LoadLocation("America/New_York")														// time of any location 
	newyorkTime := currentTime.In(nyTime)
	
	fmt.Println("Current Time:", currentTime)
	fmt.Println("Formatted Time:", formattedTime)
	fmt.Println("Future Time:", futureTime)
	fmt.Println("Past Time:", pastTime)
	fmt.Println("Parsed Time:",parsedTime)
	fmt.Println("Time in New York:", newyorkTime)
}
