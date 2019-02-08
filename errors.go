package whatsapp

import "errors"

var (
	ErrAlreadyConnected = errors.New("already connected")
	ErrAlreadyLoggedIn  = errors.New("already logged in")
	ErrInvalidSession   = errors.New("invalid session")
	ErrLoginInProgress  = errors.New("login or restore already running")
	ErrNotConnected     = errors.New("not connected")
	ErrInvalidWsData    = errors.New("received invalid data")
)
