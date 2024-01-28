package models

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/draco121/common/constants"
)

type JwtCustomClaims struct {
	Email     string         `json:"email"`
	UserId    string         `json:"userId"`
	Role      constants.Role `json:"role"`
	SessionId string         `json:"sessionId"`
}

// DefaultClaims represents the default claims for the JWT token.
type DefaultClaims struct {
	JwtCustomClaims
	jwt.StandardClaims
}

// RefreshTokenClaims represents the claims for the refresh token.
type RefreshTokenClaims struct {
	SessionId string `json:"sessionId"`
	jwt.StandardClaims
}
