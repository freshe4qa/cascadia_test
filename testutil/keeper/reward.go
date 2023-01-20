package keeper

import (
	"testing"

	"github.com/cascadiafoundation/cascadia/v1/x/reward/keeper"
	"github.com/cascadiafoundation/cascadia/v1/x/reward/types"
	"github.com/cosmos/cosmos-sdk/store"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	tmdb "github.com/tendermint/tm-db"
)

func RewardKeeper(t testing.TB) (*keeper.Keeper, sdk.Context) {
	storeKey := sdk.NewKVStoreKey(types.StoreKey)
	memStoreKey := storetypes.NewMemoryStoreKey(types.MemStoreKey)

	db := tmdb.NewMemDB()
	stateStore := store.NewCommitMultiStore(db)
	stateStore.MountStoreWithDB(storeKey, sdk.StoreTypeIAVL, db)
	stateStore.MountStoreWithDB(memStoreKey, sdk.StoreTypeMemory, nil)
	require.NoError(t, stateStore.LoadLatestVersion())

	// registry := codectypes.NewInterfaceRegistry()
	// cdc := codec.NewProtoCodec(registry)

	// paramsSubspace := typesparams.NewSubspace(cdc,
	// 	types.Amino,
	// 	storeKey,
	// 	memStoreKey,
	// 	"RewardParams",
	// )
	// k := keeper.NewKeeper(
	// 	cdc,
	// 	storeKey,
	// 	memStoreKey,
	// 	paramsSubspace,
	// 	nil,
	// )

	ctx := sdk.NewContext(stateStore, tmproto.Header{}, false, log.NewNopLogger())

	// Initialize params
	// k.SetParams(ctx, types.DefaultParams())

	return nil, ctx
}
