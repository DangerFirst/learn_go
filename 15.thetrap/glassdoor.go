package main

import "fmt"

var _ Door = &GlassDoor{}

type GlassDoor struct{}

func (*GlassDoor) Unlock() {
	fmt.Println("GlassDoor Unlock")
}

func (*GlassDoor) Lock() {
	fmt.Println("GlassDoor Lock")
}

func (*GlassDoor) Open() {
	fmt.Println("GlassDoor Open")
}

func (*GlassDoor) Close() {
	fmt.Println("GlassDoor Close")
}
