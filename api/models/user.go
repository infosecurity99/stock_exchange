package models

import "time"

type User struct {
	ID           string    `json:"id"`
	Name         string    `json:"full_name"`
	Phone        string    `json:"phone"`
	Email        string    `json:"password"`
	PasswordHash string    `json:"cash"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    string    `json:"updated_at"`
	Login        string    `json:"login"`
}

type CreateUser struct {
	Name         string `json:"full_name"`
	Phone        string `json:"phone"`
	PasswordHash string `json:"password"`
	Email        string `json:"cash"`
	Login        string `json:"login"`
}

type UpdateUser struct {
	ID           string `json:"-"`
	Name         string `json:"full_name"`
	Phone        string `json:"phone"`
	PasswordHash string `json:"password"`
	Email        string `json:"cash"`
}

type UsersResponse struct {
	Users []User `json:"users"`
	Count int    `json:"count"`
}

type UpdateUserPassword struct {
	ID          string `json:"-"`
	NewPassword string `json:"new_password"`
	OldPassword string `json:"old_password"`
}
