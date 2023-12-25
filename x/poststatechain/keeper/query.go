package keeper

import (
	"poststate-chain/x/poststatechain/types"
)

var _ types.QueryServer = Keeper{}
