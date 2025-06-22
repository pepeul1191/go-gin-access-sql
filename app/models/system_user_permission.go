package models

import (
	"time"
)

type SystemUserPermission struct {
	Created      time.Time  `gorm:"not null" json:"created"`
	SystemID     uint       `gorm:"not null;index"`          // Clave foránea para System
	UserID       uint       `gorm:"not null;index"`          // Clave foránea para User
	PermissionID uint       `gorm:"not null;index"`          // Clave foránea para Permission
	System       System     `gorm:"foreignKey:SystemID"`     // Relación opcional (belongs to)
	User         User       `gorm:"foreignKey:UserID"`       // Relación opcional (belongs to)
	Permission   Permission `gorm:"foreignKey:PermissionID"` // Relación opcional (belongs to)
}

func (SystemUserPermission) TableName() string {
	return "systems_users_permissions"
}
