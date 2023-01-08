package service

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

type JWTService interface {
	GenerateToken(name string) string
	ValidateToken(tokenString string) (*jwt.Token, error)
}

type jwtCustomClaims struct {
	Name string `json:"name"`
	jwt.StandardClaims
}

type jwtService struct {
	secretkey string
	issuer    string
}

func NewJWTService() JWTService {
	return &jwtService{
		secretkey: getSecretKey(),
		issuer:    "riwandylbs",
	}
}

func getSecretKey() string {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		secret = "riwandylbs"
	}

	return secret
}

func (jwtSrv *jwtService) GenerateToken(username string) string {

	// set custom and standard claim
	claims := &jwtCustomClaims{
		username,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
			Issuer:    jwtSrv.issuer,
			IssuedAt:  time.Now().Unix(),
		},
	}

	// create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// generating token with secret key
	t, err := token.SignedString([]byte(jwtSrv.secretkey))
	if err != nil {
		panic(err)
	}
	return t
}

// ValidateToken implements JWTService
func (jwt *jwtService) ValidateToken(tokenString string) (*jwt.Token, error) {
	panic("unimplemented")
}
