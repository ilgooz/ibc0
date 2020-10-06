package ibc0

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/test/ibc0/x/ibc0/keeper"
	"github.com/test/ibc0/x/ibc0/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	return types.DefaultGenesis()
}
