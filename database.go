package main

type ClientProfile struct {
	Email string
	Id    string
	Name  string
	Token string
}

var database = map[string]ClientProfile{
	"user1": {
		Email: "example@mail.com",
		Id:    "user1",
		Name:  "John Doe",
		Token: "token1",
	},
	"user2": {
		Email: "rose@mail.com",
		Id:    "user2",
		Name:  "Rose Doe",
		Token: "token2",
	},
}
