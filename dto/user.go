package dto

// 用户数据类型
type User struct {
	Id        int    `json:"id"`
	Username  string `json:"username"`
	Token     string `json:"token"`
	CreatedAt int    `json:"createdAt"`
	UpdatedAt int    `json:"updatedAt"`
}
