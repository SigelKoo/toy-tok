package gorm_model

import (
	"time"
)

const (
	userInfo = "user_info"
)

type UserInfo struct {
	ID            int64     `gorm:"column:id"`
	Name          string    `gorm:"column:name"`
	UserName      string    `gorm:"column:user_name"`
	Password      string    `gorm:"column:password"`
	FollowCount   int64     `gorm:"column:follow_count"`
	FollowerCount int64     `gorm:"column:follower_count"`
	CreatedAt     time.Time `gorm:"column:created_at"`
	UpdatedAt     time.Time `gorm:"column:updated_at"`
}

func (*UserInfo) TableName() string {
	return userInfo
}
