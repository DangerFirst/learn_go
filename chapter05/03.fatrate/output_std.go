package main

import "fmt"

type StdOut struct {
}

func (*StdOut) OutPut(p Person, s string) {
	fmt.Println(s)
}
