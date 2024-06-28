package models

import "github.com/google/uuid"

type User struct {
	ID   uuid.UUID  `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Name string     `gorm:"type:varchar(100)"`
	Email string    `gorm:"type:varchar(100)"`
	Password string `gorm:"type:varchar(100)"`
}