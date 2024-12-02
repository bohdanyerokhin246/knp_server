package config

import "github.com/golang-jwt/jwt/v4"

type User struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	Role     string `json:"role,omitempty"`
}

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Claims struct {
	//Username string `json:"username,omitempty"`
	Role string `json:"role,omitempty"`
	jwt.RegisteredClaims
}
