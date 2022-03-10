package main

import (
	"fmt"
	"log"
)

func main() {
	input := &inputFromStd{}
	calc := &Calc{}
	rk := &FatRateRank{}
	records := NewRecord("E:/registered_information.txt")
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

		rankResultWithBubbleSort, _ := rk.getRankWithBubbleSort(pi.Name)
		fmt.Println(pi.Name, "冒泡排序排名结果：", rankResultWithBubbleSort)

		rankResultWithQuickSort, _ := rk.getRankWithQuickSort(pi.Name)
		fmt.Println(pi.Name, "快速排序排名结果：", rankResultWithQuickSort)
	}
}
