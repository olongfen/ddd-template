package dto

import (
	"ddd-template/domain/entities"
	"github.com/jinzhu/copier"
	"time"
)

type DemoInfo struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	Message   string    `json:"message"`
}

func DemoEnt2Dto(in entities.Demo) (res *DemoInfo, err error) {
	var (
		to = new(DemoInfo)
	)
	if err = copier.Copy(to, in); err != nil {
		return
	}
	return to, nil
}
