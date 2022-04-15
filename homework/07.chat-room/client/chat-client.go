package main

import (
	"chatroom/pkg/apis"
	"context"
	"crypto/rand"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"math/big"
)

func main() {
	conn, err := grpc.Dial("localhost:9090", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	c := apis.NewChatServiceClient(conn)
	for {
		fmt.Println("请输入指令：")
		var command1 string
		var command2 string
		fmt.Scanln(&command1, &command2)
		if command1 == "chat" {
			switch command2 {
			case "register":
				register(c)
			case "login":
				login(c)
			case "list":
				login(c)
			case "history":
				login(c)
			case "subscribe":
				login(c)
			}
		} else {
			fmt.Printf("错误指令，")
		}
	}
}

func login(c apis.ChatServiceClient) {

}

func register(c apis.ChatServiceClient) {
	act := &apis.Account{}
	fmt.Print("输入昵称：")
	fmt.Scanln(&act.Name)
	fmt.Print("输入密码：")
	fmt.Scanln(&act.Password)
	num, err := rand.Int(rand.Reader, big.NewInt(100000000000))
	act.Account = num.Int64()
	ret, err := c.Register(context.TODO(), act)
	if err != nil {
		log.Fatal("注册失败：", err)
	}
	fmt.Println("注册成功，您的账号为：", ret.Account)
	log.Println("注册成功", ret)
}
