package kiemtran_test

import (
	"testing"

	keepertest "kiemtran/testutil/keeper"
	"kiemtran/testutil/nullify"
	kiemtran "kiemtran/x/kiemtran/module"
	"kiemtran/x/kiemtran/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.KiemtranKeeper(t)
	kiemtran.InitGenesis(ctx, k, genesisState)
	got := kiemtran.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
