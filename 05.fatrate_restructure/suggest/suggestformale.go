package suggest

import "fmt"

func suggestForMale(age int, fatRate float64) {
	if age >= 18 && age < 40 {
		if fatRate < 0 {
			fmt.Println("异常，请检查输入内容重新输入")
		} else if fatRate >= 0 && fatRate < 0.1 {
			fmt.Println("偏瘦，需要多吃点")
		} else if fatRate < 0.16 {
			fmt.Println("标准，不错，保持下去")
		} else if fatRate < 0.21 {
			fmt.Println("偏重，注意饮食，加强锻炼")
		} else if fatRate < 0.26 {
			fmt.Println("肥胖，控制饮食，多多锻炼")
		} else {
			fmt.Println("严重肥胖，少吃多动，坚持下去")
		}
	} else if age >= 40 && age < 60 {
		if fatRate < 0 {
			fmt.Println("异常，请检查输入内容重新输入")
		} else if fatRate >= 0 && fatRate < 0.11 {
			fmt.Println("偏瘦，需要多吃点")
		} else if fatRate < 0.17 {
			fmt.Println("标准，不错，保持下去")
		} else if fatRate < 0.22 {
			fmt.Println("偏重，注意饮食，加强锻炼")
		} else if fatRate < 0.27 {
			fmt.Println("肥胖，控制饮食，多多锻炼")
		} else {
			fmt.Println("严重肥胖，少吃多动，坚持下去")
		}
	} else {
		if fatRate < 0 {
			fmt.Println("异常，请检查输入内容重新输入")
		} else if fatRate >= 0 && fatRate < 0.13 {
			fmt.Println("偏瘦，需要多吃点")
		} else if fatRate < 0.19 {
			fmt.Println("标准，不错，保持下去")
		} else if fatRate < 0.24 {
			fmt.Println("偏重，注意饮食，加强锻炼")
		} else if fatRate < 0.29 {
			fmt.Println("肥胖，控制饮食，多多锻炼")
		} else {
			fmt.Println("严重肥胖，少吃多动，坚持下去")
		}
	}
}
