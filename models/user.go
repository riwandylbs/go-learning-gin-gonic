package models

import (
	"time"
)

type User struct {
	Id                 uint64    `gorm:"primaryKey;autoIncrement:true" json:"id"`
	Name               string    `gorm:"type: varchar(50);not null" json:"name"`
	Email              *string   `gorm:"type: varchar(100);unique;index;not null" json:"email"`
	Password           string    `gorm:"type: varchar(100)" json:"-"`
	Address            string    `gorm:"type: varchar(100)" json:"address"`
	RegistrationStatus string    `gorm:"type: varchar(100);after: email" json:"registration_status"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
}
