// // auth.go
// package middleware

// import (
// 	"errors"
// 	"net/http"

// 	"github.com/dgrijalva/jwt-go"
// 	"github.com/gin-gonic/gin"
// )

// type YourClaimStruct struct {
// 	UserID   uint   `json:"user_id"`
// 	Username string `json:"username"`
// 	jwt.StandardClaims
// }

// // AuthMiddleware is middleware for checking authentication tokens.
// func AuthMiddleware() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		// Get the token from the Authorization header
// 		token := c.GetHeader("Authorization")

// 		if token == "" {
// 			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token not found"})
// 			c.Abort()
// 			return
// 		}

// 		// Verify the token using the IsTokenValid function provided in jwt.go
// 		if !IsTokenValid(token) {
// 			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token not valid"})
// 			c.Abort()
// 			return
// 		}

// 		// Token is valid, proceed to the next handler
// 		c.Next()
// 	}
// }

// // ...

// // ParseToken is used to parse and verify JWT tokens
// func ParseToken(authorizationHeader string) (*YourClaimStruct, error) {
// 	token, err := jwt.ParseWithClaims(authorizationHeader, &YourClaimStruct{}, func(token *jwt.Token) (interface{}, error) {
// 		return jwtKey, nil
// 	})
// 	if err != nil {
// 		return nil, err
// 	}

// 	claims, ok := token.Claims.(*YourClaimStruct)
// 	if !ok || !token.Valid {
// 		return nil, errors.New("Token not valid")
// 	}

// 	return claims, nil
// }
