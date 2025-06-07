package models

// Modelo para la tabla "foots"
type Foot struct {
	ID   uint   `gorm:"primaryKey;autoIncrement"`
	Name string `gorm:"size:10;not null"`
}
