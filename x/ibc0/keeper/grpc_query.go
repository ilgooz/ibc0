package keeper

import (
	"github.com/test/ibc0/x/ibc0/types"
)

var _ types.QueryServer = Keeper{}
