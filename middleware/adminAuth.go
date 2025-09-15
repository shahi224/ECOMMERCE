package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"

	"ECOMMERCE/config"
)

func AdminAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Bypass HTML page render for login screen
		if c.Request.Method == "GET" && strings.HasPrefix(c.FullPath(), "/admin/Authentication/login") {
			c.Next()
			return
		}

		cfg, err := config.LoadConfig()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to load config"})
			c.Abort()
			return
		}

		jwtSecret := []byte(cfg.Key)
		var tokenStr string

		// Try Authorization Header first
		authHeader := c.GetHeader("Authorization")
		if authHeader != "" {
			// Split the header to handle case variations and extra spaces
			parts := strings.Fields(authHeader)
			if len(parts) == 2 && strings.EqualFold(parts[0], "bearer") {
				tokenStr = parts[1]
			} else if len(parts) == 1 {
				// Handle case where only token is provided without "Bearer"
				tokenStr = parts[0]
			} else {
				// Handle invalid header format
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Authorization header format"})
				c.Abort()
				return
			}
		} else {
			// Try Cookie
			cookie, err := c.Cookie("admin_token")
			if err == nil && cookie != "" {
				tokenStr = cookie
			}
		}

		// If no token
		if tokenStr == "" {
			// Browser request → redirect to login
			if strings.Contains(c.GetHeader("Accept"), "text/html") {
				c.Redirect(http.StatusSeeOther, "/admin/Authentication/login")
			} else {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header or session cookie missing"})
			}
			c.Abort()
			return
		}

		// Validate JWT
		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})

		if err != nil || !token.Valid {
			if strings.Contains(c.GetHeader("Accept"), "text/html") {
				c.Redirect(http.StatusSeeOther, "/admin/Authentication/login")
			} else {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			}
			c.Abort()
			return
		}

		// Extract admin ID from token and set in context
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			if userID, exists := claims["user_id"]; exists {
				if userIDFloat, ok := userID.(float64); ok {
					c.Set("adminID", uint(userIDFloat))
				}
			}
		}

		c.Next()
	}
}
