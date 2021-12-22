package suggest

import "fmt"

func Suggest(names, sexs [999]string, ages [999]int, num int, bmis, fatRates [999]float64) {
	for j := 0; j < num; j++ {
		fmt.Println("姓名：", names[j])
		fmt.Printf("BMI：%.2f\n", bmis[j])
		fmt.Printf("体脂率：%.2f\n", fatRates[j])
		fmt.Print("评估及建议：")
		if sexs[j] == "男" {
			//男性体脂建议
			suggestForMale(ages[j], fatRates[j])
		} else {
			//女性体脂建议
			suggestForFemale(ages[j], fatRates[j])
		}

	}
	return
}
