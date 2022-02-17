package models

type User struct {
	Uid string `json:"uid"`
}

type JWTHeader struct {
	Authorization string `json:"Authorization"`
}
