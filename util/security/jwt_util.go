package security

import (
	"fmt"
	"go_laundry/config"
	"go_laundry/model"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWTToken(user model.UserCredential) (string, error) {
	cfg, err := config.NewConfig()
	if err != nil {
		return "", err
	}
	now := time.Now()
	end := now.Add(time.Duration(cfg.ExpirationToken) * time.Minute)

	claims := &AppClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    cfg.ApplicationName,
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(end),
		},
		Username: user.Username,
		// Role:     "",
		// Service:  []string{},
	}
	tokenJwt := jwt.NewWithClaims(cfg.JwtSigningMethod, claims)
	tokenString, err := tokenJwt.SignedString(cfg.JwtSignatureKey)
	if err != nil {
		return "", fmt.Errorf("failed create jwt token: %v", err.Error())
	}
	return tokenString, nil
}

func VerifyJWTToken(tokenString string) (jwt.MapClaims, error) {
	cfg, err := config.NewConfig()
	if err != nil {
		return nil, err
	}
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		method, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok || method != jwt.SigningMethodHS256 {
			return nil, fmt.Errorf("invalid token signin method")
		}
		return cfg.JwtSignatureKey, nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid || claims["iss"] != cfg.ApplicationName {
		return nil, fmt.Errorf("invalid token")
	}
	return claims, nil
}
