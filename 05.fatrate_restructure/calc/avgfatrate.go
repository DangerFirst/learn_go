package calc

import "fmt"

func AvgFatRate(num int, fatRateSum float64) {
	fmt.Printf("总人数：%d，平均体脂率 %.2f\n", num, fatRateSum/float64(num))
}