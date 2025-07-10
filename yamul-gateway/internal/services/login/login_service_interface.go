package login

import "yamul-gateway/internal/dtos/commands"

// LoginServiceInterface define la interfaz para el servicio de login.
type LoginServiceInterface interface {
	CheckUserCredentials(username string, password string) (bool, commands.LoginDeniedReason)
}
