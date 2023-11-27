package security

import "github.com/golang-jwt/jwt/v5"

type AppClaims struct {
	jwt.RegisteredClaims
	Username string   `json:"username"`
	Role     string   `json:"role,omitempty"`
	Service  []string `json:"services,omitempty"`
}

//service -> resource/api method apa yg user bisa access(authorization)
