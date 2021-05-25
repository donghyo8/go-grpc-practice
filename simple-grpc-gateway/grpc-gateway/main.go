package main

import (
	"context"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"

	userpb "go-grpc/protos/v2/user"
)

const (
	portNumber           = "9000"
	gRPCServerPortNumber = "9001"
)

func main() {
	// 1. gRPC gateway와 gRPC 서버를 이어주기 위해 context 패키지 선언
	ctx := context.Background()
	/*
	   3. gRPC gateway 패키지 내 mux 선언
	   mux는 http 요청이 오면 grpc 서버에 그대로 보낼지 특정 요청들만 보낼지와 같은 작업을 해줌

	*/
	mux := runtime.NewServeMux()
	options := []grpc.DialOption{
		grpc.WithInsecure(),
	}

	/*
		2. ctx를 RegisterUserHandlerFromEndpoint에 주입
		gRPC gateway는 gRPC 서버에서 context done 신호가 오면 커넥션을 끊음
		RegisterUserHandlerFromEndpoint는 4개의 파라미터로 구성
		- ctx: gRPC 서버와 통신을 할 때 전달하기 위해 사용
		- mux: 각 request마다 미들웨어처럼 http 요청을 정의한 옵션대로 구성하는데 사용됨
		-  주소 - 통신할 주소
		- dial option: client-server 간 커넥션 시 추가적 작업을 위해 사용

	*/
	if err := userpb.RegisterUserHandlerFromEndpoint(
		ctx,
		mux,
		"localhost:"+gRPCServerPortNumber,
		options,
	); err != nil {
		log.Fatalf("failed to register gRPC gateway: %v", err)
	}

	log.Printf("start HTTP server on %s port", portNumber)
	if err := http.ListenAndServe(":"+portNumber, mux); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
