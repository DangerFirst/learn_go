package main

import (
	"encoding/json"
	"fmt"
	"github.com/armstrongli/go-bmi"
	"io/ioutil"
	"learn.go/pkg/apis"
	"log"
	"os"
)

func main() {
	input := &inputFromStd{}
	calc := &Calc{}
	rk := &FatRateRank{}
	records := NewRecord("E:/go语言/go/go_file/records.txt")
	for {
		pi := input.GetInput()
		if err := records.savePersonalInformation(pi); err != nil {
			log.Fatal("保存失败：", err)
		}
		fr, err := calc.FatRate(pi)
		if err != nil {
			log.Fatal("计算体脂率失败：", err)
		}
		rk.inputRecord(pi.Name, fr)

		rankResult, _ := rk.getRank(pi.Name)
		fmt.Println("排名结果：", rankResult)
	}
}

func caseStudy() {
	filePath := "E:/go语言/go/go_file/荣霸天.self.information.json"

	personalInformation := apis.PersonalInformation{
		Name:   `"小"黑"`,
		Sex:    "男",
		Age:    19,
		Tall:   1.77,
		Weight: 65,
	}
	fmt.Printf("%+v\n", personalInformation)

	data, err := json.Marshal(personalInformation)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("marshal的(原生)结果是：", data)
	fmt.Println("marshal的(string)结果是：", string(data))

	writeFile(filePath, data)
	readFile(filePath)
}

func readFile(filePath string) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("读取文件失败：", err)
		return
	}
	fmt.Println("读取出来的内容是：", string(data))

	personalInformation := apis.PersonalInformation{}
	json.Unmarshal(data, &personalInformation)

	fmt.Println("开始计算体脂信息：", personalInformation)
	bmi, _ := gobmi.BMI(personalInformation.Weight, personalInformation.Tall) //todo handle error
	fmt.Printf("%s的BMI是：%v\n", personalInformation.Name, bmi)
	fatRate, _ := gobmi.Fatrate(personalInformation.Age, personalInformation.Sex, bmi)
	fmt.Printf("%s的体脂率是：%v\n", personalInformation.Name, fatRate)
}

func writeFile(filePath string, date []byte) {
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("无法打开文件", filePath, "错误信息为：", err)
		os.Exit(1)
	}
	defer file.Close()
	_, err = file.Write(date)
	fmt.Println(err)
	//handle error
}
