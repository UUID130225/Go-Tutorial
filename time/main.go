package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Printf("Exercises to familiarize yourself with Time\n")

	// Lấy thời gian hiện tại
	currentTime := time.Now()
	fmt.Println("Current time:", currentTime)

	// Định dạng thời gian
	formattedTime := currentTime.Format("2006-01-02")
	fmt.Println("Formatted time:", formattedTime)

	// Chuyển chuỗi thành thời gian
	timeString := "2024-06-01"
	parsedTime, err := time.Parse("2006-01-02", timeString)
	if err != nil {
		fmt.Println("Error parsing time:", err)
	}
	fmt.Println("Parsed time:", parsedTime)

	// Cộng thêm 2 giờ
	futureTime := currentTime.Add(2 * time.Hour)
	fmt.Println("Future time (after 2 hours):", futureTime)

	// Trừ đi 30 phút
	pastTime := currentTime.Add(-30 * time.Minute)
	fmt.Println("Past time (30 minutes ago):", pastTime)
}