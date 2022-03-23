package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/armstrongli/go-bmi"
	"learn.go/pkg/apis"
	"log"
	"net"
	"time"
)

func main() {
	var port string
	flag.StringVar(&port, "port", "8080", "配置启动端口")
	flag.Parsed()

	ln, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatal(err)
	}

	rank := NewFatRateRank()

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println("warning:建立连接失败:", err)
			continue
		}
		fmt.Println(conn)
		go talk(conn, rank)
	}
}

func talk(conn net.Conn, rank *FatRateRank) {
	defer fmt.Println("结束连接：", conn)
	defer conn.Close()
	for {
		finalReceivedMessage := make([]byte, 0, 1024)
		encounterError := false
		for {
			buf := make([]byte, 1024)
			valid, err := conn.Read(buf)
			if err != nil {
				log.Println("WARNING:读取数据时出问题：", err)
				log.Println("重新读取", err)
				encounterError = true
				time.Sleep(1 * time.Second)
				break
			}
			if valid == 0 {
				break
			}
			validContent := buf[:valid]
			finalReceivedMessage = append(finalReceivedMessage, validContent...)
			if valid < len(buf) {
				break
			}
		}
		if encounterError {
			conn.Write([]byte(`服务器获取数据失败，请重新输入`))
			continue
		}
		pi := &apis.PersonalInformation{}
		if err := json.Unmarshal(finalReceivedMessage, pi); err != nil {
			conn.Write([]byte(`输入的数据无法解析，请重新输入`))
			continue
		}
		bmi, err := gobmi.BMI(pi.Weight, pi.Tall)
		if err != nil {
			conn.Write([]byte(`无法计算您的BMI，请重新输入`))
			continue
		}
		fr, err := gobmi.Fatrate(pi.Age, pi.Sex, bmi)
		rank.inputRecord(pi.Name, fr)
		rankId, _ := rank.getRank(pi.Name)

		conn.Write([]byte(fmt.Sprintf("姓名：%s, BMi：%v, 体脂率：%v,排名：%d", pi.Name, bmi, fr, rankId))) //todo handle error
		break
	}
}
