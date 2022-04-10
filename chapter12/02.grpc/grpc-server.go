package main

import (
	"context"
	context2 "golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
	"learn.go/pkg/apis"
	"log"
	"net"
)

func startGRPCServer(ctx context.Context) {
	lis, err := net.Listen("tcp", nodePort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer([]grpc.ServerOption{}...)
	apis.RegisterRankServerServer(s, &rankServer{
		persons: map[string]*apis.PersonalInformation{},
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
