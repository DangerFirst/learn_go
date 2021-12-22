package main

import (
	"fmt"
)

func guess(left, right uint) {
	guessed := (left + right) / 2
	fmt.Println("我猜的数字为：", guessed)
	fmt.Println("left and right and guessed:", left, right, guessed)
	fmt.Print("如果低了输入0，如果高了输入1，如果对了输入9：")
	var getFromInput string
	fmt.Scanln(&getFromInput)
	switch getFromInput {
	case "0":
		if left == right {
			fmt.Println("你是不是改主意了")
			return
		}
		guess(guessed+1, right)
	case "1":
		if left == right || left == guessed {
			fmt.Println("你是不是改主意了")
			return
		}
		guess(left, guessed-1)
	case "9":
		fmt.Println("你猜对了")
	}
}
