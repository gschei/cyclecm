package entity

import (
	"github.com/gschei/cyclecm/domain/club"
	"github.com/gschei/cyclecm/domain/rider"
)

type membership struct {
	ID         string       `json:"id"`
	Club       *club.Club   `json:"club"`
	Rider      *rider.Rider `json:"rider"`
	FeeBalance int          `json:"feeBalance"`
}
