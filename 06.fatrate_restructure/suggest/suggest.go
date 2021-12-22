package suggest

import "fmt"

func Suggest(name, sex string, age int, bmi, fatRate float64) {
	fmt.Println("姓名：", name)
	fmt.Printf("BMI：%.2f\n", bmi)
	fmt.Printf("体脂率：%.2f\n", fatRate)
	fmt.Print("评估及建议：")
	if sex == "男" {
		//男性体脂建议
		suggestForMale(age, fatRate)
	} else {
		//女性体脂建议
		suggestForFemale(age, fatRate)
	}
	return
}
