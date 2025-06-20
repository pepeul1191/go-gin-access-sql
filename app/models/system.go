package models

import "time"

type System struct {
	ID          uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Name        string    `gorm:"size:40;not null" json:"name"`
	Description string    `gorm:"type:text" json:"description"`
	Repository  string    `gorm:"size:100" json:"repository"`
	Created     time.Time `gorm:"not null" json:"created"`
	Updated     time.Time `gorm:"not null" json:"updated"`
}

type SystemUsersCreateRequest struct {
	//News    []IncomingPermission   `json:"news"`
	Edits []UserWithRegistrationStatus `json:"edits"` // o puedes ignorar si no las usas aún
	//Deletes []uint                       `json:"deletes"`
	//Extra   map[string]interface{}       `json:"extra"`
}
