package tokenfactory_test

import (
	"testing"

	keepertest "github.com/nezha90/tokenfactory/testutil/keeper"
	"github.com/nezha90/tokenfactory/testutil/nullify"
	tokenfactory "github.com/nezha90/tokenfactory/x/tokenfactory/module"
	"github.com/nezha90/tokenfactory/x/tokenfactory/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		DenomList: []types.Denom{
			{
				Denom: "0",
			},
			{
				Denom: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.TokenfactoryKeeper(t)
	tokenfactory.InitGenesis(ctx, k, genesisState)
	got := tokenfactory.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.DenomList, got.DenomList)
	// this line is used by starport scaffolding # genesis/test/assert
}
