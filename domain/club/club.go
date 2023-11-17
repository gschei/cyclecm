package club

import (
	"errors"
)

var (
	ErrInvalidClub = errors.New("club must have a name")
)

type Club struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

func NewClub(name string) (*Club, error) {
	if name == "" {
		return &Club{}, ErrInvalidClub
	}

	return &Club{
		Name: name,
		ID:   0,
	}, nil
}
