package utils

import "github.com/dgrijalva/jwt-go"

var secretKey = "mostSecret"

func GenerateToken(data map[string]interface{}) (string, error) {
	claims := jwt.MapClaims(data)
	parseToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := parseToken.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
