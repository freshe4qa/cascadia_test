package reward_test

import (
	"testing"

	keepertest "github.com/cascadiafoundation/cascadia/testutil/keeper"
	"github.com/cascadiafoundation/cascadia/x/reward"
	"github.com/cascadiafoundation/cascadia/x/reward/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		RewardList: []types.Reward{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.RewardKeeper(t)
	reward.InitGenesis(ctx, *k, genesisState)
	got := reward.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	require.ElementsMatch(t, genesisState.RewardList, got.RewardList)
	// this line is used by starport scaffolding # genesis/test/assert
}
