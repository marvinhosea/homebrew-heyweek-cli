package models

type Authenticate struct {
	RefreshToken string
	Token        string
}

type User struct {
	Id string
}
