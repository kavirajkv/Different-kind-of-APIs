package main

import (
	"context"
	"io"
	"log"
	"net"
	"time"

	proto "github.com/kavirajkv/api-types/GRPC/protobuf/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct {
	proto.UnimplementedChatServiceServer
}

func (s *server) GetUserInfo(ctx context.Context, req *proto.UserId) (*proto.UserInfo, error) {
	if req.Id == 0 {
		return nil, status.Error(codes.InvalidArgument, "UserId required")
	}
	log.Printf("Received request for userid: %v", req.Id)
	return &proto.UserInfo{Id: req.Id, Name: "John"}, nil

}

func (s *server) SendMessage(stream proto.ChatService_SendMessageServer) error {
	for{
		req,err:=stream.Recv()
		if err==io.EOF{
			return stream.SendAndClose(&proto.Status{Status: "Messages received successfully"})
		}
		if err!=nil{
			log.Fatalf("Error receiving messages %s",err.Error())
			return err
		}
		log.Printf("Message received from %v to %v at %v",req.Senderid,req.Receiverid,req.Time)
		log.Println("Message received: ",req.Msg)
	}

}

func (s *server)ReceiveMessage(req *proto.UserId, stream proto.ChatService_ReceiveMessageServer)error{
	for x:=range 3{
		res:=&proto.Message{Senderid: int32(x),Receiverid: req.Id,Msg: "hello",Time: time.Now().Unix()}
		err:=stream.Send(res)
		if err!=nil{
			log.Fatalln(err.Error())
		}
		time.Sleep(2*time.Second)
	}
	return nil
}

func main() {
	lis, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	proto.RegisterChatServiceServer(s, &server{})

	log.Println("Server running at port 8000")
	er := s.Serve(lis)
	if er != nil {
		log.Fatal("Server failed ", er)
	}

}
