package keeper_test

import (
	"testing"

	dualityapp "github.com/NicholasDotSol/duality/app"
	"github.com/NicholasDotSol/duality/x/dex/keeper"
	"github.com/NicholasDotSol/duality/x/dex/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/suite"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
)

type TickTestSuite struct {
	suite.Suite
	app *dualityapp.App
	ctx sdk.Context
}

func TestTickTestSuite(t *testing.T) {
	suite.Run(t, new(TickTestSuite))
}

func (s *TickTestSuite) SetupTest() {
	s.app = dualityapp.Setup(false)
	s.ctx = s.app.BaseApp.NewContext(false, tmproto.Header{})
}

func (s *TickTestSuite) TestLimitHasToken0() {
	pairId := "TokenA/TokenB"
	tick := keeper.NewTick("TokenA/TokenB", 0, 6)
	s.app.DexKeeper.SetLimitOrderPoolReserveMap(s.ctx, types.LimitOrderPoolReserveMap{
		PairId:    pairId,
		Token:     "TokenA",
		TickIndex: tick.TickIndex,
		Count:     0,
		Reserves:  newDec("5"),
	})
	s.Equal(true, s.app.DexKeeper.HasToken0(s.ctx, &tick))
}

func (s *TickTestSuite) TestEmptyHasToken0() {
	pairId := "TokenA/TokenB"
	tick := keeper.NewTick("TokenA/TokenB", 0, 6)
	s.app.DexKeeper.SetLimitOrderPoolReserveMap(s.ctx, types.LimitOrderPoolReserveMap{
		PairId:    pairId,
		Token:     "TokenA",
		TickIndex: tick.TickIndex,
		Count:     0,
		Reserves:  newDec("0"),
	})
	s.Equal(false, s.app.DexKeeper.HasToken0(s.ctx, &tick))
}

func (s *TickTestSuite) TestLiquidityHasToken0() {
	pairId := "TokenA/TokenB"
	tick := keeper.NewTick("TokenA/TokenB", 0, 6)
	tick.TickData.Reserve0AndShares[0].Reserve0 = newDec("100")
	tick.TickData.Reserve0AndShares[0].TotalShares = newDec("100")
	tick.TickData.Reserve1[0] = newDec("20")
	s.app.DexKeeper.SetLimitOrderPoolReserveMap(s.ctx, types.LimitOrderPoolReserveMap{
		PairId:    pairId,
		Token:     "TokenA",
		TickIndex: tick.TickIndex,
		Count:     0,
		Reserves:  newDec("0"),
	})
	s.Equal(true, s.app.DexKeeper.HasToken0(s.ctx, &tick))
}

func (s *TickTestSuite) TestLiquidityHighFeeHasToken0() {
	pairId := "TokenA/TokenB"
	tick := keeper.NewTick("TokenA/TokenB", 0, 6)
	tick.TickData.Reserve0AndShares[4].Reserve0 = newDec("100")
	tick.TickData.Reserve0AndShares[4].TotalShares = newDec("100")
	tick.TickData.Reserve1[4] = newDec("20")
	s.app.DexKeeper.SetLimitOrderPoolReserveMap(s.ctx, types.LimitOrderPoolReserveMap{
		PairId:    pairId,
		Token:     "TokenA",
		TickIndex: tick.TickIndex,
		Count:     0,
		Reserves:  newDec("0"),
	})
	s.Equal(true, s.app.DexKeeper.HasToken0(s.ctx, &tick))
}

func (s *TickTestSuite) TestNoLiquidityOnOneSideHasToken0() {
	pairId := "TokenA/TokenB"
	tick := keeper.NewTick("TokenA/TokenB", 0, 6)
	tick.TickData.Reserve0AndShares[4].Reserve0 = newDec("0")
	tick.TickData.Reserve0AndShares[4].TotalShares = newDec("10")
	tick.TickData.Reserve1[4] = newDec("20")
	s.app.DexKeeper.SetLimitOrderPoolReserveMap(s.ctx, types.LimitOrderPoolReserveMap{
		PairId:    pairId,
		Token:     "TokenA",
		TickIndex: tick.TickIndex,
		Count:     0,
		Reserves:  newDec("0"),
	})
	s.Equal(false, s.app.DexKeeper.HasToken0(s.ctx, &tick))
}

