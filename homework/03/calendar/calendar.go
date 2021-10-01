package calendar

import (
	"fmt"
	"time"
)

var (
	months = map[int]string{
		1:  "January",
		2:  "February",
		3:  "March",
		4:  "April",
		5:  "May",
		6:  "June",
		7:  "July",
		8:  "August",
		9:  "September",
		10: "October",
		11: "November",
		12: "December",
	}
)

func ShowCalendar(year int, month int) {
	showHeader(year, month)
	showWeekDay()
	showDay(year, month)
}

func showHeader(year int, month int) {
	monthName := months[month]
	padding := (20 - 5 - len(monthName)) / 2
	paddingStr := ""
	for i := 0; i < padding; i++ {
		paddingStr = paddingStr + " "
	}
	fmt.Printf("%s%s %d\n", paddingStr, monthName, year)
}

func showWeekDay() {
	fmt.Println("Su Mo Tu We Th Fr Sa")
}

func showDay(year int, month int) {
	// Lấy số ngày trong tháng
	daysInMonth := getDaysInMonth(year, month)

	// Lấy ngày bắt đầu của tháng
	startDay := getStartOfMonth(year, month)

	// In ngày bắt đầu của tháng
	for i := 0; i < startDay; i++ {
		fmt.Printf("%2s ", " ")
	}

	// In số ngày trong tháng
	for i := 1; i <= daysInMonth; i++ {
		fmt.Printf("%2d ", i)
		if (i + startDay) % 7 == 0 {
			fmt.Println()
		}
	}
	fmt.Println()
}

func Date(year, month, day int) time.Time {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}

// Tính toán ngày bắt đầu của tháng vào thứ mấy trong tuần (0 -> 6)
func getStartOfMonth(year, month int) int {
	t := Date(year, month, 1)
	return int(t.Weekday())
}

// Tính toán số ngày trong tháng
func getDaysInMonth(year, month int) int {
	t := Date(year, month + 1, 0)
	return t.Day()
}