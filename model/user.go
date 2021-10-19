package model

type User struct {
	ID        uint   `json:"id" gorm:"primarykey"`
	Username  string `json:"user_nsame" gorm:"not null;unique" `
	Password  string `json:"password" gorm:"not null" `
	CreatedAt int64  `gorm:"autoUpdateTime:milli" json:"created_at"`
	UpdatedAt int64  `gorm:"autoUpdateTime:milli" json:"updated_at"`
}
