package request

import "github.com/dgrijalva/jwt-go"

// Custom claims structure
type CustomClaims struct {
	ID       uint
	Name     string
	Username string
	Email    string
	jwt.StandardClaims
}
