package simple_client_server

import (
	"sync"
	"log"
	"flag"

	"google.golang.org/grpc"
	"golang.org/x/oauth2"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/oauth"
	"google.golang.org/grpc/examples/data"

	userpb "go-grpc/protos/v1/user"
)

func GetUserClientAuth(){


}

func fetchToken() *oauth2.Token {
	return &oauth2.Token{
		AccessToken: "some-secret-token",
	}
}