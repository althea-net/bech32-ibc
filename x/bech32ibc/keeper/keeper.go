package keeper

import (
	"fmt"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/althea-net/bech32-ibc/x/bech32ibc/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdkstore "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type (
	Keeper struct {
		channelKeeper types.ChannelKeeper

		cdc      codec.Codec
		storeKey sdkstore.StoreKey

		transferKeeper    types.TransferKeeper
		nftTransferKeeper types.NftTransferKeeper
	}
)

func NewKeeper(
	channelKeeper types.ChannelKeeper,
	cdc codec.Codec,
	storeKey sdkstore.StoreKey,
	transferKeeper types.TransferKeeper,
	nftTransferKeeper types.NftTransferKeeper,
) *Keeper {
	return &Keeper{
		channelKeeper:     channelKeeper,
		cdc:               cdc,
		storeKey:          storeKey,
		transferKeeper:    transferKeeper,
		nftTransferKeeper: nftTransferKeeper,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}
