package inspector

import (
	binance "github.com/binance-exchange/go-binance"
)

type Inspector interface {
	GetUserUSDBalanceBinance(string) (float64, error)
}

type inspector struct {
	binance binance.Binance
}
