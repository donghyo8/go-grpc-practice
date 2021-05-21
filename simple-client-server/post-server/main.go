package main

import (
	"context"
	"log"
	"net"
	"fmt"

	"google.golang.org/grpc"

	"go-grpc/data"
	postpb "go-grpc/protos/v1/post"
	userpb "go-grpc/protos/v1/user"
	client "go-grpc/simple-client-server"
)

const portNumber = "9001"

type postServer struct {
	postpb.PostServer
	userCli userpb.UserClient // 1. user 서비스를 사용 할 수 있도록 userpb.UserClient 타입 선언
}

// user id 전달 시, user id가 등록한 모든 정보들을(post된) 리턴하는 rpc
func (s *postServer) ListPostsByUserId(ctx context.Context, req *postpb.ListPostsByUserIdRequest) (*postpb.ListPostsByUserIdResponse, error) {
	userID := req.UserId

	// 기존 생성한 user gRPC 서버의 GetUser rpc 호출(user id를 받아오기 위함) 
	resp, err := s.userCli.GetUser(ctx, &userpb.GetUserRequest{UserId: userID})
	fmt.Print(resp)
	if err != nil {
		return nil, err
	}

	var postMessages []*postpb.PostMessage

	for _, up := range data.UserPosts {
		if up.UserID != userID {
			continue
		}

		for _, p := range up.Posts {
			p.Author = resp.UserMessage.Name
		}

		postMessages = up.Posts
		break
	}

	return &postpb.ListPostsByUserIdResponse{
		PostMessages: postMessages,
	}, nil
}

// 서비스에 등록된 모든 정보들을(post된) 리턴하는 rpc
// post 서비스는 어떤 user id가 post로 정보를 등록했는지 알지만, user id에 해당하는 Author는 모른다고 가정
// protobuf 정의에 의해 post 서비스의 rpc들은 Author 필드에 이름을 채워 전달
func (s *postServer) ListPosts(ctx context.Context, req *postpb.ListPostsRequest) (*postpb.ListPostsResponse, error) {
	var postMessages []*postpb.PostMessage
	for _, up := range data.UserPosts {
		resp, err := s.userCli.GetUser(ctx, &userpb.GetUserRequest{UserId: up.UserID})
		if err != nil {
			return nil, err
		}

		for _, p := range up.Posts {
			p.Author = resp.UserMessage.Name
		}

		postMessages = append(postMessages, up.Posts...)
	}

	return &postpb.ListPostsResponse{
		PostMessages: postMessages,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":"+portNumber)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// 2. 커넥션 선언
	// client의 커넥션을 리턴하는 함수(GetUserClient)의 구현부분은 user_client_.go
	userCli := client.GetUserClient("localhost:9000")
	grpcServer := grpc.NewServer()
	
	// 2. user gRPC 서버와 통신할 수 있는 Client를 struct에 넣어줌
	// post gRPC 서버 내에서 user gRPC 서버에 접근할 수 있는 client를 struct로 가지고 있어서 코드 내 접근 가능
	postpb.RegisterPostServer(grpcServer, &postServer{
		userCli: userCli,
	})

	log.Printf("start gRPC server on %s port", portNumber)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
// simple/main.go, post-server/main.go 두개 서버를 동시에 실행
// 포트번호 9000, 9001
// 
