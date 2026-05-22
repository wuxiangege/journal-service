package model

import "time"

type User struct {
	ID           uint64    `gorm:"primaryKey;autoIncrement;comment:用户ID"`
	Username     string    `gorm:"size:128;uniqueIndex;not null;comment:登录账号（邮箱或用户名）"`
	PasswordHash string    `gorm:"size:255;not null;comment:密码哈希（bcrypt）"`
	CreatedAt    time.Time `gorm:"autoCreateTime;comment:创建时间"`
}

func (User) TableName() string {
	return "users"
}
