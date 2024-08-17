package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "kiemtran/testutil/keeper"
	"kiemtran/x/kiemtran/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.KiemtranKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
