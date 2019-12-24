package main

import (
	"dst-client-go/src/proto"
	"fmt"
"context"
"google.golang.org/grpc"
"google.golang.org/grpc/grpclog"
"net"
	"testing"
)

// 定义helloService并实现约定的接口
type dstStringService struct{}

// HelloService Hello服务
var ssss = dstStringService{}

// SayHello 实现Hello服务接口
func (h dstStringService) Put(ctx context.Context, in *com_distkv_drpc_pb_generated.PutRequest) (*com_distkv_drpc_pb_generated.PutResponse, error) {
	resp := new(com_distkv_drpc_pb_generated.PutResponse)
	resp.Status = 1
	fmt.Println(in.Value)
	fmt.Println(in.Key)
	return resp, nil
}

// SayHello 实现Hello服务接口
func (h dstStringService) Get(ctx context.Context, in *com_distkv_drpc_pb_generated.GetRequest) (*com_distkv_drpc_pb_generated.GetResponse, error) {
	resp := new(com_distkv_drpc_pb_generated.GetResponse)
	resp.Status = 1
	fmt.Println("get")
	return resp, nil
}

func Test_b_test(t *testing.T)  {
	fmt.Println("start")

	listen, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		grpclog.Fatalf("Failed to listen: %v", err)
	}

	// 实例化grpc Server
	s := grpc.NewServer()

	// 注册HelloService

	com_distkv_drpc_pb_generated.RegisterDstStringServiceServer(s,ssss)

	grpclog.Println("Listen on 127.0.0.1:8080" )
	s.Serve(listen)

}
