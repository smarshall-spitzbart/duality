package keeper

import (
	"context"

	"github.com/NicholasDotSol/duality/x/dex/types"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type TickIterator struct {
	iter sdk.Iterator
	cdc  codec.BinaryCodec
}

func (k Keeper) NewTickIterator(
	// NOTE: both start and end are inclusive
	ctx context.Context,
	startInclusive int64,
	endInclusive int64,
	pairId *types.PairId,
	scanLeft bool,
) types.TickIteratorI {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	prefixStore := prefix.NewStore(sdkCtx.KVStore(k.storeKey), types.TickPrefix(pairId))

	if scanLeft {
		return TickIterator{
			iter: prefixStore.ReverseIterator(
				types.TickIndexToBytes(endInclusive),
				types.TickIndexToBytes(startInclusive+1),
			),
			cdc: k.cdc,
		}
	} else {
		return TickIterator{
			iter: prefixStore.Iterator(
				types.TickIndexToBytes(startInclusive),
				types.TickIndexToBytes(endInclusive+1),
			),
			cdc: k.cdc,
		}
	}
}

func (ti TickIterator) Valid() bool {
	return ti.iter.Valid()
}

func (ti TickIterator) Close() error {
	return ti.iter.Close()
}

func (ti TickIterator) Value() types.Tick {
	var tick types.Tick
	ti.cdc.MustUnmarshal(ti.iter.Value(), &tick)
	return tick
}

func (ti TickIterator) Next() {
	ti.iter.Next()
}
