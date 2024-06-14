package models

import "time"

type UserPushService struct {
	UserId      string    `gorm:"column:userId;primaryKey" json:"userId"`
	ServiceCode string    `gorm:"column:serviceCode;primaryKey" json:"serviceCode"`
	IsAllowed   bool      `gorm:"column:isAllowed;" json:"isAllowed"`
	CreatedAt   time.Time `gorm:"column:createdAt;" json:"createdAt"`
	UpdatedAt   time.Time `gorm:"column:updatedAt;" json:"updatedAt"`
}

func (UserPushService) TableName() string {
	return `public."userPushServices"`
}
