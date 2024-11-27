package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "github.com/nezha90/tokenfactory/testutil/keeper"
	"github.com/nezha90/tokenfactory/x/tokenfactory/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.TokenfactoryKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
