package keeper

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/althea-net/bech32-ibc/x/bech32ibc/types"
)

var _ types.QueryServer = Keeper{}

func (k Keeper) HrpIbcRecords(ctx context.Context, _ *types.QueryHrpIbcRecordsRequest) (*types.QueryHrpIbcRecordsResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)

	records := k.GetHrpIbcRecords(sdkCtx)

	return &types.QueryHrpIbcRecordsResponse{HrpIbcRecords: records}, nil
}

func (k Keeper) HrpIbcRecord(ctx context.Context, req *types.QueryHrpIbcRecordRequest) (*types.QueryHrpIbcRecordResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	sdkCtx := sdk.UnwrapSDKContext(ctx)

	record, err := k.GetHrpIbcRecord(sdkCtx, req.GetHrp())

	if err != nil {
		return nil, err
	}

	return &types.QueryHrpIbcRecordResponse{HrpIbcRecord: record}, nil
}

func (k Keeper) NativeHrp(ctx context.Context, _ *types.QueryNativeHrpRequest) (*types.QueryNativeHrpResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)

	hrp, err := k.GetNativeHrp(sdkCtx)
	if err != nil {
		return nil, err
	}

	return &types.QueryNativeHrpResponse{NativeHrp: hrp}, nil
}
