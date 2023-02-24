package main

import (
	"fmt"
	"github.com/bojanz/currency"
	"pillowww/titw/internal/bootstrap"
)

func init() {
	bootstrap.InitConfig()
	bootstrap.InitDb()
}

func main() {
	amount, _ := currency.NewAmount("345.333333", "USD")
	fmt.Println(*amount.BigInt())
}
