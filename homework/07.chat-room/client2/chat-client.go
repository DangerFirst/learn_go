package main

import (
	"bufio"
	"chatroom/pkg/apis"
	"context"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"io/ioutil"
	"log"
	"math/big"
	"os"
	"strings"
)

func main() {
	fmt.Println("chat register：注册|chat login：登录|chat list：查看在线用户|" +
		"chat with：发起聊天|chat history：查看聊天记录|chat subscribe：接收消息|" +
		"not auto login：清除保存的密码|exit 0：退出")
	conn, err := grpc.Dial("localhost:9090", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	c := apis.NewChatServiceClient(conn)
	filePath := "E:/password.json"
	records := NewRecord(filePath)
	for {
		fmt.Print("请输入指令：")
		command := stdInput()
		if command == "chat register" {
			register(c)
		} else if command == "chat login" {
			act := login(c)
			if err := records.saveAccount(act); err != nil {
				log.Fatal("保存失败：", err)
			}
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
					} else if commandAfterLogin == "chat subscribe" {
						revMessage(c, act)
					} else if commandAfterLogin == "exit 0" {
						fmt.Println("退出登录成功")
						logOut(c, act)
						break
					} else {
						fmt.Printf("错误指令，")
					}
				}
			}
		} else if command == "not auto login" {
			notAutoLogin()
		} else if command == "exit 0" {
			fmt.Println("退出程序")
			break
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

func notAutoLogin() {
	filePath := "E:/password.json"
	records := NewRecord(filePath)
	records.saveAccount(&apis.Account{
		Id:       1,
		Account:  1,
		Password: "1",
		Name:     "1",
	})
	fmt.Println("清除保存的密码成功")
}

func revMessage(c apis.ChatServiceClient, act *apis.Account) {
	conn := connectDb()
	var chatServer apis.ChatServiceServer = NewDbChat(conn, &dbChat{})
	_, err := c.RevMessage(context.TODO(), act)
	if err != nil {
		log.Fatal("接收失败：", err)
	}
	_, err = chatServer.RevMessage(context.TODO(), act)
	if err != nil {
		fmt.Println("查询失败：", err)
	}
}

func chatHistory(c apis.ChatServiceClient, act *apis.Account) {
	conn := connectDb()
	var chatServer apis.ChatServiceServer = NewDbChat(conn, &dbChat{})
	ctx := context.Background()
	var toAct string
	_, err := c.ChatRecord(ctx, act)
	if err != nil {
		log.Fatal("查询失败：", err)
	}
	fmt.Print("输入要查询的对象账号：")
	fmt.Scanln(&toAct)
	ctx = context.WithValue(ctx, "toAct", toAct)
	_, err = chatServer.ChatRecord(ctx, act)
	if err != nil {
		fmt.Println("查询失败：", err)
	}
}

func chatWith(c apis.ChatServiceClient, act *apis.Account) {
	conn := connectDb()
	var chatServer apis.ChatServiceServer = NewDbChat(conn, &dbChat{})
	ctx := context.Background()
	var toAct string
	var text string
	fmt.Print("输入要交流的对象账号：")
	fmt.Scanln(&toAct)
	fmt.Print("内容：")
	fmt.Scanln(&text)
	ctx = context.WithValue(ctx, "toAct", toAct)
	ctx = context.WithValue(ctx, "text", text)
	_, err := chatServer.Chat(ctx, act)
	if err != nil {
		log.Fatal("发起失败：", err)
	}
	ctx = context.WithValue(ctx, "mg", "111")
	ret, err := c.Chat(ctx, act)
	fmt.Println(ret)
	if err != nil {
		log.Fatal("发起失败：", err)
	}

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
	data, err := ioutil.ReadFile("E:/password.json")
	json.Unmarshal(data, act)
	if err != nil || act.Account == 0 {
		fmt.Print("输入账号：")
		fmt.Scanln(&act.Account)
		fmt.Print("输入密码：")
		fmt.Scanln(&act.Password)
	}
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
