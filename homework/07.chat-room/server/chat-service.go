package main

import (
	"chatroom/pkg/apis"
	"golang.org/x/net/context"
	"log"
)

var _ apis.ChatServiceServer = &chatServer{}

var ol *apis.OnlineUserList = &apis.OnlineUserList{}

type chatServer struct {
	messageCh chan *apis.Message
}

func (c chatServer) LogOut(ctx context.Context, account *apis.Account) (*apis.Account, error) {
	log.Printf("%s退出登录成功\n", account.Name)
	for i, v := range ol.OnlineUsers {
		if v.Account == account.Account {
			ol.OnlineUsers = append(ol.OnlineUsers[:i], ol.OnlineUsers[i+1:]...)
		}
	}
	return account, nil
}

func (c chatServer) Login(ctx context.Context, account *apis.Account) (*apis.Account, error) {
	log.Printf("%s登录成功\n", account.Name)
	ol.OnlineUsers = append(ol.OnlineUsers, &apis.OnlineUser{
		Account: account.Account,
		Name:    account.Name,
	})
	return account, nil
}

func (c chatServer) Register(ctx context.Context, account *apis.Account) (*apis.Account, error) {
	log.Printf("收到新注册人：%s\n", account.String())
	return account, nil
}

func (c chatServer) OnlineUser(ctx context.Context, null *apis.Null) (*apis.OnlineUserList, error) {
	log.Printf("在线用户：%s\n", ol.OnlineUsers)
	return ol, nil
}

func (c chatServer) Chat(ctx context.Context, account *apis.Account) (*apis.ChatHistory, error) {
	log.Printf("%s开始聊天\n", account.Name)
	chy := &apis.ChatHistory{}
	return chy, nil
}

func (c chatServer) ChatRecord(ctx context.Context, account *apis.Account) (*apis.ChatHistory, error) {
	return nil, nil
}

func (c chatServer) RevMessage(ctx context.Context, account *apis.Account) (*apis.Message, error) {
	return nil, nil
}
