package custommodels

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type TokenUser struct {
	Data           interface{}
	StandardClaims jwt.StandardClaims
}

// Valid implements jwt.Claims
func (t *TokenUser) Valid() error {
	now := time.Now().Unix()
	fmt.Println("entro")
	if now > t.StandardClaims.ExpiresAt {
		return fmt.Errorf("Token has expired")
	}
	return nil
}
