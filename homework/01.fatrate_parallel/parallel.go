package main

import "fmt"

func main() {
	//定义变量
	var x1, y1, x2, y2, x3, y3, x4, y4, k1, k2, b1, b2 float64
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
	//输入的直线坐标点相同时
	if (x1 == x2 && x2 == y1 && y1 == y2) || (x3 == x4 && x4 == y3 && y3 == y4) {
		fmt.Println("输入的参数不能确定为直线，无法判断")
		//斜率都存在时
	} else if x1 != x2 && x3 != x4 {
		k1 = (y2 - y1) / (x2 - x1)
		k2 = (y4 - y3) / (x4 - x3)
		if k1 == k2 {
			//直线方程式y=kx+b
			b1 = y1 - k1*x1
			b2 = y3 - k2*x3
			//判断两条直线是否在一条直线方程上
			if b1 == b2 {
				fmt.Println("两条直线不平行，且重合")
			} else {
				fmt.Println("两条直线平行")
			}
		} else {
			fmt.Println("两条直线不平行")
		}
		//斜率都不存在时
	} else if x1 == x2 && x3 == x4 {
		//判断两条线段是否在一条直线方程上
		if x1 == x3 {
			fmt.Println("两条直线不平行,垂直于x轴且重合")
		} else {
			fmt.Println("两条直线平行,且垂直于x轴")
		}
		//其中一条斜率不存在时
	} else {
		fmt.Println("两条直线不平行")
	}
}
