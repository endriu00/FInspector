package inspector

import (
	"fmt"
	binance "github.com/binance-exchange/go-binance"
	"time"
)

func (inspector *inspector) GetUserUSDBalanceBinance(user string) (float64, error) {
	var userAmount float64
	account, err := inspector.binance.Account(binance.AccountRequest{
		RecvWindow: 0,
		Timestamp:  time.Now(),
	})
	if err != nil {
		return -1, err
	}
	for _, balance := range account.Balances {
		// Get asset last price in asset/USDT value
		ticker, err := inspector.binance.Ticker24(binance.TickerRequest{
			Symbol: balance.Asset + "USDT",
		})
		if err != nil {
			return -1, err
		}
		userAmount += ticker.LastPrice
	}

	fmt.Println(userAmount)
	return userAmount, nil
}
