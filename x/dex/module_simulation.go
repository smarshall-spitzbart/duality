package dex

import (
	"math/rand"

	"github.com/NicholasDotSol/duality/testutil/sample"
	dexsimulation "github.com/NicholasDotSol/duality/x/dex/simulation"
	"github.com/NicholasDotSol/duality/x/dex/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = dexsimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgDeposit = "op_weight_msg_deposit"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeposit int = 100

	opWeightMsgWithdrawl = "op_weight_msg_withdrawl"
	// TODO: Determine the simulation weight value
	defaultWeightMsgWithdrawl int = 100

	opWeightMsgSwap = "op_weight_msg_swap"
	// TODO: Determine the simulation weight value
	defaultWeightMsgSwap int = 100

	opWeightMsgPlaceLimitOrder = "op_weight_msg_place_limit_order"
	// TODO: Determine the simulation weight value
	defaultWeightMsgPlaceLimitOrder int = 100

	opWeightMsgWithdrawFilledLimitOrder = "op_weight_msg_withdrawl_withdrawn_limit_order"
	// TODO: Determine the simulation weight value
	defaultWeightMsgWithdrawFilledLimitOrder int = 100

	opWeightMsgCancelLimitOrder = "op_weight_msg_cancel_limit_order"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCancelLimitOrder int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	dexGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&dexGenesis)
}

// ProposalContents doesn't return any content functions for governance proposals
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// RandomizedParams creates randomized  param changes for the simulator
func (am AppModule) RandomizedParams(_ *rand.Rand) []simtypes.ParamChange {

	return []simtypes.ParamChange{}
}

// RegisterStoreDecoder registers a decoder
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgDeposit int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeposit, &weightMsgDeposit, nil,
		func(_ *rand.Rand) {
			weightMsgDeposit = defaultWeightMsgDeposit
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeposit,
		dexsimulation.SimulateMsgDeposit(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgWithdrawl int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgWithdrawl, &weightMsgWithdrawl, nil,
		func(_ *rand.Rand) {
			weightMsgWithdrawl = defaultWeightMsgWithdrawl
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgWithdrawl,
		dexsimulation.SimulateMsgWithdrawl(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgSwap int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgSwap, &weightMsgSwap, nil,
		func(_ *rand.Rand) {
			weightMsgSwap = defaultWeightMsgSwap
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgSwap,
		dexsimulation.SimulateMsgSwap(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgPlaceLimitOrder int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgPlaceLimitOrder, &weightMsgPlaceLimitOrder, nil,
		func(_ *rand.Rand) {
			weightMsgPlaceLimitOrder = defaultWeightMsgPlaceLimitOrder
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgPlaceLimitOrder,
		dexsimulation.SimulateMsgPlaceLimitOrder(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgWithdrawFilledLimitOrder int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgWithdrawFilledLimitOrder, &weightMsgWithdrawFilledLimitOrder, nil,
		func(_ *rand.Rand) {
			weightMsgWithdrawFilledLimitOrder = defaultWeightMsgWithdrawFilledLimitOrder
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgWithdrawFilledLimitOrder,
		dexsimulation.SimulateMsgWithdrawFilledLimitOrder(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCancelLimitOrder int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCancelLimitOrder, &weightMsgCancelLimitOrder, nil,
		func(_ *rand.Rand) {
			weightMsgCancelLimitOrder = defaultWeightMsgCancelLimitOrder
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCancelLimitOrder,
		dexsimulation.SimulateMsgCancelLimitOrder(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
