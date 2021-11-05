package inspector

import (
	binance "github.com/binance-exchange/go-binance"
	"time"
)

// `GetUserInfoBinance` returns the user binance account.
func (inspector *inspector) GetUserInfoBinance() (*binance.Account, error) {
	var userAccount *binance.Account

	rcvWindow, err := time.ParseDuration("3000ms")
	if err != nil {
		inspector.log.WithError(err).Error("Cannot parse time duration")
		return userAccount, err
	}
	userAccount, err = inspector.binance.Account(binance.AccountRequest{
		RecvWindow: rcvWindow,
		Timestamp:  time.Now(),
	})
	if err != nil {
		inspector.log.WithError(err).Error("Cannot request account from Binance")
		return userAccount, err
	}
	return userAccount, nil
}
