package cli

import (
	"strconv"

	"github.com/NicholasDotSol/duality/x/mev/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdSend() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "send [amount-in] [token-in]",
		Short: "Broadcast message send",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			argAmountIn := args[0]
			argTokenIn := args[1]

			argAmountInInt, ok := sdk.NewIntFromString(argAmountIn)

			if !ok {
				return sdkerrors.Wrapf(types.ErrIntOverflowTx, "AmountIn Overflower error")
			}
			msg := types.NewMsgSend(
				clientCtx.GetFromAddress().String(),
				argAmountInInt,
				argTokenIn,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
