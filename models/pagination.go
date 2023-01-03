package models

import (
	"github.com/morkid/paginate"
	protobuffer "github.com/mxbikes/protobuf/mod"
)

type ListQuery struct {
	Size    int    `url:"size,omitempty"`
	Page    int    `url:"page,omitempty"`
	OrderBy string `url:"orderBy,omitempty"`
}

func PaginationToProto(pagination *paginate.Page) *protobuffer.Pagination {
	return &protobuffer.Pagination{
		Page:       int64(pagination.Page),
		Size:       int64(pagination.Size),
		MaxPage:    int64(pagination.MaxPage),
		TotalPages: int64(pagination.TotalPages),
		Total:      int64(pagination.Total),
		Last:       pagination.Last,
		First:      pagination.First,
		Visible:    int64(pagination.Visible),
	}
}
