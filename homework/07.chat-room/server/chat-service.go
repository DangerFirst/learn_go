package main

import (
	"chatroom/pkg/apis"
	"golang.org/x/net/context"
	"log"
)

var _ apis.ChatServiceServer = &chatServer{}

type chatServer struct {
	messageCh chan *apis.Message
}

func (c chatServer) Register(ctx context.Context, account *apis.Account) (*apis.Account, error) {
	log.Printf("收到新注册人：%s\n", account.String())
	return account, nil
}

func (c chatServer) Login(ctx context.Context, account *apis.Account) (*apis.Null, error) {
	return nil, nil
}

func (c chatServer) OnlineUser(ctx context.Context, null *apis.Null) (*apis.OnlineUserList, error) {
	return nil, nil
}

func (c chatServer) Chat(ctx context.Context, account *apis.Account) (*apis.ChatHistory, error) {
	return nil, nil
}

func (c chatServer) ChatRecord(ctx context.Context, account *apis.Account) (*apis.ChatHistory, error) {
	return nil, nil
}

func (c chatServer) RevMessage(ctx context.Context, account *apis.Account) (*apis.Message, error) {
	return nil, nil
}
