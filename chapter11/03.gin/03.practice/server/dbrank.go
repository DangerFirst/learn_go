package main

import (
	"learn.go/chapter11/02.practice/frinterface"
	"learn.go/pkg/apis"
)

var _ frinterface.ServerInterface = &dbRank{}

type dbRank struct {
}

func (d dbRank) RegisterPersonalInformation(pi *apis.PersonalInformation) error {
	panic("implement me")
}

func (d dbRank) UpdatePersonalInformation(pi *apis.PersonalInformation) (*apis.PersonalInformationFatRate, error) {
	panic("implement me")
}

func (d dbRank) GetFatRate(name string) (*apis.PersonalRank, error) {
	panic("implement me")
}

func (d dbRank) GetTop() ([]*apis.PersonalRank, error) {
	panic("implement me")
}
