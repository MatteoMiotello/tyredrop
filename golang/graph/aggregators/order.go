package aggregators

import (
	"pillowww/titw/graph/converters"
	"pillowww/titw/graph/model"
	"pillowww/titw/models"
)

func AggregateOrderModels(orderModels models.OrderSlice) []*model.Order {
	var ordersGraph []*model.Order

	for _, order := range orderModels {
		ordersGraph = append(ordersGraph, converters.OrderToGraphQL(order))
	}

	return ordersGraph
}

func RowsFromOrderModel(order *models.Order) []*model.OrderRow {
	rows := order.R.OrderRows
	currency := order.R.Currency

	var rowsGraph []*model.OrderRow

	for _, row := range rows {
		r, err := converters.OrderRowToGraphQL(row, currency)

		if err != nil {
			continue
		}

		rowsGraph = append(rowsGraph, r)
	}

	return rowsGraph
}
