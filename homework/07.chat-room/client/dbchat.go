package main

import (
	"chatroom/pkg/apis"
	"context"
	"fmt"
	"gorm.io/gorm"
	"log"
	"time"
)

var _ apis.ChatServiceServer = &dbChat{}

func NewDbChat(conn *gorm.DB,
	embedChat apis.ChatServiceServer) apis.ChatServiceServer {
	if conn == nil {
		log.Fatal("数据库连接为空")
	}
	if embedChat == nil {
		log.Fatal("内存为空")
	}
	return &dbChat{
		conn:      conn,
		embedChat: embedChat,
	}
}

type dbChat struct {
	conn      *gorm.DB
	embedChat apis.ChatServiceServer
}

func (d dbChat) LogOut(ctx context.Context, account *apis.Account) (*apis.Account, error) {
	return nil, nil
}

func (d dbChat) Register(ctx context.Context, account *apis.Account) (*apis.Account, error) {
	resp := d.conn.Create(account)
	if err := resp.Error; err != nil {
		fmt.Printf("创建%s时失败：%v\n", account.Name, err)
		return nil, err
	}
	//fmt.Printf("创建%s成功\n", account.Name)
	return account, nil
}

func (d dbChat) Login(ctx context.Context, account *apis.Account) (*apis.Account, error) {
	var tables []*apis.Account
	resp := d.conn.Where("account=? and password=?", account.Account, account.Password).Find(&tables)
	if err := resp.Error; err != nil {
		fmt.Println("登录失败：", err)
		return nil, err
	}
	if len(tables) == 0 {
		fmt.Println("账号密码错误")
		return nil, nil
	}
	return tables[0], nil
}

func (d dbChat) OnlineUser(ctx context.Context, null *apis.Null) (*apis.OnlineUserList, error) {
	return nil, nil
}

func (d dbChat) Chat(ctx context.Context, account *apis.Account) (*apis.ChatHistory, error) {
	var tables []*apis.ChatHistory
	toAct := ctx.Value("toAct")
	text := fmt.Sprint(ctx.Value("text"))
	resp := d.conn.Where("(talker_account=? and listener_account=?) or (talker_account=? and listener_account=?)", account.Account, toAct, toAct, account.Account).Order("id desc").Limit(20).Find(&tables)
	if err := resp.Error; err != nil {
		fmt.Println("查询失败：", err)
		return nil, err
	}
	fmt.Printf("%s:%s\n", account.Name, text)
	fmt.Println("历史聊天记录：")
	if len(tables) == 0 {
		fmt.Println("没有聊天记录")
	} else {
		for _, v := range tables {
			fmt.Printf("%s|%d:%s->%s|%d %s\n", v.Talker, v.TalkerAccount, v.Record, v.Listener, v.ListenerAccount, v.CreateDate)
		}
	}

	var table *apis.Account
	resp = d.conn.Where("account=?", toAct).Find(&table)
	if err := resp.Error; err != nil {
		fmt.Println("查询交流账户失败：", err)
		return nil, err
	}

	chy := &apis.ChatHistory{
		TalkerAccount:   account.Account,
		Talker:          account.Name,
		ListenerAccount: table.Account,
		Listener:        table.Name,
		Record:          text,
		CreateDate:      time.Now().Local().String(),
	}

	resp = d.conn.Create(chy)
	if err := resp.Error; err != nil {
		fmt.Printf("创建%s的聊天记录失败：%v\n", chy.Talker, err)
		return nil, err
	}
	return nil, nil
}

func (d dbChat) ChatRecord(ctx context.Context, account *apis.Account) (*apis.ChatHistory, error) {
	var tables []*apis.ChatHistory
	toAct := ctx.Value("toAct")
	resp := d.conn.Where("(talker_account=? and listener_account=?) or (talker_account=? and listener_account=?)", account.Account, toAct, toAct, account.Account).Find(&tables)
	if err := resp.Error; err != nil {
		fmt.Println("查询失败：", err)
		return nil, err
	}
	if len(tables) == 0 {
		fmt.Println("没有聊天记录")
		return nil, nil
	}
	for _, v := range tables {
		fmt.Printf("%s|%d:%s->%s|%d %s\n", v.Talker, v.TalkerAccount, v.Record, v.Listener, v.ListenerAccount, v.CreateDate)
	}
	return nil, nil
}

func (d dbChat) RevMessage(ctx context.Context, account *apis.Account) (*apis.ChatHistory, error) {
	var tables []*apis.ChatHistory
	resp := d.conn.Where("listener_account=? and is_read=0", account.Account).Find(&tables)
	if err := resp.Error; err != nil {
		fmt.Println("接收失败：", err)
		return nil, err
	}
	if len(tables) == 0 {
		fmt.Println("没有新消息")
		return nil, nil
	}
	for _, v := range tables {
		fmt.Printf("%s|%d:%s %s\n", v.Talker, v.TalkerAccount, v.Record, v.CreateDate)
	}
	resp = d.conn.Model(&tables).Where("listener_account=? and is_read=0", account.Account).Update("is_read", 1)
	if err := resp.Error; err != nil {
		fmt.Println("标记为已读失败", err)
		return nil, err
	}
	return nil, nil
}
