package main

import (
	"chatroom/pkg/apis"
	"fmt"
	"golang.org/x/net/context"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var _ apis.ChatServiceServer = &chatServer{}

var ol *apis.OnlineUserList = &apis.OnlineUserList{}

type chatServer struct {
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
	for _, v := range ol.OnlineUsers {
		if v.Account == account.Account {
			return account, nil
		}
	}
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
	mx := &apis.ChatHistory{}
	return mx, nil
}

func (c chatServer) ChatRecord(ctx context.Context, account *apis.Account) (*apis.ChatHistory, error) {
	log.Printf("查询%s聊天记录\n", account.Name)
	chy := &apis.ChatHistory{}
	return chy, nil
}

func (c chatServer) RevMessage(ctx context.Context, account *apis.Account) (*apis.ChatHistory, error) {
	log.Printf("%s接收消息\n", account.Name)
	chy := &apis.ChatHistory{}
	return chy, nil
}

func dbCharHistory() {
	conn, err := gorm.Open(mysql.Open("root:123qwe@tcp(localhost)/testdb"))
	if err != nil {
		log.Fatal("数据库连接失败：", err)
	}
	log.Println("连接数据库成功")
	var tables []*apis.ChatHistory
	resp := conn.Find(&tables)
	if err := resp.Error; err != nil {
		fmt.Println("查询聊天记录失败：", err)
	}
	records := NewRecord("E:/CharHistory.json")
	if err := records.saveAccount(tables); err != nil {
		log.Fatal("备份聊天记录失败：", err)
	}
	log.Println("备份聊天记录成功")
}
