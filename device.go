package models

import "time"

type Device struct {
	DeviceId   string    `gorm:"column:deviceId;primaryKey" json:"deviceId"`
	UserId     string    `gorm:"column:userId;primaryKey" json:"userId"`
	DeviceType string    `gorm:"column:deviceType" json:"deviceType"`
	Token      string    `gorm:"column:token" json:"token"`
	CreatedAt  time.Time `gorm:"column:createdAt" json:"createdAt"`
	UpdatedAt  time.Time `gorm:"column:updatedAt" json:"updatedAt"`
}

func (Device) TableName() string {
	return `"public"."userDevices"`
}
