package suggest

import "fmt"

func suggestForFemale(age int, fatRate float64) {
	if age >= 18 && age < 40 {
		if fatRate < 0 {
			fmt.Println("异常，请检查输入内容重新输入")
		} else if fatRate >= 0 && fatRate < 0.2 {
			fmt.Println("偏瘦，需要多吃点")
		} else if fatRate < 0.27 {
			fmt.Println("标准，不错，保持下去")
		} else if fatRate < 0.34 {
			fmt.Println("偏重，注意饮食，加强锻炼")
		} else if fatRate < 0.39 {
			fmt.Println("肥胖，控制饮食，多多锻炼")
		} else {
			fmt.Println("严重肥胖，少吃多动，坚持下去")
		}
	} else if age >= 40 && age < 60 {
		if fatRate < 0 {
			fmt.Println("异常，请检查输入内容重新输入")
		} else if fatRate >= 0 && fatRate < 0.21 {
			fmt.Println("偏瘦，需要多吃点")
		} else if fatRate < 0.28 {
			fmt.Println("标准，不错，保持下去")
		} else if fatRate < 0.35 {
			fmt.Println("偏重，注意饮食，加强锻炼")
		} else if fatRate < 0.4 {
			fmt.Println("肥胖，控制饮食，多多锻炼")
		} else {
			fmt.Println("严重肥胖，少吃多动，坚持下去")
		}
	} else {
		if fatRate < 0 {
			fmt.Println("异常，请检查输入内容重新输入")
		} else if fatRate >= 0 && fatRate < 0.22 {
			fmt.Println("偏瘦，需要多吃点")
		} else if fatRate < 0.29 {
			fmt.Println("标准，不错，保持下去")
		} else if fatRate < 0.36 {
			fmt.Println("偏重，注意饮食，加强锻炼")
		} else if fatRate < 0.41 {
			fmt.Println("肥胖，控制饮食，多多锻炼")
		} else {
			fmt.Println("严重肥胖，少吃多动，坚持下去")
		}
	}
}
