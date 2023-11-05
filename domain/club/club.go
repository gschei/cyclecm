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

func NewClub(name string) (Club, error) {
	if name == "" {
		return Club{}, ErrInvalidClub
	}

	return Club{
		Name: name,
		ID:   1,
	}, nil
}

func (c *Club) GetID() int64 {
	return c.ID
}
func (c *Club) SetID(id int64) {
	c.ID = id
}

func (c *Club) SetName(name string) {
	c.Name = name
}

func (c *Club) GetName() string {
	return c.Name
}
