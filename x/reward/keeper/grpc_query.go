package keeper

import (
	"github.com/cascadiafoundation/cascadia/v1/x/reward/types"
)

var _ types.QueryServer = Keeper{}
