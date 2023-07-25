package converters

import "pillowww/titw/graph/model"

func PaginationToGraphql(pagination *model.PaginationInput, totalCount int) *model.Pagination {
	pageCount := totalCount / pagination.Limit
	currentPage := pagination.Offset / pagination.Limit

	return &model.Pagination{
		Offset:      &pagination.Offset,
		Limit:       &pagination.Limit,
		Totals:      &totalCount,
		CurrentPage: &currentPage,
		PageCount:   &pageCount,
	}
}
