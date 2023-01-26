package main

import (
	"context"
	"io"
	"log"
	"strconv"
	"time"

	pb "grpc-example/hellostream"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewGreeterClient(conn)

	RunSayHello(client)
	RunGetStream(client)
	RunSetStream(client)
	RunAllStream(client)
}

func RunSayHello(client pb.GreeterClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := client.SayHello(ctx, &pb.HelloRequest{Data: "你好"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("收到服务端来的数据: %s", r.GetData())
}

// RunGetStream 服务端流(客户端发起一个请求，从服务端得到一个消息流，客户端可以一直从流中读取数据直到读取完服务端返回的所有数据)
func RunGetStream(client pb.GreeterClient) {
	stream, err := client.GetStream(context.Background(), &pb.HelloRequest{Data: "hello"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	for {
		// 接收服务端返回的流式数据，当收到io.EOF或错误时退出
		res, err := stream.Recv()

		if err != nil {
			if err == io.EOF {
				log.Println("服务端发送完毕")
				break
			}
			log.Println("其他错误", err.Error())
		}
		log.Println("收到从服务端返回的数据", res.Data)
	}
}

// RunSetStream 客户端流（客户端发起一个消息流到服务端，服务端收到数据流后执行响应）
func RunSetStream(client pb.GreeterClient) {
	// 客户端流式RPC
	stream, err := client.SetStream(context.Background())
	if err != nil {
		log.Fatalf("c.LotsOfGreetings failed, err: %v", err)
	}
	names := []string{"项羽", "安琪拉", "妲己"}
	for _, name := range names {
		// 发送流式数据
		err := stream.Send(&pb.HelloRequest{Data: name})
		if err != nil {
			log.Fatalf("client.RunSetStream stream.Send(%v) failed, err: %v", name, err)
		}
	}
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("client.RunSetStream failed: %v", err)
	}
	log.Printf("got reply: %v", res.GetData())
}

func RunAllStream(client pb.GreeterClient) {
	//双向流
	stream, err := client.AllStream(context.Background())
	if err != nil || stream == nil {
		panic("init failed")
	}
	//发送流给服务端
	go func() {

		for i := 1; i < 6; i++ {

			if err := stream.Send(&pb.HelloRequest{
				Data: "client" + strconv.Itoa(i),
			}); err != nil {
				log.Println("send failed", err.Error())
				return
			}
		}
		//发送流完成关闭
		if err := stream.CloseSend(); err != nil {
			log.Println("send close failed", err.Error())
			return
		}
	}()

	//接收服务端发送过来的流数据

	go func() {

		for {
			data, err := stream.Recv()
			if err != nil {
				if err == io.EOF {
					log.Println("server close conn")
				} else {
					log.Println("recv from server failed", err.Error())
				}
				return
			}
			log.Println("recv from server", data.Data)
		}
	}()
	select {}
}
