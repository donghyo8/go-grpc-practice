package data

import (
	userpb "go-grpc/protos/v1/user"
)

// 이후 DB에서 데이터를 가져오는 형식으로 변경
// 연습이니 static하게 선언
// user.proto에서 정의한 UserMessage 필드와 매핑되는 값들임
var UserData = []*userpb.UserMessage{
	{
		UserId: "1",
		Name: "donghyo1",
		PhoneNumber: "01012345678",
		Age: 30,
	},
	{
		UserId: "2",
		Name: "donghyo2",
		PhoneNumber: "01912345678",
		Age: 29,
	},
	{
		UserId: "3",
		Name: "donghyo3",
		PhoneNumber: "01712345678",
		Age: 29,
	},
}