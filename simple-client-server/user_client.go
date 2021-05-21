package simple_client_server

import (
	"sync"

	"google.golang.org/grpc"

	userpb "go-grpc/protos/v1/user"
)

var (
	once sync.Once // gRPC 서버 내 싱글톤으로 초기에 한번만 client를 생성하고, 각 rpc 내에서는 같은 client를 사용하기 위함으로 사용됨
	cli  userpb.UserClient
)

func GetUserClient(serviceHost string) userpb.UserClient {
	once.Do(func() {
		// grpc.Dial: gRPC 서비스와의 커넥션을 생성해주는 함수
		// 첫번째 인자: 커넥션 맺을 타겟정보
		// 그 외 인자: 커넥션을 맺을때 필요한 추가적인 옵션
		conn, _ := grpc.Dial(serviceHost,
			grpc.WithInsecure(), // transport security 비활성화(통신 시 보안 X)
			grpc.WithBlock(), // 커넥션이 맺어지기 전까지 블록
		)

		cli = userpb.NewUserClient(conn)
	})

	return cli
}