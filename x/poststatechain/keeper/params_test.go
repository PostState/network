package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "poststate-chain/testutil/keeper"
	"poststate-chain/x/poststatechain/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.PoststatechainKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
