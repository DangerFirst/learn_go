package main

import (
	"fmt"
	"github.com/spf13/cobra"
)

func main() {
	var (
		name   string
		sex    string
		age    int
		tall   float64
		weight float64
	)
	cmd:=&cobra.Command{
		Use:                        "",
		Short:                      "",
		Long:                       "",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("name：", name)
			fmt.Println("sex：", sex)
			fmt.Println("age：", age)
			fmt.Println("tall：", tall)
			fmt.Println("weight：", weight)
		},
	}
	cmd.Flags().StringVar(&name,"name1","","姓名")
	cmd.Flags().StringVar(&sex,"sex","","性别")
	cmd.Flags().IntVar(&age,"age",0,"年龄")
	cmd.Flags().Float64Var(&tall,"tall",0,"身高")
	cmd.Flags().Float64Var(&weight,"weight",0,"体重")
	cmd.Execute()
}
