package calendar

import (
	"fmt"
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
}

/*
Chiều rộng calendar = 6 * 3 + 2 = 20 ký tự
space + năm chiếm 4 ký tự
   September 2021
Su Mo Tu We Th Fr Sa
          1  2  3  4
 5  6  7  8  9 10 11
12 13 14 15 16 17 18
19 20 21 22 23 24 25
26 27 28 29 30
*/
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