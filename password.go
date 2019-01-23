package password_service

// Creator describes operations with Password
type Creator interface {
	Create(lettersLen, specChLen, numbersLen int) Password
}

// Password describes password instance.
type Password string
