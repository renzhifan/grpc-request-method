package controllers

import (
	"context"
	pb "grpc-example/hellostream"
	"log"
)

type TestService struct {
	pb.UnimplementedTestServer
}

// SayHello 一元RPC方法(客户端发起一个请求，从服务端得到一个响应，和普通函数调用相似)
func (s *TestService) Test(ctx context.Context, req *pb.HelloReq) (*pb.HelloRly, error) {
	log.Println("收到客户端来的数据", req.GetData())
	var reply pb.HelloRly

	reply.MapField = make(map[uint32]*pb.WikiInfo)
	reply.MapField[1] = &pb.WikiInfo{
		WikiId:   1,
		WikiName: "哈哈",
	}
	return &reply, nil
}
