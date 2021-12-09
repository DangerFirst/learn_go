package main

import "fmt"

func main(){
	//定义数组
	scores:=map[string]int{
		"小黑":23,
		"小红":27,
		"小绿":53,
		"小蓝":63,
		"小白":73,
		"小橙":43,
		"小黄":35,
		"小明":93,
		"小张":83,
		"小赵":73,
		"小李":67,
		"小松":77,
		"小宋":78,
		"小马":87,
		"小六":73,
		"小刘":91,
		"小振":99,
		"小唐":96,
		"小楼":100,
		"小周":87,
	}
	//计算平均分
	scoreSum:=0
	for _,score:=range scores{
		scoreSum+=score
	}
	fmt.Println("平均分为：",scoreSum/len(scores))
}
