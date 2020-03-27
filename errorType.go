package main

import (
	"errors"
	"fmt"
)

var (
	ErrInvalid    = errors.New("invalid argument")
	ErrPermission = errors.New("permission denied")
	ErrExist      = errors.New("file already exists")
	ErrNotExist   = errors.New("file does not exist")
	ErrClosed     = errors.New("file already closed")
)

func main() {
	err := createErrClosed()
	if IsNotExist(err) {
		fmt.Println(err)
	}

	if IsClosed(err) {
		fmt.Println(err)
	}
}

func createErrClosed() error {
	return ErrClosed
}

func IsInvalid(e error) bool {
	switch e {
	case ErrInvalid:
		return true
	default:
		return false
	}
}

func IsPermission(e error) bool {
	switch e {
	case ErrPermission:
		return true
	default:
		return false
	}
}

func IsExist(e error) bool {
	switch e {
	case ErrExist:
		return true
	default:
		return false
	}
}

func IsNotExist(e error) bool {
	switch e {
	case ErrNotExist:
		return true
	default:
		return false
	}
}

func IsClosed(e error) bool {
	switch e {
	case ErrClosed:
		return true
	default:
		return false
	}
}
