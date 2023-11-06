package jwt

import (
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
)

var (
	secretKey = os.Getenv("JWT_SECRET") // JWT secret key
)

// GenerateJwtToken generates a JWT token.
func GenerateJwtToken(payload jwt.Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	return token.SignedString([]byte(secretKey))
}

// TokenData struct contains the data required to create a JWT token.
type TokenData struct {
	LoginUser   string
	UserID      int
	Roles       []string
	EmpNO       string
	Avatar      string
	AccessToken string
}

// CreateJwtToken creates a JWT token and returns the generated token string.
func CreateJwtToken(data TokenData) (string, error) {
	durTime := os.Getenv("JWT_EXPIRATION")
	if durTime == "" {
		durTime = "2h" // Default JWT expiration time is 2 hours
	}

	expiration, err := time.ParseDuration(durTime)
	if err != nil {
		return "", err
	}

	now := time.Now()
	expiresAt := now.Add(expiration)

	return GenerateJwtToken(&PayloadClaims{
		StandardClaims: jwt.StandardClaims{
			Audience:  strconv.Itoa(data.UserID),     // Audience
			ExpiresAt: expiresAt.Unix(),              // Expiration time
			IssuedAt:  now.Unix(),                    // Issuance time
			Issuer:    "Golang Project",              // Issuer
			Subject:   "Web JWT Token - API Service", // Subject
		},
		Username:    data.LoginUser,
		EmpNO:       data.EmpNO,
		Profile:     data.Avatar,
		Certificate: data.AccessToken,
		Roles:       data.Roles,
	})
}

// ParseJwtToken parses the JWT token and returns the PayloadClaims struct.
func ParseJwtToken(jwtToken string) (*PayloadClaims, error) {
	tokenClaims, err := jwt.ParseWithClaims(jwtToken, &PayloadClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*PayloadClaims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}

// ValidateJwt validates the JWT token's validity and returns the PayloadClaims struct if valid.
func ValidateJwt(jwtToken string) (*PayloadClaims, error) {
	claims, err := ParseJwtToken(jwtToken)
	if err != nil {
		return nil, err
	}
	if err = claims.Valid(); err != nil {
		return nil, err
	}
	return claims, nil
}
