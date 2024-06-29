package models

import "github.com/google/uuid"

type User struct {
	ID   uuid.UUID  `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Name string     `gorm:"type:varchar(100)"`
	Email string    `gorm:"type:varchar(100)"`
	Password string `gorm:"type:varchar(100)"`
	AccountNumber int64 `gorm:"type:bigserial;autoincrement"`
	Balance float32 `gorm:"type:numeric(10,2);default:0"`
}