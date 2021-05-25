package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"

	userpb "go-grpc/protos/v2/user"

	"go-grpc/data"
)

const portNumber = "9001"

type userServer struct {
	userpb.UserServer
}

// 2. user.proto에 정의한 user_id로 user의 정보를 갖고오는 rpc
func (s *userServer) GetUser(ctx context.Context, req *userpb.GetUserRequest) (*userpb.GetUserResponse, error) {
	userID := req.UserId

	var userMessage *userpb.UserMessage
	for _, u := range data.UsersV2 {
		if u.UserId != userID {
			continue
		}
		userMessage = u
		break
	}

	return &userpb.GetUserResponse{
		UserMessage: userMessage,
	}, nil
}

// 3. user.proto에 정의한 user의 모든 정보를 조회하는 rpc
func (s *userServer) ListUsers(ctx context.Context, req *userpb.ListUsersRequest) (*userpb.ListUsersResponse, error) {
	userMessages := make([]*userpb.UserMessage, len(data.UsersV2))
	for i, u := range data.UsersV2 {
		userMessages[i] = u
	}

	return &userpb.ListUsersResponse{
		UserMessages: userMessages,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":"+portNumber)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// 1. user_grpc.pb.go 파일의 RegisterUserServer 함수를 가져와서 user 서비스를 등록하면 user 서비스를 담당하는 grpc server가 생성됨
	grpcServer := grpc.NewServer()
	userpb.RegisterUserServer(grpcServer, &userServer{})

	log.Printf("start gRPC server on %s port", portNumber)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
