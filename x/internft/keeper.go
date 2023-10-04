package internft

import (
	ibcnfttransfertypes "github.com/bianjieai/nft-transfer/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/nft"
	nftkeeper "github.com/cosmos/cosmos-sdk/x/nft/keeper"
)

// Naive wrapper for nft keeper that implements the expected keeper interface for the nft-transfer module

type InterNftKeeperWrapper struct {
	nftKeeper *nftkeeper.Keeper
}

type InterClass struct {
	ID   string
	URI  string
	Data string
}

type InterToken struct {
	ClassID string
	ID      string
	URI     string
	Data    string
}

func (ic InterClass) GetID() string      { return ic.ID }
func (ic InterClass) GetURI() string     { return ic.URI }
func (ic InterClass) GetData() string    { return ic.Data }
func (it InterToken) GetClassID() string { return it.ClassID }
func (it InterToken) GetID() string      { return it.ID }
func (it InterToken) GetURI() string     { return it.URI }
func (it InterToken) GetData() string    { return it.Data }

func NewInterNftKeeperWrapper(nftKeeper *nftkeeper.Keeper) ibcnfttransfertypes.NFTKeeper {
	return InterNftKeeperWrapper{
		nftKeeper,
	}
}

func (ik InterNftKeeperWrapper) CreateOrUpdateClass(ctx sdk.Context,
	classID,
	classURI,
	classData string,
) error {
	class := nft.Class{
		Id:  classID,
		Uri: classURI,
	}

	if ik.nftKeeper.HasClass(ctx, classID) {
		return ik.nftKeeper.UpdateClass(ctx, class)
	}
	return ik.nftKeeper.SaveClass(ctx, class)
}

func (ik InterNftKeeperWrapper) Mint(ctx sdk.Context, classID, tokenID, tokenURI string, tokenData string, receiver sdk.AccAddress) error {
	token := nft.NFT{
		ClassId: classID,
		Id:      tokenID,
		Uri:     tokenURI,
		UriHash: "",
		Data:    nil,
	}
	return ik.nftKeeper.Mint(ctx, token, receiver)
}

func (ik InterNftKeeperWrapper) Transfer(ctx sdk.Context, classID string, tokenID string, tokenData string, receiver sdk.AccAddress) error {
	return ik.nftKeeper.Transfer(ctx, classID, tokenID, receiver)
}

func (ik InterNftKeeperWrapper) GetClass(ctx sdk.Context, classID string) (ibcnfttransfertypes.Class, bool) {
	class, exist := ik.nftKeeper.GetClass(ctx, classID)
	if !exist {
		return nil, false
	}

	return InterClass{
		ID:  class.Id,
		URI: class.Uri,
	}, true
}

func (ik InterNftKeeperWrapper) GetNFT(ctx sdk.Context, classID, tokenID string) (ibcnfttransfertypes.NFT, bool) {
	fetchedNft, has := ik.nftKeeper.GetNFT(ctx, classID, tokenID)
	if !has {
		return nil, false
	}

	return InterToken{
		ClassID: classID,
		ID:      tokenID,
		URI:     fetchedNft.Uri,
	}, true
}

// Burn implement the method of ICS721Keeper.Burn
func (ik InterNftKeeperWrapper) Burn(ctx sdk.Context, classID string, tokenID string) error {
	return ik.nftKeeper.Burn(ctx, classID, tokenID)
}

// GetOwner implement the method of ICS721Keeper.GetOwner
func (ik InterNftKeeperWrapper) GetOwner(ctx sdk.Context, classID string, tokenID string) sdk.AccAddress {
	return ik.nftKeeper.GetOwner(ctx, classID, tokenID)
}

// HasClass implement the method of ICS721Keeper.HasClass
func (ik InterNftKeeperWrapper) HasClass(ctx sdk.Context, classID string) bool {
	return ik.nftKeeper.HasClass(ctx, classID)
}