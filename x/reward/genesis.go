package reward

import (
	"github.com/cascadiafoundation/cascadia/x/reward/keeper"
	"github.com/cascadiafoundation/cascadia/x/reward/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the reward
	for _, elem := range genState.RewardList {
		k.SetReward(ctx, elem)
	}
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.RewardList = k.GetAllReward(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
