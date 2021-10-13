package dto

type User struct {
	Id         int    `json:"id"`
	Username   string `json:"username"`
	Token      string `json:"token"`
	Created_at int    `json:"created_at"`
	Updated_at int    `json:"updated_at"`
}
