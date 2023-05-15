package owl

import "errors"

type Repository interface {
	Get() ([]Match, error)
	GetActive() ([]ActiveMatch, error)
}

var ErrNotFound = errors.New("not found")
