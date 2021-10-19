package dto

// 用户数据类型
type User struct {
	Id        int    `json:"id"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Token     string `json:"token"`
	CreatedAt int64  `gorm:"autoUpdateTime:milli" json:"created_at"`
	UpdatedAt int64  `gorm:"autoUpdateTime:milli" json:"updated_at"`
}
