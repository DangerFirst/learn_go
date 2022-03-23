package main

import (
	"crypto/rand"
	"learn.go/pkg/apis"
	"math/big"
)

type Interface interface {
	ReadPersonalInformation() (apis.PersonalInformation, error)
}

var _ Interface = &fakeInterface{}

type fakeInterface struct {
	name       string
	sex        string
	baseWeight float64
	baseTall   float64
	baseAge    int
}

func (f fakeInterface) ReadPersonalInformation() (apis.PersonalInformation, error) {
	r, _ := rand.Int(rand.Reader, big.NewInt(1000))
	out := float64(r.Int64()) / 1000
	if r.Int64()%2 == 0 {
		out = 0 - out
	}
	pi := apis.PersonalInformation{
		Name:   f.name,
		Sex:    f.sex,
		Tall:   f.baseTall,
		Weight: f.baseWeight,
		Age:    f.baseAge,
	}
	f.baseWeight += out
	return pi, nil
}