func (s *TickTestSuite) TestLimitHasToken1() {
	pairId := "TokenA/TokenB"
	tick := keeper.NewTick("TokenA/TokenB", 0, 6)
	s.app.DexKeeper.SetLimitOrderPoolReserveMap(s.ctx, types.LimitOrderPoolReserveMap{
		PairId:    pairId,
		Token:     "TokenB",
		TickIndex: tick.TickIndex,
		Count:     0,
		Reserves:  newDec("5"),
	})
	s.Equal(true, s.app.DexKeeper.HasToken1(s.ctx, &tick))
}

func (s *TickTestSuite) TestEmptyHasToken1() {
	pairId := "TokenA/TokenB"
	tick := keeper.NewTick("TokenA/TokenB", 0, 6)
	s.app.DexKeeper.SetLimitOrderPoolReserveMap(s.ctx, types.LimitOrderPoolReserveMap{
		PairId:    pairId,
		Token:     "TokenB",
		TickIndex: tick.TickIndex,
		Count:     0,
		Reserves:  newDec("0"),
	})
	s.Equal(false, s.app.DexKeeper.HasToken1(s.ctx, &tick))
}

func (s *TickTestSuite) TestLiquidityHasToken1() {
	pairId := "TokenA/TokenB"
	tick := keeper.NewTick("TokenA/TokenB", 0, 6)
	tick.TickData.Reserve0AndShares[0].Reserve0 = newDec("100")
	tick.TickData.Reserve0AndShares[0].TotalShares = newDec("100")
	tick.TickData.Reserve1[0] = newDec("20")
	s.app.DexKeeper.SetLimitOrderPoolReserveMap(s.ctx, types.LimitOrderPoolReserveMap{
		PairId:    pairId,
		Token:     "TokenB",
		TickIndex: tick.TickIndex,
		Count:     0,
		Reserves:  newDec("0"),
	})
	s.Equal(true, s.app.DexKeeper.HasToken1(s.ctx, &tick))
}

func (s *TickTestSuite) TestLiquidityHighFeeHasToken1() {
	pairId := "TokenA/TokenB"
	tick := keeper.NewTick("TokenA/TokenB", 0, 6)
	tick.TickData.Reserve0AndShares[4].Reserve0 = newDec("100")
	tick.TickData.Reserve0AndShares[4].TotalShares = newDec("100")
	tick.TickData.Reserve1[4] = newDec("20")
	s.app.DexKeeper.SetLimitOrderPoolReserveMap(s.ctx, types.LimitOrderPoolReserveMap{
		PairId:    pairId,
		Token:     "TokenB",
		TickIndex: tick.TickIndex,
		Count:     0,
		Reserves:  newDec("0"),
	})
	s.Equal(true, s.app.DexKeeper.HasToken1(s.ctx, &tick))
}

func (s *TickTestSuite) TestNoLiquidityOnOneSideHasToken1() {
	pairId := "TokenA/TokenB"
	tick := keeper.NewTick("TokenA/TokenB", 0, 6)
	tick.TickData.Reserve0AndShares[4].Reserve0 = newDec("100")
	tick.TickData.Reserve0AndShares[4].TotalShares = newDec("10")
	tick.TickData.Reserve1[4] = newDec("0")
	s.app.DexKeeper.SetLimitOrderPoolReserveMap(s.ctx, types.LimitOrderPoolReserveMap{
		PairId:    pairId,
		Token:     "TokenB",
		TickIndex: tick.TickIndex,
		Count:     0,
		Reserves:  newDec("0"),
	})
	s.Equal(false, s.app.DexKeeper.HasToken1(s.ctx, &tick))
}