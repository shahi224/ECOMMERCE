package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"ECOMMERCE/utils/helper"
	"ECOMMERCE/utils/response"
)

func UserAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get token from header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
			return
		}

		// Strip "Bearer " prefix
		tokenString := helper.GetTokenFromHeader(authHeader)

		// Optional: support cookie fallback
		if tokenString == "" {
			var err error
			tokenString, err = c.Cookie("Authorization")
			if err != nil {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token missing"})
				return
			}
		}

		// ✅ Extract user data from JWT
		userID, userEmail, userRole, err := helper.ExtractUserIDFromToken(tokenString)
		if err != nil {
			fmt.Println("JWT Error:", err)
			res := response.ClientResponse(http.StatusUnauthorized, "Invalid Token", nil, err.Error())
			c.AbortWithStatusJSON(http.StatusUnauthorized, res)
			return
		}

		if userID == 0 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid user ID in token"})
			return
		}

		// ✅ Set user info in Gin context
		c.Set("user_id", userID)
		c.Set("user_email", userEmail)
		c.Set("user_role", userRole)

		c.Next()
	}
}
