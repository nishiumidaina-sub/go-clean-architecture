package domain

import (
    "time"
)

type Users []User

type User struct {
	ID        int       `json:"id"`
	UserKey   string    `json:"user_key"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
    UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
