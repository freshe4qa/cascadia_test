package keeper_test

import (
	"testing"

	testkeeper "github.com/cascadiafoundation/cascadia/testutil/keeper"
	"github.com/cascadiafoundation/cascadia/x/reward/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.RewardKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
	require.EqualValues(t, params.EnableReward, k.EnableReward(ctx))
}
