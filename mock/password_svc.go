package mock

import "github.com/valentyn88/password_service"

// PasswordSvc service for test.
type PasswordSvc struct {
}

// Create makes fake password.
func (ps PasswordSvc) Create(lettersLen, specChLen, numbersLen int) password_service.Password {
	return password_service.Password("8G628:z0&")
}
