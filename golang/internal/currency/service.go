package currency

import (
	currency2 "github.com/bojanz/currency"
	"strconv"
)

func ToFloat(price int, isoCode string) (float64, error) {
	amount, err := currency2.NewAmountFromInt64(int64(price), isoCode)
	if err != nil {
		return 0, err
	}

	floatValue, err := strconv.ParseFloat(amount.Number(), 64)

	if err != nil {
		return 0, err
	}

	return floatValue, nil
}
