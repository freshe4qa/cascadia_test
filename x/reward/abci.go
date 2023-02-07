package reward

import (
	"github.com/cascadiafoundation/cascadia/x/reward/keeper"
	"github.com/cascadiafoundation/cascadia/x/reward/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
)

// EndBlocker called every block, process inflation, update validator set.
func EndBlocker(ctx sdk.Context, keeper keeper.Keeper) {
	blockHeight := ctx.BlockHeight()

	var contractAddress common.Address
	var err error
	// Deploy Voting Escrow contract to the chain
	if blockHeight == 1 {
		contractAddress, err = keeper.DeployVotingEscrowContract(ctx)

		if err != nil {
			panic(err)
		}

		gasFeeShares, _ := sdk.NewDecFromStr("0.33")
		blockRewardShares, _ := sdk.NewDecFromStr("0.33")

		keeper.SetReward(ctx, types.Reward{
			Index:             "vecontract",
			Contract:          contractAddress.String(),
			GasFeeShares:      gasFeeShares,
			BlockRewardShares: blockRewardShares,
		})
	}
}
