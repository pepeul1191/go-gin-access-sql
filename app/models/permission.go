package models

import (
	"time"
)

type Permission struct {
	ID      uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Name    string    `gorm:"size:40;not null" json:"name"`
	Created time.Time `gorm:"not null" json:"created"`
	Updated time.Time `gorm:"not null" json:"updated"`
	RoleId  uint      `gorm:"not null" json:"role_id"`                 // Hace referencia a System.ID
	Role    Role      `gorm:"foreignKey:RoleId" json:"role,omitempty"` // Relación con System
}

type PermissionCreateRequest struct {
	News    []IncomingPermission   `json:"news"`
	Edits   []Permission           `json:"edits"` // o puedes ignorar si no las usas aún
	Deletes []uint                 `json:"deletes"`
	Extra   map[string]interface{} `json:"extra"`
}

type IncomingPermission struct {
	ID   string `json:"id"` // tmp_...
	Name string `json:"name"`
}

type CreatedPermissionResponse struct {
	Tmp string `json:"tmp"`
	ID  string `json:"id"`
}

type ExtPermission struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}
