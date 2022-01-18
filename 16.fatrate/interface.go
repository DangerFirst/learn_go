package main

type InputService interface {
	GetInput() Person
}

type OutPutService interface {
	OutPut(Person, string)
}
