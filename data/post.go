package data


import (
	postpb "go-grpc/protos/v1/post"
)

type PostData struct {
	UserID string
	Posts []*postpb.PostMessage
}

var UserPosts = []*PostData {
	{
		UserID: "1",
		Posts: []*postpb.PostMessage {

			{
				PostId: "1",
				Author: "",
				Title: "gRPC 구축(1)",
				Body: "gRPC 구축은 이렇게",
				Tags: []string{"gRPC", "Golang", "Server", "protobuf"},
			},

			{
				PostId: "12",
				Author: "",
				Title: "gRPC 구축(2)",
				Body: "gRPC 구축은 이렇게",
				Tags: []string{"gRPC", "Golang", "Server", "protobuf"},
			},
		},
	},
}
