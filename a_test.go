package main

import (
	"bufio"
	"context"
	"dst-client-go/src/proto"
	"encoding/json"
	"fmt"
	"google.golang.org/grpc"
	"testing"
)

type writer struct {
	w *bufio.Writer

}

//func (w *writer) writeArgs()  {
//	var req = &Request{
//		RequestId:
//
//	}
//
//}

//func generateId(){
//	t := time.Now().Unix()
//}

func  Test_a_test(t *testing.T) {

	conn, err := grpc.Dial("127.0.0.1:8080", grpc.WithInsecure())
	if err != nil{
		fmt.Println(err.Error())
	}
	defer conn.Close()
	client := com_distkv_drpc_pb_generated.NewDstStringServiceClient(conn)
	req := &com_distkv_drpc_pb_generated.PutRequest{Key:"dstPut",Value:"PutValue"}
	p,err := client.Put(context.Background(),req)
	if err != nil{
		fmt.Println(1)
		fmt.Println(err.Error())
	}
	a,err := json.Marshal(p)
	if err != nil{
		fmt.Println(2)
		fmt.Println(err)
	}
	fmt.Println(string(a))
	}


