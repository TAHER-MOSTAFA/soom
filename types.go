package main

type User struct {
	Name  string
	Offer string
}

type Room struct {
	ID    string
	Users []*User
}

type Msg struct {
	Type    string
	Content interface{}
}
