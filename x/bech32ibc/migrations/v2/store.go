package v2

import (
	v1 "github.com/althea-net/bech32-ibc/x/bech32ibc/migrations/v1"
	"github.com/althea-net/bech32-ibc/x/bech32ibc/types"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// MigrateStore performs in-place store migrations from v1 to v2. The migration includes:
//
// - Migrate all HrpIbcRecords to the new format which includes an NFT channel for each hrp
func MigrateStore(ctx sdk.Context, storeKey storetypes.StoreKey, cdc codec.BinaryCodec) error {
	ctx.Logger().Info("bech32-ibc module in-store migration from v1 to v2 started")

	store := ctx.KVStore(storeKey)
	migrateHrpIbcRecord(store, cdc)

	ctx.Logger().Info("bech32-ibc module in-store migration completed successfully")
	return nil
}

// migrateHrpIbcRecord migrates HrpIbcRecord from v1 to v2
// It iterates through all the records and updates the record to the new format
func migrateHrpIbcRecord(store sdk.KVStore, cdc codec.BinaryCodec) error {
	prefixStore := prefix.NewStore(store, types.HrpIBCRecordStorePrefix)
	iterator := prefixStore.Iterator(nil, nil)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var oldRecord v1.HrpIbcRecord
		if err := cdc.Unmarshal(iterator.Value(), &oldRecord); err != nil {
			return err
		}

		newRecord := types.HrpIbcRecord{
			Hrp:                   oldRecord.Hrp,
			FungibleSourceChannel: oldRecord.SourceChannel,
		}

		b, err := cdc.Marshal(&newRecord)
		if err != nil {
			return err
		}
		prefixStore.Set(iterator.Key(), b)
	}

	return nil
}
