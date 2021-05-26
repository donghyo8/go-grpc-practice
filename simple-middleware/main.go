// https://github.com/grpc-ecosystem/go-grpc-middleware

package main

import (
	"context"
	"log"
	"net"
	"time"

	"go-grpc/data"
	userpb "go-grpc/protos/v1/user"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const portNumber = "9000"

type userServer struct {
	userpb.UserServer
}

// gRPC 서버를 시작하고 서버가 호출 될 때마다 customMiddleware() 함수를 거쳐감
func customMiddleware() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (_ interface{}, err error) {
		log.Print("Requested at:", time.Now())

		resp, err := handler(ctx, req)
		return resp, err
	}
}

// user.proto에 정의한 user의 모든 정보를 조회하는 rpc
func (s *userServer) ListUsers(ctx context.Context, req *userpb.ListUsersRequest) (*userpb.ListUsersResponse, error) {
	userMessages := make([]*userpb.UserMessage, len(data.UserData))
	for i, u := range data.UserData {
		userMessages[i] = u
	}

	return &userpb.ListUsersResponse{
		UserMessages: userMessages,
	}, nil
}

func customAuthFunc(ctx context.Context) (context.Context, error) {
	token, err := grpc_auth.AuthFromMD(ctx, "bearer")
	if err != nil {
		return nil, err
	}

	if token != "customToken" {
		return nil, status.Errorf(codes.Unauthenticated, "invalid auth token: %v", err)
	}

	newCtx := context.WithValue(ctx, "token", token)

	return newCtx, nil
}

func main() {
	lis, err := net.Listen("tcp", ":"+portNumber)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	logrus.ErrorKey = "grpc.error"
	logrusEntry := logrus.NewEntry(logrus.StandardLogger())

	// gRPC 서버에 미들웨어를 추가하는 방법이 여러가지 존재함
	// grpcServer := grpc.NewServer() 기존 gRPC middleware 없는 방식
	grpcServer := grpc.NewServer(
		// Unary format의 서버에 interceptor를 추가해주는 서버 옵션
		// 이 방식은 하나의 interceptor만 추가가 가능함
		// 여러개의 interceptor들을 추가하고 싶어도 추가하지 못하는 방식임
		// 이를 위해, grpc_middleware.ChainUnaryServer 사용(이 함수의 파라미터로 여러개의 interceptor를 주입해 여러개의 interceptor들을 체인처럼 하나씩 순서대로 실행할 수 있게 해줌)
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			customMiddleware(),
			// 로깅
			grpc_logrus.UnaryServerInterceptor(logrusEntry),
			// 토큰을 보내지 않으면 인증 실패
			grpc_auth.UnaryServerInterceptor(customAuthFunc),
			// rpc내에서 panic이 발생해도 서버 종료 X
			grpc_recovery.UnaryServerInterceptor(),
		)),
	)
	userpb.RegisterUserServer(grpcServer, &userServer{})

	log.Printf("start gRPC server on %s port", portNumber)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
