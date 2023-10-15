package v2_test

import (
	"github.com/althea-net/bech32-ibc/app/params"
	v1 "github.com/althea-net/bech32-ibc/x/bech32ibc/migrations/v1"
	v2 "github.com/althea-net/bech32-ibc/x/bech32ibc/migrations/v2"
	"github.com/althea-net/bech32-ibc/x/bech32ibc/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	"github.com/cosmos/cosmos-sdk/testutil"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"testing"
)

var testData = []v1.HrpIbcRecord{
	{
		Hrp:               "stars",
		SourceChannel:     "channel-0",
		IcsToHeightOffset: 0,
		IcsToTimeOffset:   0,
	},
	{
		Hrp:               "cosmos",
		SourceChannel:     "channel-1",
		IcsToHeightOffset: 10,
		IcsToTimeOffset:   10,
	},
	{
		Hrp:               "osmo",
		SourceChannel:     "channel-42",
		IcsToHeightOffset: 420,
		IcsToTimeOffset:   1337,
	},
}

func TestMigrateStore(t *testing.T) {
	cdc := params.MakeEncodingConfig().Codec
	storeKey := sdk.NewKVStoreKey(types.StoreKey)
	ctx := testutil.DefaultContext(storeKey, sdk.NewTransientStoreKey("transient_test"))
	store := prefix.NewStore(ctx.KVStore(storeKey), types.HrpIBCRecordStorePrefix)

	for _, record := range testData {
		bz, err := cdc.Marshal(&record)
		require.NoError(t, err)
		store.Set([]byte(record.Hrp), bz)
	}

	err := v2.MigrateStore(ctx, storeKey, cdc)
	require.NoError(t, err)

	iterator := store.Iterator(nil, nil)
	defer iterator.Close()
	count := 0
	for ; iterator.Valid(); iterator.Next() {
		count++
	}
	require.Equal(t, len(testData), count)

	for _, record := range testData {
		var newRecord types.HrpIbcRecord
		bz := store.Get([]byte(record.Hrp))
		require.NoError(t, cdc.Unmarshal(bz, &newRecord))
		require.Equal(t, record.Hrp, newRecord.Hrp)
		require.Equal(t, record.SourceChannel, newRecord.FungibleSourceChannel)
		require.Equal(t, "", newRecord.NftSourceChannel)
	}
}
