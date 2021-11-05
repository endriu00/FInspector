package inspector

import (
	binance "github.com/binance-exchange/go-binance"
)

// `GetAssetLastPrice` returns the asset `asset` last price, calculating
// its value with the `referenceAsset`. A reference asset is an asset
// used to determine another asset price in relation to the reference
// asset value.
func (inspector *inspector) GetAssetLastPrice(asset, referenceAsset string) (float64, error) {
	ticker, err := inspector.binance.Ticker24(binance.TickerRequest{
		Symbol: asset + referenceAsset,
	})
	if err != nil {
		inspector.log.WithField("Asset", asset).WithError(err).Error("Cannot get balance for asset")
		// Return the amount calculated and the error met
		return 0, err
	}
	return ticker.LastPrice, nil
}
