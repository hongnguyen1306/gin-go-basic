package common

import "strings"

type Paging struct {
	Page       int    `json:"page"`
	Limit      int    `json:"limit"`
	Total      int    `json:"total"`
	FakeCursor string `json:"cursor"`
	NextCursor string `json:"nextCursor"`
}

func (p *Paging) Fulfill() {
	if p.Page <= 0 {
		p.Page = 1
	}

	if p.Limit <= 0 {
		p.Limit = 2
	}

	p.FakeCursor = strings.TrimSpace(p.FakeCursor)
}
