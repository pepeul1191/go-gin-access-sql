package models

import "github.com/golang-jwt/jwt/v5"

type Claims struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	UserID   uint   `json:"user_id"`
	SystemID uint   `json:"system_id"`
	jwt.RegisteredClaims
}
