package db

type Error string

func (e Error) Error() string { return string(e) }

const (
	ErrAlreadyExists = Error("already exists")
	ErrDoesntExist   = Error("doesn't exist")
)
