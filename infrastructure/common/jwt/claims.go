package jwt

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
)

// PayloadClaims struct represents the payload claims of a JWT.
type PayloadClaims struct {
	jwt.StandardClaims
	Username    string   `json:"name"`    // Username
	EmpNO       string   `json:"emp_no"`  // Employee number
	Profile     string   `json:"profile"` // Profile picture link
	Certificate string   `json:"cert"`    // Access token
	Roles       []string `json:"roles"`   // Roles list
}

// Valid method validates the JWT's validity.
func (c PayloadClaims) Valid() error {
	if err := c.StandardClaims.Valid(); err != nil {
		return err
	}
	if c.Username == "" {
		return errors.New("missing account name")
	}
	if len(c.Roles) == 0 {
		return errors.New("missing account role")
	}
	if c.Certificate == "" {
		return errors.New("missing certificate")
	}
	if c.ExpiresAt < time.Now().Unix() {
		return errors.New("token has expired")
	}
	return nil
}
