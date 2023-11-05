package memory

import (
	"sync"

	"github.com/gschei/cyclecm/domain/club"
)

var idCounter int64 = 0

type MemoryRepository struct {
	clubs map[int64]club.Club
	sync.Mutex
}

func New() *MemoryRepository {
	return &MemoryRepository{
		clubs: make(map[int64]club.Club),
	}
}

func (mr *MemoryRepository) Get(id int64) (club.Club, error) {
	if club, ok := mr.clubs[id]; ok {
		return club, nil
	}

	return club.Club{}, club.ErrClubNotFound
}

func (mr *MemoryRepository) Add(c club.Club) (club.Club, error) {
	mr.Lock()
	idCounter++
	c.ID = idCounter
	mr.clubs[c.GetID()] = c
	mr.Unlock()
	return c, nil
}

func (mr *MemoryRepository) Update(c club.Club) error {
	if _, ok := mr.clubs[c.GetID()]; !ok {
		return club.ErrClubNotFound
	}
	mr.Lock()
	mr.clubs[c.GetID()] = c
	mr.Unlock()
	return nil
}
