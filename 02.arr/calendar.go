package main

import "fmt"

func main() {
	//定义变量
	var year, month, day, dayAdd int
	//输入年月
	fmt.Println("输入年份")
	fmt.Scanln(&year)
	fmt.Println("输入月份")
	fmt.Scanln(&month)
	//计算一年的第一天为星期几
	YearFirstDay := (5*(year-1)/4-(year-1)/100+(year-1)/400)%7 + 1
	//日历数组
	calendar := [6][7]int{}
	//判断月份天数
	switch month {
	case 1, 3, 5, 7, 8, 10, 12:
		day = 31
	case 4, 6, 9, 11:
		day = 30
	case 2:
		if (year%4 == 0 && year%100 != 0) || (year%400 == 0) {
			day = 29
		} else {
			day = 28
		}
	}
	//计算月份第一天为星期几
	daySum := 0
	for m := 1; m < month; m++ {
		switch m {
		case 1, 3, 5, 7, 8, 10, 12:
			dayAdd = 31
		case 4, 6, 9, 11:
			dayAdd = 30
		case 2:
			if (year%4 == 0 && year%100 != 0) || (year%400 == 0) {
				dayAdd = 29
			} else {
				dayAdd = 28
			}
		}
		daySum += dayAdd
	}
	weekFirstDay := (daySum + YearFirstDay) % 7

	//写入输入月份的日期
	dayIs := 1
	for j := 0; dayIs <= day; j++ {
		if j == 0 {
			for i := weekFirstDay - 1; i < 7; i++ {
				calendar[j][i] = dayIs
				dayIs++
			}
		} else {
			for i := 0; i < 7 && dayIs <= day; i++ {
				calendar[j][i] = dayIs
				dayIs++
			}
		}
	}

	sit := " "
	//calendarHeader := []string{"一", "二", "三", "四", "五", "六", "日"}
	calendarHeader := []string{"Mo", "Tu", "We", "Th", "Fr", "Sa", "Su"}
	//输出该月日期
	for h := 0; h < 7; h++ {
		fmt.Printf("%3s ", calendarHeader[h])
	}
	fmt.Println()
	for w := 0; w < 6; w++ {
		for d := 0; d < 7; d++ {
			if calendar[w][d] == 0 {
				fmt.Printf("%3s ", sit)
			} else {
				fmt.Printf("%3d ", calendar[w][d])
			}
		}
		fmt.Println()
	}
}
