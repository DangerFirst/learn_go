package main

import "fmt"

func main() {
	//定义变量
	var year, month, day int
	//输入年月
	fmt.Println("输入年份")
	fmt.Scanln(&year)
	fmt.Println("输入月份")
	fmt.Scanln(&month)
	//计算第一天为星期几
	week := (5*(year-1)/4-(year-1)/100+(year-1)/400)%7 + 1
	//日历数组
	calendar := [...][7]string{
		{"一", "二", "三", "四", "五", "六", "日"},
	}
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

	fmt.Println(calendar)
	fmt.Println(day)
	fmt.Println(week)
}
