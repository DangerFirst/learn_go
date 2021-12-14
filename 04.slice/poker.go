package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	//定义牌组数组
	var initPoker, temPoker []string
	//调用初始化牌组函数
	initPoker = initPokerOrder()
	//打印初始化的牌组数组
	fmt.Println("初始牌组为：", initPoker)
	//洗牌
	//洗切牌重复五次
	for t := 0; t < 5; t++ {
		//将牌分为两叠，然后均匀洗在一起
		for site := 0; site < 27; site++ {
			temPoker = append(temPoker, initPoker[site], initPoker[site+27])
		}
		//切5次牌
		for cut := 0; cut < 5; cut++ {
			rand.Seed(time.Now().Unix())
			cutSite := rand.Intn(54)
			temPoker = append(temPoker[cutSite:], temPoker[:cutSite]...)
		}
		initPoker = temPoker
		//初始化临时数组
		temPoker = nil
	}
	fmt.Println("打乱牌组为：", initPoker)
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
	initPoker = append(initPoker, "大王", "小王")
	return initPoker
}
