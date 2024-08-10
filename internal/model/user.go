package model

import "github.com/google/uuid"

type User struct { // User logged In
	ID       uuid.UUID `json:"id"`
	Fullname string    `json:"fullname"`
	Email    string    `json:"email"`
	Phone    string    `json:"phone"`
}
