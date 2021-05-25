package data

import (
	userpb "go-grpc/protos/v2/user"
)

var UsersV2 = []*userpb.UserMessage{
	{
		UserId:      "1",
		Name:        "user1",
		PhoneNumber: "01012341234",
		Age:         10,
	},
	{
		UserId:      "2",
		Name:        "user2",
		PhoneNumber: "01098128734",
		Age:         21,
	},
	{
		UserId:      "3",
		Name:        "user3",
		PhoneNumber: "01056785678",
		Age:         32,
	},
	{
		UserId:      "4",
		Name:        "user4",
		PhoneNumber: "01099999999",
		Age:         45,
	},
	{
		UserId:      "5",
		Name:        "user5",
		PhoneNumber: "01012344321",
		Age:         54,
	},
}
