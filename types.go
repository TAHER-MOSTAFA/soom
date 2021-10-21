package main

type User struct {
	Name  string `json:"name"`
	Offer string `json:"offer"`
}

type Room struct {
	ID    string
	Users []*User
}

type Msg struct {
	Type    string
	Content interface{}
}
