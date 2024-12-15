package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "aigent/testutil/keeper"
	"aigent/x/aigent/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.AigentKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
