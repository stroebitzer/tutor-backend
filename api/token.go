package api

import (
	"fmt"

	"github.com/stroebitzer/tutor-backend/app"
)

func verifyToken(token string) error {
	if app.GetToken() != token {
		return fmt.Errorf("invalid token %s", token)
	}
	return nil
}
