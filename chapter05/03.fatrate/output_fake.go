package main

type fakeOutput struct {
	p Person
	s string
}

func (o *fakeOutput) OutPut(p Person, s string) {
	o.p = p
	o.s = s
}
