package main

import "fmt"

func main() {
	vg := &voteGame{
		students: []*student{
			&student{name: "1"},
			&student{name: "2"},
			&student{name: "3"},
			&student{name: "4"},
		}}
	leader := vg.goRun()
	fmt.Println(leader)
	leader.Distribute()

	var stdXS = &student{
		name: "小宋",
	}
	var ldXS Leader = Leader(*stdXS)
	fmt.Println(ldXS)
}

type voteGame struct {
	students []*student
}

type Leader student

func (l *Leader) Distribute() {
	fmt.Println("发作业了。。。。")
}

func (g *voteGame) goRun() *Leader {
	for _, item := range g.students {
		item.voteA(g.students[0])
	}
	maxScore := -1
	maxScoreIndex := -1
	for i, item := range g.students {
		if maxScore < item.agree {
			maxScore = item.agree
			maxScoreIndex = i
		}
	}
	if maxScore < 0 {
		return nil
	}
	return (*Leader)(g.students[maxScoreIndex])
}

type student struct {
	name            string
	agree, disagree int
}

func (std *student) voteA(target *student) {
	target.agree++
}

func (std *student) voteD(target *student) {
	target.disagree++
}
