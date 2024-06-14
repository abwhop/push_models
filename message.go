package models

import (
	"context"
	"database/sql/driver"
	"encoding/json"
	"github.com/lib/pq"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

type Message struct {
	Id             string            `json:"id"`
	Data           map[string]string `json:"data"`
	Type           string            `json:"type"`
	Title          string            `json:"title"`
	Source         string            `json:"source"`
	Message        string            `json:"message"`
	LoginAd        string            `json:"login_ad"`
	PersonalNumber string            `json:"personalNumber"`
	ObjectId       string            `json:"objectId"`
	UserId         int               `json:"user_id"`
	MessageDetail  string            `json:"messageDetail"`
}

func (u *Message) GormValue(_ context.Context, _ *gorm.DB) clause.Expr {
	if user, err := json.Marshal(u); err == nil {
		return clause.Expr{
			SQL:  "?",
			Vars: []interface{}{string(user)},
		}
	} else {
		return clause.Expr{
			SQL:  "?",
			Vars: []interface{}{nil},
		}
	}
}

func (u *Message) Value() (driver.Value, error) {
	var user []byte
	var err error
	if user, err = json.Marshal(u); err != nil {
		return nil, err
	}
	return string(user), nil
}
func (u *Message) Scan(v interface{}) error {
	return json.Unmarshal(v.([]byte), &u)
}

type UserPushMessage struct {
	ID        string         `gorm:"column:id;type:uuid;primaryKey"`
	UserID    string         `gorm:"column:userId;notNull;index"`
	Message   *Message       `gorm:"column:message;type:jsonb;notNull;default:'{}'"`
	Devices   pq.StringArray `gorm:"column:devices;type:varchar[];notNull;default:'{}'"`
	IsViewed  bool           `gorm:"column:isViewed;type:boolean;notNull;default:false"`
	CreatedAt time.Time      `gorm:"column:createdAt;type:timestamp with time zone;notNull;default:now()"`
	UpdatedAt time.Time      `gorm:"column:updatedAt;type:timestamp with time zone;notNull;default:now()"`
}

func (UserPushMessage) TableName() string {
	return `"public"."userPushMessages"`
}

// AccessUserPushMessages Model message for access from mobile app
type AccessUserPushMessages struct {
	Id        string    `gorm:"column:id;primaryKey" json:"id"`
	IsViewed  bool      `gorm:"column:isViewed" json:"isViewed"`
	Type      string    `gorm:"column:type" json:"type"`
	Title     string    `gorm:"column:title" json:"title"`
	Source    string    `gorm:"column:source" json:"source"`
	ObjectID  string    `gorm:"column:objectID" json:"objectID"`
	Body      string    `gorm:"column:body" json:"body"`
	CreatedAt time.Time `gorm:"column:createdAt" json:"createdAt"`
}

func (AccessUserPushMessages) TableName() string {
	return `public."userPushMessages"`
}
