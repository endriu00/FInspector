package inspector

import (
	"errors"
	binance "github.com/binance-exchange/go-binance"
	"github.com/sirupsen/logrus"
)

// `Inspector` is the interface for the inspector.
type Inspector interface {
	GetUserTotalBalanceBinance(string, string) (float64, error)
	GetUserAssetBalanceBinance(string, string, string) (float64, error)
}

// `Config` is the configuration struct for the inspector.
type Config struct {
	Binance binance.Binance
	Log     *logrus.Entry
}

// `inspector` collects the balance of a user wallet.
type inspector struct {
	// binance is the binance interface for accessing the Binance API.
	binance binance.Binance

	// log is the inspector log.
	log *logrus.Entry
}

// `New` creates a new Inspector starting from the configurations `cfg`.
// If some mandatory parameters in `cfg` are missing, the function will
// return an error.
func New(cfg Config) (Inspector, error) {
	if cfg.Binance == nil {
		return nil, errors.New("binance not initialized")
	}
	if cfg.Log == nil {
		return nil, errors.New("logger not initialized")
	}
	return &inspector{
		binance: cfg.Binance,
		log:     cfg.Log,
	}, nil
}
