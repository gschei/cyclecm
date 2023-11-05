package ride

import (
	"time"

	"github.com/gschei/cyclecm/domain/rider"
)

type ride struct {
	ID           string       `json:"id"`
	Name         string       `json:"name"`
	StartDate    time.Time    `json:"startDate"`
	Route        string       `json:"route"`
	Participants *rider.Rider `json:"participants"`
}
