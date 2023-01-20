package simulation

import (
	"math/rand"

	"github.com/cascadiafoundation/cascadia/v1/x/reward/keeper"
	"github.com/cascadiafoundation/cascadia/v1/x/reward/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func SimulateMsgRegisterVeContractReward(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgRegisterVeContractReward{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the RegisterReward simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "RegisterReward simulation not implemented"), nil, nil
	}
}
