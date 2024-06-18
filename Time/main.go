package main

import (
	"fmt"
	"time"
)

func main() {
	//dd-mm-yyyy
	//"dd/mm/yyyy"
	//2006-01-02  15:04:05
	currentTime := time.Now()
	fmt.Println(currentTime)
	fmt.Printf("Type of currtime %T\n", currentTime)

	formatted := currentTime.Format("2006/01/02, 03:04 PM")
	fmt.Println(formatted)

	layout := "2006-01-02"
	dateStr := "2023-11-25"
	formatted_time, _ := time.Parse(layout, dateStr)
	fmt.Println(formatted_time)

	//add 1 more day to current Time

	dates := currentTime.Add(24 * time.Hour)
	fmt.Println("new_date time", dates)

	format_dates := dates.Format("2006/01/02")
	fmt.Println(format_dates)

	newYear := currentTime.AddDate(1, 0, 0)
	fmt.Println(newYear)
}
