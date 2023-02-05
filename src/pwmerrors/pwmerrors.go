package pwmerrors

import "errors"

var (
	ErrPassExists   = errors.New("Password already exists")
	ErrPassMismatch = errors.New("Passwords aren't equal")
)
