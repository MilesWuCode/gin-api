package plugin

import (
	"errors"
)

type Pagination struct {
	Page int `form:"page"`
	Size int `form:"size"`
	Sort int `form:"sort"`
}

func (p *Pagination) Ready() (limit int, offset int, err error) {
	limit = p.Size

	offset = (p.Page - 1) * p.Size

	err = nil

	if limit < 0 || offset < 0 {
		limit = 0

		offset = 0

		err = errors.New("Pagination params error")
	}

	return
}
