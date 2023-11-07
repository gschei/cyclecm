package club

import (
	"errors"
)

var (
	ErrClubNotFound    = errors.New("the club was not found in the repository")
	ErrFailedToAddClub = errors.New("failed to add the club to the repository")
	ErrUpdateClub      = errors.New("failed to update the club in the repository")
)

type ClubRepository interface {
	Get(int64) (Club, error)
	Add(Club) (Club, error)
	Update(Club) error
	GetAll() []Club
}
