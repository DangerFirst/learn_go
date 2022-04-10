package main

import (
	"context"
	"learn.go/pkg/apis"
	"sync"
)

var _ apis.RankServerServer = &rankServer{}

type rankServer struct {
	sync.Mutex
	persons map[string]*apis.PersonalInformation
}

func (r *rankServer) mustEmbedUnimplementedRankServerServer() {
	return
}

func (r *rankServer) Register(ctx context.Context, information *apis.PersonalInformation) (*apis.PersonalInformation, error) {
	r.Lock()
	defer r.Unlock()
	r.persons[information.Name] = information
	return information, nil
}
