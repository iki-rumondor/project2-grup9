// package middleware

// import (
// 	"time"

// 	"github.com/dgrijalva/jwt-go"
// )

// var jwtKey = []byte("rahasia-super-rahasia") // Replace with a secure secret key

// // Claims is the data structure embedded in the JWT token
// type Claims struct {
// 	UserID   uint   `json:"user_id"`
// 	Username string `json:"username"`
// 	jwt.StandardClaims
// }

// // GenerateToken is used to create a JWT token
// func GenerateToken(userID uint, username string) (string, error) {
// 	// Prepare the Claims structure
// 	claims := &Claims{
// 		UserID:   userID,
// 		Username: username,
// 		StandardClaims: jwt.StandardClaims{
// 			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(), // Token expires in 24 hours
// 		},
// 	}

// 	// Create a token with HS256 (HMAC with SHA-256) and Claims
// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

// 	// Sign the token with the secret key
// 	tokenString, err := token.SignedString(jwtKey)
// 	if err != nil {
// 		return "", err
// 	}

// 	return tokenString, nil
// }

// // IsTokenValid is used to verify a JWT token
// func IsTokenValid(tokenString string) bool {
// 	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
// 		return jwtKey, nil
// 	})
// 	if err != nil || !token.Valid {
// 		return false
// 	}
// 	return true
// }
