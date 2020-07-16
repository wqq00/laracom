package test

import (
	"log"
	"context"
	"google.golang.org/grpc"

	pb "laracom/demo-service/proto/demo"
)

const (
	address = "localhost:9999"
)


func main()  {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("连接到 gRPC 服务器失败: %v", err)
	}

	defer conn.Close()

	client := pb.NewDemoServiceClient(conn)
	req := &pb.DemoRequest{Name: "学院君"}
	rsp, err := client.SayHello(context.Background(), req)
	if err != nil {
		log.Fatalf("调用 gRPC 服务接口失败: %v", err)
	}
	log.Printf("%s", rsp.Text)
}