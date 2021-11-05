package main

import (
	"context"
	binance "github.com/binance-exchange/go-binance"
	"github.com/endriu00/FInspector/service/inspector"
	"github.com/go-kit/log"
	"github.com/sirupsen/logrus"
	"os"
)

func main() {
	// Start loggers
	binanceLog := log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	binanceLog = log.With(binanceLog, "time", log.DefaultTimestampUTC, "caller", log.DefaultCaller)
	log := logrus.NewEntry(logrus.StandardLogger())

	// Get configuration
	config, err := getConfigurationEnv()
	if err != nil {
		log.WithError(err).Error("Could not load configuration from environment")
		return
	}

	// initialize the signer for the binance API secret key
	// and the API service.
	hmacSigner := &binance.HmacSigner{
		Key: []byte(config.Binance.SecretKey),
	}
	binanceService := binance.NewAPIService(
		config.Binance.ApiURL,
		config.Binance.ApiKey,
		hmacSigner,
		binanceLog,
		context.Background(),
	)
	binance := binance.NewBinance(binanceService)

	// Initialize the inspector
	inspector, err := inspector.New(inspector.Config{
		Binance: binance,
		Log:     log,
	})
	if err != nil {
		binanceLog.Log("error", err)
		return
	}

	// START TEST
	totalBalance, err := inspector.GetUserTotalBalanceBinance("", "USDT")
	if err != nil {
		return
	}
	log.Info("You have USDT")
	log.Info(totalBalance)

	btcBalance, err := inspector.GetUserAssetBalanceBinance("", "BTC", "USDT")
	if err != nil {
		return
	}
	log.Info("You have BTC")
	log.Info(btcBalance)
	// END TEST

}
