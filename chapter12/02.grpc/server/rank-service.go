package main

import (
	"context"
	"io"
	"learn.go/chapter12/02.grpc/server/frinterface"
	"learn.go/pkg/apis"
	"log"
	"sync"
)

var _ apis.RankServiceServer = &rankServer{}

type rankServer struct {
	sync.Mutex
	apis.UnimplementedRankServiceServer
	rankS    frinterface.ServerInterface
	personCh chan *apis.PersonalInformation
}

func (r rankServer) Update(ctx context.Context, information *apis.PersonalInformation) (*apis.PersonalInformationFatRate, error) {
	r.regPerson(information)
	return r.rankS.UpdatePersonalInformation(information)
}

func (r rankServer) GetFR(ctx context.Context, information *apis.PersonalInformation) (*apis.PersonalRank, error) {
	return r.rankS.GetFatRate(information.Name)
}

func (r rankServer) GetTop(ctx context.Context, null *apis.Null) (*apis.PersonalRanks, error) {
	top, err := r.rankS.GetTop()
	if err != nil {
		log.Println("获取榜单出错：", err)
		return nil, err
	}
	return &apis.PersonalRanks{PersonalRanks: top}, nil
}

func (r rankServer) regPerson(pi *apis.PersonalInformation) {
	r.rankS.RegisterPersonalInformation(pi) //todo handle error
	r.personCh <- pi
}

func (r rankServer) WatchPersons(null *apis.Null, server apis.RankService_WatchPersonsServer) error {
	for pi := range r.personCh {
		if err := server.Send(pi); err != nil {
			log.Println("发送失败，结束：", err)
			return err
		}
	}
	return nil
}

func (r rankServer) RegisterPersons(server apis.RankService_RegisterPersonsServer) error {
	pis := &apis.PersonalInformationList{}
	for {
		pi, err := server.Recv()
		if err == io.EOF {
			break
		}
		if err != nil && err != io.EOF {
			log.Printf("WARNING：获取人员注册时失败：%v\n", err)
			return err
		}
		pis.Items = append(pis.Items, pi)
		r.regPerson(pi)
	}
	log.Println("连续得到注册清单：", pis.String())
	return server.SendAndClose(pis)
}

func (r rankServer) Register(ctx context.Context, information *apis.PersonalInformation) (*apis.PersonalInformation, error) {
	r.regPerson(information)
	log.Printf("收到新注册人：%s\n", information.String())
	return information, nil
}
