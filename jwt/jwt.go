package jwt

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"

	"github.com/draco121/common/models"
)

// JWTSecretKey is a secret key used to sign the JWT tokens.
var JWTSecretKey = []byte(os.Getenv("JWT_SECRET"))

// GenerateJWT creates a new JWT token with default claims.
func GenerateJWT(customClaims *models.JwtCustomClaims) (string, error) {
	claims := models.DefaultClaims{
		JwtCustomClaims: *customClaims,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 1).Unix(), // Token expires in 1 hour
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(JWTSecretKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

// GenerateRefreshToken creates a new refresh token.
func GenerateRefreshToken(sessionId string) (string, error) {
	claims := models.RefreshTokenClaims{
		SessionId: sessionId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24 * 7).Unix(), // Refresh token expires in 7 days
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	refreshToken, err := token.SignedString(JWTSecretKey)
	if err != nil {
		return "", err
	}

	return refreshToken, nil
}

func VerfiyRefreshToken(refreshToken string) (*models.RefreshTokenClaims, error) {
	// Validate the refresh token and get its claims
	token, err := jwt.ParseWithClaims(refreshToken, &models.RefreshTokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return JWTSecretKey, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*models.RefreshTokenClaims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("invalid refresh token")
	}

	return claims, nil
}

// GenerateNewToken generates a new JWT token based on a refresh token.
func VerifyJwtToken(jwtToken string) (*models.DefaultClaims, error) {
	// Validate the refresh token and get its claims
	token, err := jwt.ParseWithClaims(jwtToken, &models.DefaultClaims{}, func(token *jwt.Token) (interface{}, error) {
		return JWTSecretKey, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*models.DefaultClaims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("invalid jwt token")
	}

	return claims, nil
}
