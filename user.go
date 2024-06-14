package models

import "time"

type User struct {
	Id                 string    `gorm:"column:id;primaryKey" json:"id"`
	BitrixUserId       int       `gorm:"column:bitrixUserId;" json:"bitrixUserId"`
	FullPersonalNumber string    `gorm:"column:fullPersonalNumber" json:"fullPersonalNumber"`
	LoginAd            string    `gorm:"column:loginAd" json:"loginAd"`
	CreatedAt          time.Time `gorm:"column:createdAt" json:"createdAt"`
	UpdatedAt          time.Time `gorm:"column:updatedAt" json:"updatedAt"`
}

func (User) TableName() string {
	return "public.users"
}
