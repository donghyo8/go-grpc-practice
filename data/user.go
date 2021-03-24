package data

import (
	userpb "go-grpc/protos/v1/user"
)

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