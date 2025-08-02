package models

import (
	"time"
)

type SystemUser struct {
	ID       uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	SystemID uint      `gorm:"not null" json:"system_id"`
	UserID   uint      `gorm:"not null" json:"user_id"`
	Created  time.Time `gorm:"not null" json:"created"`
}

func (SystemUser) TableName() string {
	return "systems_users"
}

type SystemUserView struct {
	ID        uint   `gorm:"primaryKey;column:id" json:"id"`
	Username  string `gorm:"column:username" json:"username"`
	Password  string `gorm:"column:password" json:"-"`
	Email     string `gorm:"column:email" json:"email"`
	Activated bool   `gorm:"column:activated" json:"activated"`
}

func (SystemUserView) TableName() string {
	return "vw_system_users"
}

type ExtSystemUsersUsernameInput struct {
	SystemID uint   `json:"system_id"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type ExtSystemUsersEmailInput struct {
	SystemID uint   `json:"system_id"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type ExtSystemUsersOutput struct {
	ID       uint                     `json:"id"`
	Username string                   `json:"username"`
	Email    string                   `json:"email"`
	SystemID uint                     `json:"system_id"`
	Token    string                   `json:"token"`
	Roles    []ExtRoleWithPermissions `json:"roles"`
}
