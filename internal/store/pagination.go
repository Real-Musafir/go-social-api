package store

import (
	"net/http"
	"strconv"
)

type PaginateFeedQuery struct {
	Limit 	int			`json:"limit" validate:"gte=1,lte=20"`
	Offset 	int			`json:"offset" validate:"gte=0"`
	Sort 	string		`json:"sort" validate:"oneof=asc desc"`
}

func (fq PaginateFeedQuery) Parse(r *http.Request) (PaginateFeedQuery, error) {
	qs := r.URL.Query()
	limit := qs.Get("limit")

	if limit != "" {
		l, err := strconv.Atoi(limit)
		if err != nil {
			return fq, nil
		}
		fq.Limit = l
	}

	offset := qs.Get("offset")
	if offset != "" {
		l, err := strconv.Atoi(offset)
		if err != nil {
			return fq, nil
		}
		fq.Offset = l
	}

	sort := qs.Get("sort")
	if sort != "" {
		fq.Sort = sort
	}

	return fq, nil
}