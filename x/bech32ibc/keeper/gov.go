package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/althea-net/bech32-ibc/x/bech32ibc/types"
)

func (k Keeper) HandleUpdateHrpIbcChannelProposal(ctx sdk.Context, p *types.UpdateHrpIbcChannelProposal) error {
	err := types.ValidateHrp(p.Hrp)
	if err != nil {
		return err
	}

	// Check that the fungible channel exists unless it's empty (in which case it will be "unset" in the record)
	if p.FungibleSourceChannel != "" {
		_, fungibleChannelFound := k.channelKeeper.GetChannel(ctx, k.transferKeeper.GetPort(ctx), p.FungibleSourceChannel)
		if !fungibleChannelFound {
			return sdkerrors.Wrap(types.ErrInvalidIBCData, fmt.Sprintf("fungible channel not found: %s", p.FungibleSourceChannel))
		}
	}

	// Check that the nft channel exists unless it's empty (in which case it will be "unset" in the record)
	if p.NftSourceChannel != "" {
		_, nftChannelFound := k.channelKeeper.GetChannel(ctx, k.nftTransferKeeper.GetPort(ctx), p.NftSourceChannel)
		if !nftChannelFound {
			return sdkerrors.Wrap(types.ErrInvalidIBCData, fmt.Sprintf("nft channel not found: %s", p.NftSourceChannel))
		}
	}

	return k.setHrpIbcRecord(ctx, types.HrpIbcRecord{
		Hrp:                   p.Hrp,
		FungibleSourceChannel: p.FungibleSourceChannel,
		NftSourceChannel:      p.NftSourceChannel,
	})
}
