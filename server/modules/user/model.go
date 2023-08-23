package user

import (
	"database/sql"
	"time"
)

type User struct {
	ID          uint         `gorm:"primaryKey" json:"id"`
	FirstName   string       `json:"firstName"`
	LastName    string       `json:"lastName"`
	Email       *string      `gorm:"unique;not null" json:"email" validate:"required,email"`
	Password    string       `gorm:"not null" json:"-" validate:"required,min=8"`
	ActivatedAt sql.NullTime `json:"activatedAt"`
	CreatedAt   time.Time    `json:"createdAt"`
	UpdatedAt   time.Time    `json:"updatedAt"`
}

type SignUp struct {
	FirstName   string       `json:"firstName"`
	LastName    string       `json:"lastName"`
	Email       *string      `json:"email" validate:"required,email"`
	Password    string       `json:"password" validate:"required,min=8"`
	ActivatedAt sql.NullTime `json:"activatedAt"`
	CreatedAt   time.Time    `json:"createdAt"`
	UpdatedAt   time.Time    `json:"updatedAt"`
}

type SignIn struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type SignedInUser struct {
	Token string
	User
}
