package converters

import (
	"math"
	"pillowww/titw/graph/model"
)

func PaginationToGraphql(pagination *model.PaginationInput, totalCount int) *model.Pagination {
	pageCount := float64(totalCount) / float64(pagination.Limit)
	currentPage := pagination.Offset / pagination.Limit

	roundedCount := int(math.Ceil(pageCount))

	return &model.Pagination{
		Offset:      &pagination.Offset,
		Limit:       &pagination.Limit,
		Totals:      &totalCount,
		CurrentPage: &currentPage,
		PageCount:   &roundedCount,
	}
}
