package model

import "time"

type User struct {
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Varifed   bool      `json:"varifed"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAT time.Time `json:"UpdatedAt"`
}
