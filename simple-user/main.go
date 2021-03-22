import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"go-grpc\data"
	userpb "go-grpc\protos\v1\user"
)

const port = "9000"

type userServer struct {
	userpb.UserServer
}

// userId 별로 user message 리턴
func (s *userServer) GetUser (ctx context.Context, req *userpb.GetUserRequest) (*userpb.GetUserResponse, error){
	userID := req.UserId

	var UserMessage *userpb.UserMessage

	for _, u := range data.Users {
		if u.UserId != userID {
			continue
		}
		UserMessage = u
		break
	}

	return &userpb.GetUserResponse{
		UserMessage: userMessage,
	}, nil

}

// user message 리턴
func (s *userServer) ListUsers(ctx context.Context, req *userpb.ListUsersRequest) (*userpb.ListUsersResponse, error) {
	userMessages := make([]*userpb.UserMessage, len(data.Users))
	for i, u := range data.Users {
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
	userpb.RegisterUserServer(grpcServer, &userServer{})

	log.Printf("start gRPC server on %s port", portNumber)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}