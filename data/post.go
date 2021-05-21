package data

import (
	postpb "go-grpc/protos/v1/post"
)

// 이후 DB에서 데이터를 가져오는 형식으로 변경
// 연습이니 static하게 선언
// post.proto에서 정의한 PostMessage 필드와 매핑되는 값들임
type PostData struct {
	UserID string
	Posts []*postpb.PostMessage
}

// Author : post 서비스는 앞서 만든 user 서비스에 GetUser rpc로 user id를 받아서 채움
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
