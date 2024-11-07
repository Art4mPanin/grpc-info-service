package jwt

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

//	func GetToken(authHeader string) (*jwt.Token, error) {
//		token, err := jwt.Parse(authHeader, func(token *jwt.Token) (interface{}, error) {
//			return []byte{}, nil
//		})
//		if err != nil {
//			return nil, fmt.Errorf("invalid token from get: %w", err)
//		}
//		return token, nil
//	}
func GetToken(authHeader string) (*jwt.Token, error) {
	token, _, err := new(jwt.Parser).ParseUnverified(authHeader, jwt.MapClaims{})
	if err != nil {
		return nil, fmt.Errorf("invalid token from get: %w", err)
	}
	return token, nil
}
func ValidateToken(token *jwt.Token) (int, bool, error) {
	claims, ok := token.Claims.(jwt.MapClaims)

	if exp, ok := claims["exp"].(float64); ok {
		if float64(time.Now().Unix()) > exp {
			return 0, false, fmt.Errorf("token expired")
		}
	} else {
		return 0, false, fmt.Errorf("invalid token claims")
	}
	subFloat, ok := claims["sub"].(float64)
	if !ok {
		return 0, false, fmt.Errorf("invalid token subject")
	}
	sub := int(subFloat)

	superuser, ok := claims["superuser"].(bool)
	if !ok {
		return 0, false, fmt.Errorf("invalid rights")
	}

	return sub, superuser, nil
}
