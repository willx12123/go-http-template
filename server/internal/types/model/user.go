package model

import "time"

type User struct {
	ID             uint `gorm:"primaryKey,autoIncrement"`
	Name           string
	Email          string
	PasswordDigest string
	CreatedAt      time.Time `gorm:"autoCreateTime"`
	UpdatedAt      time.Time `gorm:"autoCreateTime,autoUpdateTime"`
}

func (*User) TableName() string {
	return "users"
}
