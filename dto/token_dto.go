package dto

import (
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

type AccessTokenDTO struct {
	UserId      uuid.UUID `json:"user_id"`
	Permissions []string  `json:"permissions"`
	jwt.StandardClaims
}

type RefreshTokenDTO struct {
	UserId uuid.UUID `json:"user_id"`
	jwt.StandardClaims
}