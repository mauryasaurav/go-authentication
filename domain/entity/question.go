package entity

import (
	"time"

	"github.com/google/uuid"
)

/* Question Schema */
type QuestionSchema struct {
	Id         uuid.UUID `gorm:"primaryKey;type:uuid;default:public.uuid_generate_v4()" json:"id"`
	CreatedAt  time.Time `json:"created_at" gorm:"default=now()"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"default=now()"`
	Question   string    `json:"question"`
	Answer     string    `json:"answer"`
	UserID     uuid.UUID `json:"user_id"`
	CategoryID uuid.UUID `json:"category_id"`
	Role      string    `json:"role" gorm:"default=user"`
}
