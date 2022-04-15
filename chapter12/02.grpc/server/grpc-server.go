package main

import (
	"context"
	"learn.go/chapter12/02.grpc/server/ranks"

	//context2 "golang.org/x/net/context"
	"google.golang.org/grpc"
	"learn.go/pkg/apis"
	"log"
	"net"
)

func main() {
	ctx, cancel := context.WithCancel(context.TODO())
	defer cancel()
	startGRPCServer(ctx)
}

func startGRPCServer(ctx context.Context) {
	lis, err := net.Listen("tcp", "0.0.0.0:9090")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer([]grpc.ServerOption{}...)
	apis.RegisterRankServiceServer(s, &rankServer{
		rankS:    ranks.NewFatRateRank(),
		personCh: make(chan *apis.PersonalInformation, 1024),
	})
	go func() {
		select {
		case <-ctx.Done():
			s.Stop()
		}
	}()
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
