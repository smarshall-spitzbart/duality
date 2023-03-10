package keeper

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/NicholasDotSol/duality/x/dex/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const LPsharesRegexpStr = "^" + types.DepositSharesPrefix + "-" +
	// Token0 (regexp from cosmos-sdk.types.coin.reDnmString)
	"([a-zA-Z][a-zA-Z0-9/-]{2,127})" + "-" +
	// Token1
	"([a-zA-Z][a-zA-Z0-9/-]{2,127})" + "-" +
	// Tickindex
	`t(\d+)` + "-" +
	// feeIndex
	`f(\d+)`

var LPSharesRegexp = regexp.MustCompile(LPsharesRegexpStr)

func CreateSharesId(token0 string, token1 string, tickIndex int64, feeIndex uint64) (denom string) {
	t0 := strings.ReplaceAll(token0, "-", "")
	t1 := strings.ReplaceAll(token1, "-", "")
	return fmt.Sprintf("%s-%s-%s-t%d-f%d", types.DepositSharesPrefix, t0, t1, tickIndex, feeIndex)
}

func ParseDepositShares(shares sdk.Coin) (matches []string, valid bool) {
	// NOTE: Since dashes are removed as part of CreateSharesId, if either side of the LP position are denoms that contain dashes
	// they will not be parsed correctly and the correct dneom will not be returned
	matchArr := LPSharesRegexp.FindAllStringSubmatch(shares.Denom, -1)
	if matchArr == nil {
		return nil, false

	}
	return matchArr[0][1:5], true
}

func DepositSharesToData(shares sdk.Coin, feeTiers []types.FeeTier) (types.DepositRecord, error) {
	matches, valid := ParseDepositShares(shares)

	if !valid {
		return types.DepositRecord{}, types.ErrInvalidDepositShares
	}

	pairId := CreatePairId(matches[0], matches[1])
	tickIndex, err := strconv.ParseInt(matches[2], 10, 0)
	if err != nil {
		return types.DepositRecord{}, types.ErrInvalidDepositShares
	}
	feeIndex, err := strconv.ParseUint(matches[3], 10, 0)
	if err != nil {
		return types.DepositRecord{}, types.ErrInvalidDepositShares
	}

	return types.DepositRecord{
		PairId:          pairId,
		SharesOwned:     shares.Amount,
		CenterTickIndex: tickIndex,
		LowerTickIndex:  tickIndex - feeTiers[feeIndex].Fee,
		UpperTickIndex:  tickIndex + feeTiers[feeIndex].Fee,
		FeeIndex:        feeIndex,
	}, nil
}
