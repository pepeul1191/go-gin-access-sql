package models

import (
	"time"
)

type Role struct {
	ID       uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Name     string    `gorm:"size:40;not null" json:"name"`
	Created  time.Time `gorm:"not null" json:"created"`
	Updated  time.Time `gorm:"not null" json:"updated"`
	SystemID uint      `gorm:"not null" json:"system_id"`                   // Hace referencia a System.ID
	System   System    `gorm:"foreignKey:SystemID" json:"system,omitempty"` // Relación con System
}

type RoleCreateRequest struct {
	News    []IncomingRole         `json:"news"`
	Edits   []Role                 `json:"edits"` // o puedes ignorar si no las usas aún
	Deletes []uint                 `json:"deletes"`
	Extra   map[string]interface{} `json:"extra"`
}

type IncomingRole struct {
	ID   string `json:"id"` // tmp_...
	Name string `json:"name"`
}

type CreatedRoleResponse struct {
	Tmp string `json:"tmp"`
	ID  string `json:"id"`
}
