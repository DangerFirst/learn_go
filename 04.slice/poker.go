package main

import (
	"fmt"
	"strconv"
)

func main() {
	//定义牌组数组
	initPoker := []string{}
	//调用初始化牌组函数
	initPoker=initPokerOrder()
	//打印初始化的牌组数组
	fmt.Println("初始牌组为：",initPoker)
	//洗牌

}

//初始化牌组
func initPokerOrder() (initPoker []string) {
	pokerColor := []string{"红心", "黑桃", "方块", "梅花"}
	for size := 1; size <= 13; size++ {
		for _, color := range pokerColor {
			switch size {
			case 11:
				initPoker = append(initPoker, color+"J")
			case 12:
				initPoker = append(initPoker, color+"Q")
			case 13:
				initPoker = append(initPoker, color+"K")
			default:
				initPoker = append(initPoker, color+strconv.Itoa(size))
			}
		}
	}
	initPoker = append(initPoker, "大王","小王")
	return initPoker
}
