package main

//
//import (
//	"context"
//	"google.golang.org/grpc"
//	"io"
//	"chatroom/pkg/apis"
//	"log"
//)
//
//func main() {
//	conn, err := grpc.Dial("localhost:9090", grpc.WithInsecure())
//	if err != nil {
//		log.Fatalf("did not connect: %v", err)
//	}
//	c := apis.NewChatServiceClient(conn)
//	w, err := c.WatchOnline(context.TODO(), &apis.Null{})
//	if err != nil {
//		log.Fatal("启动watcher失败：", err)
//	}
//	out:=&apis.OnlineUserList{}
//	for {
//		ol, err := w.Recv()
//		if err != nil {
//			if err == io.EOF {
//				log.Println("服务器告知说完了")
//				break
//			}
//			log.Fatal("接收异常：", err)
//		}
//		out=append(out,ol)
//		log.Println("收到新变动：", ol.String())
//	}
//}
