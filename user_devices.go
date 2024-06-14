package models

import "github.com/lib/pq"

type UserDevicesAndServices struct {
	User
	Devices   pq.StringArray `gorm:"column:devices;type:varchar[]" json:"devices"`
	IsAllowed bool           `gorm:"column:isAllowed" json:"isAllowed"`
}

func (UserDevicesAndServices) TableName() string {
	return `"public"."users"`
}
