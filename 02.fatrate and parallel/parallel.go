package main

import "fmt"

func main() {
	//定义变量
	var x1, y1, x2, y2, x3, y3, x4, y4, k1, k2 float64
	//输入参数
	fmt.Print("输入直线一其中一端点x轴方向的坐标：")
	fmt.Scanln(&x1)
	fmt.Print("输入直线一其中一端点y轴方向的坐标：")
	fmt.Scanln(&y1)
	fmt.Print("输入直线一另一端点x轴方向的坐标：")
	fmt.Scanln(&x2)
	fmt.Print("输入直线一另一端点y轴方向的坐标：")
	fmt.Scanln(&y2)
	fmt.Print("输入直线二其中一端点x轴方向的坐标：")
	fmt.Scanln(&x3)
	fmt.Print("输入直线二其中一端点y轴方向的坐标：")
	fmt.Scanln(&y3)
	fmt.Print("输入直线二另一端点x轴方向的坐标：")
	fmt.Scanln(&x4)
	fmt.Print("输入直线二另一端点y轴方向的坐标：")
	fmt.Scanln(&y4)
	fmt.Printf("第一条直线的两点坐标为：(%.f,%.f),(%.f,%.f)\n", x1, y1, x2, y2)
	fmt.Printf("第二条直线的两点坐标为：(%.f,%.f),(%.f,%.f)\n", x3, y3, x4, y4)
	//判断直线是否平行
	//斜率都存在时
	if x1 != x2 && x3 != x4 {
		k1 = (y2 - y1) / (x2 - x1)
		k2 = (y4 - y3) / (x4 - x3)
		if k1 == k2 {
			fmt.Println("两条直线平行")
		} else {
			fmt.Println("两条直线不平行")
		}
		//斜率都不存在时
	} else if x1 == x2 && x3 == x4 {
		fmt.Println("两条直线平行")
		//其中一条斜率不存在时
	} else {
		fmt.Println("两条直线不平行")
	}
}
