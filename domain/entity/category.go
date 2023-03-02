package entity

import (
	"time"

	"github.com/google/uuid"
)

/* Category Validation */
type CategorySchema struct {
	Id        uuid.UUID `gorm:"primaryKey;type:uuid;default:public.uuid_generate_v4()" json:"id"`
	CreatedAt time.Time `json:"created_at" gorm:"default=now()"`
	UpdatedAt time.Time `json:"updated_at" gorm:"default=now()"`
	Role      string    `json:"role" gorm:"default=user"`
	Name      string    `json:"name"`
	UserID    uuid.UUID `json:"user_id"`
}
