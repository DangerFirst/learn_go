package main

import (
	"bufio"
	"chatroom/pkg/apis"
	"context"
	"crypto/rand"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"math/big"
	"os"
	"strings"
)

func main() {
	conn, err := grpc.Dial("localhost:9090", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	c := apis.NewChatServiceClient(conn)
	for {
		fmt.Print("请输入指令：")
		command := stdInput()
		if command == "chat register" {
			register(c)
		} else if command == "chat login" {
			act := login(c)
			if act != nil {
				for {
					fmt.Print("请输入指令：")
					commandAfterLogin := stdInput()
					if commandAfterLogin == "chat list" {
						onlineUser(c)
					} else if commandAfterLogin == "chat with" {
						chatWith(c, act)
					} else if commandAfterLogin == "chat history" {
						chatHistory(c, act)
					} else if commandAfterLogin == "exit 0" {
						fmt.Println("退出登录成功")
						logOut(c, act)
						break
					} else {
						fmt.Printf("错误指令，")
					}
				}
			}
		} else {
			fmt.Printf("错误指令，")
		}
	}
}

func stdInput() string {
	var command string
	reader := bufio.NewReader(os.Stdin)  // 标准输入输出
	command, _ = reader.ReadString('\n') // 回车结束
	command = strings.TrimSpace(command)
	return command
}

func connectDb() *gorm.DB {
	conn, err := gorm.Open(mysql.Open("root:123qwe@tcp(localhost)/testdb"))
	if err != nil {
		log.Fatal("数据库连接失败：", err)
	}
	//fmt.Println("连接数据库成功")
	return conn
}

func logOut(c apis.ChatServiceClient, act *apis.Account) {
	_, err := c.LogOut(context.TODO(), act)
	if err != nil {
		log.Fatal("退出失败：", err)
	}
}

func chatHistory(c apis.ChatServiceClient, act *apis.Account) {
	conn := connectDb()
	var chatServer apis.ChatServiceServer = NewDbChat(conn, &dbChat{})
	_, err := c.ChatRecord(context.TODO(), act)
	if err != nil {
		log.Fatal("查询失败：", err)
	}
	_, err = chatServer.ChatRecord(context.TODO(), act)
	if err != nil {
		log.Fatal("查询失败：", err)
	}
}

func chatWith(c apis.ChatServiceClient, act *apis.Account) {
	//conn := connectDb()
	//var chatServer apis.ChatServiceServer=NewDbChat(conn,&dbChat{})
	//var toAct string
	//chy := &apis.ChatHistory{}
	//fmt.Print("输入要交流的对象账号：")
	//fmt.Scanln(&toAct)
	//fmt.Print("内容：")
	//fmt.Scanln(&chy.Record)
	//
	//fmt.Printf("%s\n：",act.Name)
	//chatServer.Chat(context.TODO(),act)

}

func onlineUser(c apis.ChatServiceClient) {
	ret, err := c.OnlineUser(context.TODO(), &apis.Null{})
	if err != nil {
		log.Fatal("查询失败：", err)
	}
	fmt.Println("在线用户：")
	for _, v := range ret.OnlineUsers {
		fmt.Printf("账号：%d 昵称：%s\n", v.Account, v.Name)
	}
}

func login(c apis.ChatServiceClient) (act *apis.Account) {
	conn := connectDb()
	var chatServer apis.ChatServiceServer = NewDbChat(conn, &dbChat{})
	act = &apis.Account{}
	fmt.Print("输入账号：")
	fmt.Scanln(&act.Account)
	fmt.Print("输入密码：")
	fmt.Scanln(&act.Password)
	ret, err := chatServer.Login(context.TODO(), act)
	if err != nil {
		log.Fatal("登录失败：", err)
	}
	if ret != nil {
		_, err = c.Login(context.TODO(), ret)
		if err != nil {
			fmt.Println("登录失败：", err)
		}
		fmt.Printf("登录成功，欢迎你回来，%s!\n", ret.Name)
	}
	return ret
}

func register(c apis.ChatServiceClient) {
	conn := connectDb()
	var chatServer apis.ChatServiceServer = NewDbChat(conn, &dbChat{})
	act := &apis.Account{}
	fmt.Print("输入昵称：")
	fmt.Scanln(&act.Name)
	fmt.Print("输入密码：")
	fmt.Scanln(&act.Password)
	num, err := rand.Int(rand.Reader, big.NewInt(100000000000))
	act.Account = num.Int64()
	_, err = chatServer.Register(context.TODO(), act)
	if err != nil {
		log.Fatal("注册失败：", err)
	}
	_, err = c.Register(context.TODO(), act)
	if err != nil {
		log.Fatal("注册失败：", err)
	}
	fmt.Println("注册成功，您的账号为：", act.Account)
}
