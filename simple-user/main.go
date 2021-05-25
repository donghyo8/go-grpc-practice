package main

import (
	"context"
	"log"
	"net"

	"go-grpc/data"
	userpb "go-grpc/protos/v1/user"

	"google.golang.org/grpc"
)

const portNumber = "9000"

type userServer struct {
	userpb.UserServer
}

// 2. user.proto에 정의한 user_id로 user의 정보를 갖고오는 rpc
func (s *userServer) GetUser(ctx context.Context, req *userpb.GetUserRequest) (*userpb.GetUserResponse, error) {
	userID := req.UserId

	var userMessage *userpb.UserMessage
	for _, u := range data.UserData {
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
	userMessages := make([]*userpb.UserMessage, len(data.UserData))
	for i, u := range data.UserData {
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

	grpcServer := grpc.NewServer()
	// 1. user_grpc.pb.go 파일의 RegisterUserServer 함수를 가져와서 user 서비스를 등록하면 user 서비스를 담당하는 grpc server가 생성됨
	userpb.RegisterUserServer(grpcServer, &userServer{})

	log.Printf("start gRPC server on %s port", portNumber)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}

// BoolmRPC로 user.proto 파일을 import 하면 user.proto에 정의된 rpc들을 확인할 수 있음
// 서버 실행 후 해당 포트로 reuqest 형식에 맞게 데이터를 request하면 response됨
// getUser: 이미 static하게 저장해둔 데이터 필드들을 id값에 맞게 보여줌
// ListUser: 저장된 모든 데이터들을 보여줌
