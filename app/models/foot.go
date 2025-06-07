package models

// Modelo para la tabla "foots"
type Foot struct {
	ID   uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Name string `gorm:"size:10;not null" json:"name"`
}
