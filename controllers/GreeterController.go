package controllers

import (
	"context"
	pb "grpc-example/hellostream"
	"io"
	"log"
	"strconv"
	"sync"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GreeterService struct {
	pb.UnimplementedGreeterServer
}

// SayHello 一元RPC方法(客户端发起一个请求，从服务端得到一个响应，和普通函数调用相似)
func (s *GreeterService) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	log.Println("收到客户端来的数据", req.GetData())
	return &pb.HelloResponse{Data: "Hello " + req.GetData()}, nil
}

// GetStream 服务端流(客户端发起一个请求，从服务端得到一个消息流，客户端可以一直从流中读取数据直到读取完服务端返回的所有数据)
func (s *GreeterService) GetStream(req *pb.HelloRequest, res pb.Greeter_GetStreamServer) error {
	if req.Data == "hello" {
		for i := 0; i < 6; i++ {
			err := res.Send(&pb.HelloResponse{
				Data: time.Now().Format("2006-01-02 15:04:05"),
			})
			if err != nil {
				log.Println("send to client failed", err.Error())
				return status.Error(codes.Canceled, "send to client failed")

			}
			time.Sleep(time.Second)
		}

	} else {
		log.Println("err request")
	}
	return nil

}

// SetStream 客户端流（客户端发起一个消息流到服务端，服务端收到数据流后执行响应）
func (s *GreeterService) SetStream(req pb.Greeter_SetStreamServer) error {

	for {
		// 接收客户端发来的流式数据
		data, err := req.Recv()

		if err != nil {
			if err == io.EOF {
				log.Println("客户端发送完毕")
				break
			} else {
				log.Println("其他错误", err.Error())
			}
			return err
		} else {
			log.Println("收到从客户端返回的数据", data.Data)
			continue
		}
	}
	//返回数据给客户端
	req.SendMsg(&pb.HelloResponse{
		Data: "SetStream 执行完毕",
	})
	return nil
}

// AllStream 双向流 （客户端发起一个消息流到服端，服务端收到流数据后返回另外一个数据流给客户端)
func (s *GreeterService) AllStream(req pb.Greeter_AllStreamServer) error {
	w := sync.WaitGroup{}
	w.Add(1)
	go func() {
		for {
			recvMsg, err := req.Recv()
			if err != nil {
				if err == io.EOF {
					//客户端主动关闭发送
					log.Println("客户端发送完毕", err.Error())
				} else {
					log.Println("其他错误", err.Error())
				}
				break
			}
			log.Println("收到客户端来的数据：", recvMsg.Data)
		}
		w.Done()
	}()

	w.Add(1)
	go func() {
		for i := 1; i < 3; i++ {
			if err := req.Send(&pb.HelloResponse{
				Data: "从服务端返回：" + strconv.Itoa(i),
			}); err != nil {
				log.Println("发送失败", err.Error())
				break
			}
			time.Sleep(time.Second)

		}
		w.Done()
	}()
	w.Wait()
	return nil
}
