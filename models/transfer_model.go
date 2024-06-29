package models

import (
	"time"

	"github.com/google/uuid"
)

type Transfer struct {
	ID            uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	FromAccountID int64  	`gorm:"type:bigserial;not null"`
	ToAccountID   int64 	`gorm:"type:bigserial;not null"`
	Amount        float32   `gorm:"type:numeric(10,2);not null"`
	CreatedAt     time.Time `gorm:"type:timestamp;default:current_timestamp"`
}