package frinterface

import (
	"moments/pkg/apis"
)

type ServeInterface interface {
	RegisterPersonalInformation(pi *apis.PersonalInformation) error
	UpdatePersonalInformation(pi *apis.PersonalInformation) (*apis.PersonalInformationFatRate, error)
	GetFatRate(name string) (*apis.PersonalRank, error)
	GetTop() ([]*apis.PersonalRank, error)
}

type MomentsServerInterface interface {
	Publish(ps *apis.PersonalShow) error
	Delete(ps *apis.PersonalShow) error
	Moments() ([]*apis.GoMoments, error)
}

type RankInitInterface interface {
	Init() error
}
