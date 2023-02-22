package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/cascadiafoundation/cascadia/testutil/keeper"
	"github.com/cascadiafoundation/cascadia/x/reward/keeper"
	"github.com/cascadiafoundation/cascadia/x/reward/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.RewardKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
