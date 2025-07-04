package models

import (
	"time"
)

type User struct {
	ID            uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Username      string    `gorm:"size:20;not null" json:"username"`
	Password      string    `gorm:"size:100;not null" json:"password"`
	ActivationKey string    `gorm:"size:30" json:"activation_key,omitempty"`
	ResetKey      string    `gorm:"size:30" json:"reset_key,omitempty"`
	Email         string    `gorm:"size:50;unique;not null" json:"email"`
	Activated     bool      `gorm:"not null;default:false" json:"activated"`
	Created       time.Time `gorm:"not null" json:"created"`
	Updated       time.Time `gorm:"not null" json:"updated"`
}

type UserSummary struct {
	ID        uint      `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Activated bool      `json:"activated"`
	Created   time.Time `gorm:"not null" json:"created"`
	Updated   time.Time `gorm:"not null" json:"updated"`
}

type CreateUserInput struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
}

type UpdatePasswordUserInput struct {
	Pasword string `json:"password" binding:"required"`
}

type UpdateActivationKeyUserInput struct {
	ActivationKey string `json:"activation_key" binding:"required"`
}

type UpdateResetKeyUserInput struct {
	ResetKey string `json:"reset_key" binding:"required"`
}

type UpdateActivatedUserInput struct {
	Activated bool `json:"activated"`
}

type UpdateUserInput struct {
	ID       uint   `json:"id"`
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
}

type UserWithRegistrationStatus struct {
	ID         uint   `json:"id" gorm:"column:id"`
	Username   string `json:"username" gorm:"column:username"`
	Email      string `json:"email" gorm:"column:email"`
	Registered bool   `json:"registered" gorm:"column:registered"`
}

type UserPermissionSystemRequest struct {
	//News    []IncomingRole         `json:"news"`
	Edits []UserPermissionWithRegistrationStatus `json:"edits"` // o puedes ignorar si no las usas aún
	//Deletes []uint `json:"deletes"`
	//Extra   map[string]interface{} `json:"extra"`
}

type UserPermissionWithRegistrationStatus struct {
	PermissionID uint   `json:"id" gorm:"column:id"`
	Name         string `json:"name" binding:"required"`
	Registered   bool   `json:"registered" gorm:"column:registered"`
}
