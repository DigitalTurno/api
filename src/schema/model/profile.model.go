package model

import (
	"time"
)

type Profile struct {
	ID        int64     `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID    int64     `gorm:"not null;index" json:"userId"` // Configurar como clave foránea y agregar índice
	Firstname *string   `gorm:"type:varchar(50); null;" json:"firstname" validate:"max=50"`
	Lastname  *string   `gorm:"type:varchar(50); null;" json:"lastname" validate:"max=50"`
	Email     string    `gorm:"type:varchar(100);not null;unique" json:"email" validate:"required,email"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	// Relación con el modelo User
	User User `gorm:"foreignKey:UserID;references:ID"` // Relacionar con el modelo User
}

type ProfileInput struct {
	Firstname *string `json:"firstname" validate:"max=50"`
	Lastname  *string `json:"lastname" validate:"max=50"`
	Email     string  `json:"email" validate:"required,email"`
}
