package main

import (
	"context"
	"log"
	"net"

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
