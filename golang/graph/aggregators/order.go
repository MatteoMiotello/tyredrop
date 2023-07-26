package aggregators

import (
	"pillowww/titw/graph/converters"
	"pillowww/titw/graph/model"
	"pillowww/titw/models"
)

func AggregateOrderModels(orderModels models.OrderSlice) ([]*model.Order, error) {
	var ordersGraph []*model.Order

	for _, order := range orderModels {
		order, err := converters.OrderToGraphQL(order)

		if err != nil {
			return nil, err
		}

		ordersGraph = append(ordersGraph, order)
	}

	return ordersGraph, nil
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
