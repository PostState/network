package poststatechain_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "poststate-chain/testutil/keeper"
	"poststate-chain/testutil/nullify"
	"poststate-chain/x/poststatechain/module"
	"poststate-chain/x/poststatechain/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.PoststatechainKeeper(t)
	poststatechain.InitGenesis(ctx, k, genesisState)
	got := poststatechain.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
