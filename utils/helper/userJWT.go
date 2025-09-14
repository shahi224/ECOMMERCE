package helper

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"ECOMMERCE/config"
)

// Struct for custom claims
type AuthUserClaims struct {
	UserID uint   `json:"user_id"`
	Email  string `json:"email"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

// ✅ GenerateJWT creates a new JWT token with user ID, email, role, and 24h expiry
func GenerateJWT(userID uint, email string, role string) (string, error) {
	cfg, err := config.LoadConfig()
	if err != nil {
		return "", fmt.Errorf("config error: %w", err)
	}

	claims := AuthUserClaims{
		UserID: userID,
		Email:  email,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), // Token expires in 24 hours
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign and return token string
	return token.SignedString([]byte(cfg.Key))
}

// ✅ GetTokenFromHeader extracts the token string from the Authorization header
func GetTokenFromHeader(header string) string {
	if len(header) >= 7 && header[:7] == "Bearer " {
		return header[7:]
	}
	return header
}

// ✅ ExtractUserIDFromToken validates token and extracts user ID, email, and role
func ExtractUserIDFromToken(tokenString string) (uint, string, string, error) {
	cfg, err := config.LoadConfig()
	if err != nil {
		return 0, "", "", fmt.Errorf("config error: %w", err)
	}

	// Parse the token
	token, err := jwt.ParseWithClaims(tokenString, &AuthUserClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Validate signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(cfg.Key), nil
	})
	if err != nil {
		return 0, "", "", err
	}

	// Type assertion to custom claims
	claims, ok := token.Claims.(*AuthUserClaims)
	if !ok || !token.Valid {
		return 0, "", "", errors.New("invalid token or claims")
	}

	return claims.UserID, claims.Email, claims.Role, nil
}
