package suggest

import (
	"fatrate/input"
	"fmt"
)

func Suggest(person *input.Person) {
	fmt.Println("姓名：", person.Name)
	fmt.Printf("BMI：%.2f\n", person.Bmi)
	fmt.Printf("体脂率：%.2f\n", person.FatRate)
	fmt.Print("评估及建议：")
	//男性体脂建议
	suggestForMale(person.Age, person.FatRate)
	if person.Sex == "女" {
		//女性体脂建议
		suggestForFemale(person.Age, person.FatRate)
	}
	return
}
