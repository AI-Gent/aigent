package aigent_test

import (
	"testing"

	keepertest "aigent/testutil/keeper"
	"aigent/testutil/nullify"
	aigent "aigent/x/aigent/module"
	"aigent/x/aigent/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.AigentKeeper(t)
	aigent.InitGenesis(ctx, k, genesisState)
	got := aigent.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
