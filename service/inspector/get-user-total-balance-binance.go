package inspector

// `GetUserTotalBalanceBinance` returns the total balance for the user `user`
// having `asset` as the reference asset. The reference asset is the asset
// used for determining the price of another asset in its value.
// For example, in BTC/USDT the reference asset is USDT, so the price of
// BTC will be calculated in relation to the price of USDT.
// The recomended reference asset is USDT, because almost every other asset
// is listed in the exchanges with it.
func (inspector *inspector) GetUserTotalBalanceBinance(user, asset string) (float64, error) {
	var userAmount float64
	var lastPrice float64

	// Get user information for his Binance account
	account, err := inspector.GetUserInfoBinance()
	if err != nil {
		inspector.log.WithField("Provider", "binance").WithError(err).Error("Could not get user info.")
		return -1, err
	}

	// Sum each asset of his wallet to the total amount
	for _, balance := range account.Balances {
		// Do not consider 0 value balances
		if balance.Free == 0 && balance.Locked == 0 {
			continue
		}

		switch balance.Asset {
		// LD is not recognized as a crypto
		case "LDBNB":
			fallthrough
		case "LDUSDT":
			continue
		// BETH has no market value in USDT because it is tightly coupled with ETH.
		// BETH is a binance wrapped version of ETH for ETH 2.0 staking.
		// Its value SHOULD always be pegged to ETH, so, for obtaining BETH
		// last price, ETH/USDT pair is used instead.
		case "BETH":
			lastPrice, err = inspector.GetAssetLastPrice("ETH", asset)
			if err != nil {
				inspector.log.WithError(err).Error("Could not get asset last price")
				return userAmount, err
			}
		default:
			// Get asset last price in asset/USDT value
			lastPrice, err = inspector.GetAssetLastPrice(balance.Asset, asset)
			if err != nil {
				inspector.log.WithError(err).Error("Could not get asset last price")
				return userAmount, err
			}
		}

		// Sum free and locked assets quantity
		totalAsset := balance.Free + balance.Locked

		// Calculate asset amount in USDT and sum it to userAmount
		userAmount += lastPrice * totalAsset
	}
	return userAmount, nil
}
