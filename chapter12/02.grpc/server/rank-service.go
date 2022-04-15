package main

import (
	"context"
	"learn.go/pkg/apis"
	"log"
	"sync"
)

var _ apis.RankServiceServer = &rankServer{}

type rankServer struct {
	sync.Mutex
	apis.UnimplementedRankServiceServer
	persons map[string]*apis.PersonalInformation
}

func (r rankServer) Register(ctx context.Context, information *apis.PersonalInformation) (*apis.PersonalInformation, error) {
	r.Lock()
	defer r.Unlock()
	r.persons[information.Name] = information
	log.Printf("收到新注册人：%s\n", information.String())
	return information, nil
}
