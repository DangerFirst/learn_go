package main

import (
	"fmt"
)

func main() {
	//定义map
	scores := map[string]int{
		"小黑": 23,
		"小红": 27,
		"小绿": 53,
		"小蓝": 63,
		"小白": 73,
		"小橙": 43,
		"小黄": 35,
		"小明": 93,
		"小张": 83,
		"小赵": 73,
		"小李": 67,
		"小松": 77,
		"小宋": 78,
		"小马": 87,
		"小六": 73,
		"小刘": 91,
		"小振": 99,
		"小唐": 96,
		"小楼": 100,
		"小周": 87,
	}
	fmt.Println("各人员分数情况(未排序前):", scores)
	//计算平均分
	scoreSum := 0
	for _, score := range scores {
		scoreSum += score
	}
	fmt.Println("平均分为：", scoreSum/len(scores))
	//定义名字和分数的切片
	var name []string
	var score []int
	//将map分为两个切片
	for k, v := range scores {
		name = append(name, k)
		score = append(score, v)
	}
	//按分数从高到低排序
	for i := 0; i < len(score)-1; i++ {
		for j := i + 1; j < len(score); j++ {
			if score[i] < score[j] {
				//名字跟随分数同时排序
				temp := score[i]
				temp2 := name[i]
				score[i] = score[j]
				name[i] = name[j]
				score[j] = temp
				name[j] = temp2
			} else {
				continue
			}
		}
	}
	fmt.Println("排名结果如下")
	same := 0
	//输出排名
	for r := 0; r < len(name)-1; r++ {
		//同分数判断
		if score[r] == score[r+1] {
			fmt.Printf("%d %s %d ", r+1+same, name[r], score[r])
			same--
		} else {
			fmt.Println(r+1+same, name[r], score[r])
		}
	}
	fmt.Println(len(name)+same, name[len(name)-1], score[len(score)-1])
}
