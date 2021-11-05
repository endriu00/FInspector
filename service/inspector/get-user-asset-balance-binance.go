package inspector

import (
	"github.com/sirupsen/logrus"
)

// `GetUserAssetBalanceBinance` returns the specific `asset` balance for
// the user `user` having `asset` as the reference asset.
// The reference asset is the asset used for determining the price of another
// asset in its value. For example, in BTC/USDT the reference asset is USDT,
// so the price of BTC will be calculated in relation to the price of USDT.
// The recomended reference asset is USDT, because almost every other asset
// is listed in the exchanges with it.
func (inspector *inspector) GetUserAssetBalanceBinance(user, asset, referenceAsset string) (float64, error) {
	var userBalance float64
	var totalAsset float64

	// Get user information for his Binance account
	account, err := inspector.GetUserInfoBinance()
	if err != nil {
		inspector.log.WithField("Provider", "binance").WithError(err).Error("Could not get user info.")
		return -1, err
	}

	// Take the latest price of the desired asset in relation to the referenceAsset
	for _, balance := range account.Balances {
		if balance.Asset == asset {
			lastPrice, err := inspector.GetAssetLastPrice(asset, referenceAsset)
			if err != nil {
				inspector.log.WithFields(logrus.Fields{
					"asset":           asset,
					"reference asset": referenceAsset,
				}).WithError(err).Error("Could not get asset last price")
				return -1, err
			}

			// Sum up the locked and free amount of the asset to get
			// the user total amount of the asset
			totalAsset = balance.Free + balance.Locked

			// userBalance of the asset asset is asset lastPrice dot user totalAsset
			userBalance = lastPrice * totalAsset
			break
		}
	}
	return userBalance, nil
}
