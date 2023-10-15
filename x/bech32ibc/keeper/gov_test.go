package keeper_test

import (
	"github.com/althea-net/bech32-ibc/app"
	"github.com/althea-net/bech32-ibc/x/bech32ibc/keeper"
	"github.com/althea-net/bech32-ibc/x/bech32ibc/testutil"
	"github.com/althea-net/bech32-ibc/x/bech32ibc/types"
	sdktestutil "github.com/cosmos/cosmos-sdk/testutil"
	sdk "github.com/cosmos/cosmos-sdk/types"
	channeltypes "github.com/cosmos/ibc-go/v5/modules/core/04-channel/types"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"testing"
)

func TestHandleUpdateHrpIbcChannelProposal(t *testing.T) {
	tests := map[string]struct {
		proposal *types.UpdateHrpIbcChannelProposal
		expErr   string
		// setting manualMockSetup to nil will use the happy path default mock setup
		manualMockSetup func(*testutil.MockChannelKeeper, *testutil.MockTransferKeeper, *testutil.MockNftTransferKeeper)
	}{
		"happy path": {
			proposal: &types.UpdateHrpIbcChannelProposal{
				Title:                 "Test title",
				Description:           "Test description",
				Hrp:                   "stars",
				FungibleSourceChannel: "channel-0",
				NftSourceChannel:      "channel-1",
			},
		},
		"happy path with empty nft channel": {
			proposal: &types.UpdateHrpIbcChannelProposal{
				Title:                 "Test title",
				Description:           "Test description",
				Hrp:                   "stars",
				FungibleSourceChannel: "channel-0",
				NftSourceChannel:      "",
			},
		},
		"happy path with empty fungible channel": {
			proposal: &types.UpdateHrpIbcChannelProposal{
				Title:                 "Test title",
				Description:           "Test description",
				Hrp:                   "stars",
				FungibleSourceChannel: "",
				NftSourceChannel:      "channel-42",
			},
		},
		"happy path with empty channels (should delete the record)": {
			proposal: &types.UpdateHrpIbcChannelProposal{
				Title:                 "Test title",
				Description:           "Test description",
				Hrp:                   "stars",
				FungibleSourceChannel: "",
				NftSourceChannel:      "",
			},
		},
		"empty hrp": {
			proposal: &types.UpdateHrpIbcChannelProposal{
				Title:                 "Test title",
				Description:           "Test description",
				Hrp:                   "",
				FungibleSourceChannel: "channel-0",
				NftSourceChannel:      "channel-1",
			},
			expErr: "empty HRP",
			manualMockSetup: func(channelKeeper *testutil.MockChannelKeeper, transferKeeper *testutil.MockTransferKeeper, nftTransferKeeper *testutil.MockNftTransferKeeper) {
				// No mocks because we expect it to not get very far
			},
		},
		"invalid (not found) fungible channel": {
			proposal: &types.UpdateHrpIbcChannelProposal{
				Title:                 "Test title",
				Description:           "Test description",
				Hrp:                   "stars",
				FungibleSourceChannel: "channel-1337",
				NftSourceChannel:      "channel-42",
			},
			expErr: "fungible channel not found: channel-1337",
			manualMockSetup: func(channelKeeper *testutil.MockChannelKeeper, transferKeeper *testutil.MockTransferKeeper, nftTransferKeeper *testutil.MockNftTransferKeeper) {
				transferKeeper.EXPECT().GetPort(gomock.Any()).Return("transfer")
				channelKeeper.EXPECT().GetChannel(gomock.Any(), gomock.Any(), "channel-1337").Return(channeltypes.Channel{}, false)
			},
		},
		"invalid (not found) nft channel": {
			proposal: &types.UpdateHrpIbcChannelProposal{
				Title:                 "Test title",
				Description:           "Test description",
				Hrp:                   "stars",
				FungibleSourceChannel: "channel-420",
				NftSourceChannel:      "channel-9001",
			},
			expErr: "nft channel not found: channel-9001",
			manualMockSetup: func(channelKeeper *testutil.MockChannelKeeper, transferKeeper *testutil.MockTransferKeeper, nftTransferKeeper *testutil.MockNftTransferKeeper) {
				transferKeeper.EXPECT().GetPort(gomock.Any()).Return("transfer")
				channelKeeper.EXPECT().GetChannel(gomock.Any(), gomock.Any(), "channel-420").Return(channeltypes.Channel{}, true)

				nftTransferKeeper.EXPECT().GetPort(gomock.Any()).Return("nfttransfer")
				channelKeeper.EXPECT().GetChannel(gomock.Any(), gomock.Any(), "channel-9001").Return(channeltypes.Channel{}, false)
			},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			encCfg := app.MakeEncodingConfig()
			key := sdk.NewKVStoreKey(types.StoreKey)

			mockChannelKeeper := testutil.NewMockChannelKeeper(ctrl)
			mockTransferKeeper := testutil.NewMockTransferKeeper(ctrl)
			mockNftTransferKeeper := testutil.NewMockNftTransferKeeper(ctrl)

			if tc.manualMockSetup != nil {
				tc.manualMockSetup(mockChannelKeeper, mockTransferKeeper, mockNftTransferKeeper)
			} else {
				if tc.proposal.FungibleSourceChannel != "" {
					mockChannelKeeper.EXPECT().GetChannel(gomock.Any(), gomock.Any(), tc.proposal.FungibleSourceChannel).Return(channeltypes.Channel{}, true)
					mockTransferKeeper.EXPECT().GetPort(gomock.Any()).Return("transfer")
				}

				if tc.proposal.NftSourceChannel != "" {
					mockChannelKeeper.EXPECT().GetChannel(gomock.Any(), gomock.Any(), tc.proposal.NftSourceChannel).Return(channeltypes.Channel{}, true)
					mockNftTransferKeeper.EXPECT().GetPort(gomock.Any()).Return("nfttransfer")
				}
			}

			testKeeper := keeper.NewKeeper(mockChannelKeeper, encCfg.Codec, key, mockTransferKeeper, mockNftTransferKeeper)

			ctx := sdktestutil.DefaultContext(key, sdk.NewTransientStoreKey("transient_test"))

			err := testKeeper.HandleUpdateHrpIbcChannelProposal(ctx, tc.proposal)
			if tc.expErr != "" {
				require.ErrorContains(t, err, tc.expErr)
			} else {
				require.NoError(t, err)

				records := testKeeper.GetHrpIbcRecords(ctx)
				if tc.proposal.FungibleSourceChannel == "" && tc.proposal.NftSourceChannel == "" {
					require.Len(t, records, 0)
				} else {
					require.Len(t, records, 1)
					require.Equal(t, tc.proposal.Hrp, records[0].Hrp)
					require.Equal(t, tc.proposal.FungibleSourceChannel, records[0].FungibleSourceChannel)
					require.Equal(t, tc.proposal.NftSourceChannel, records[0].NftSourceChannel)
				}
			}
		})
	}
}
