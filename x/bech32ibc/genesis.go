package bech32ibc

import (
	"github.com/althea-net/bech32-ibc/x/bech32ibc/keeper"
	"github.com/althea-net/bech32-ibc/x/bech32ibc/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	k.SetNativeHrp(ctx, genState.NativeHRP)
	k.SetHrpIbcRecords(ctx, genState.HrpIBCRecords)
}

// ExportGenesis returns the module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	hrpIbcRecords := k.GetHrpIbcRecords(ctx)
	nativeHrp, err := k.GetNativeHrp(ctx)
	if err != nil {
		panic(err)
	}

	return &types.GenesisState{
		NativeHRP:     nativeHrp,
		HrpIBCRecords: hrpIbcRecords,
	}
}
