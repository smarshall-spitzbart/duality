package types

import (
	"testing"

	"github.com/NicholasDotSol/duality/testutil/sample"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestMsgCancelLimitOrder_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgCancelLimitOrder
		err  error
	}{
		{
			name: "invalid creator",
			msg: MsgCancelLimitOrder{
				Creator:   "invalid_address",
				Receiver:  sample.AccAddress(),
				TokenA:    "TokenA",
				TokenB:    "TokenB",
				TickIndex: 0,
				KeyToken:  "TokenA",
				Key:       0,
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "invalid receiver",
			msg: MsgCancelLimitOrder{
				Creator:   sample.AccAddress(),
				Receiver:  "invalid address",
				TokenA:    "TokenA",
				TokenB:    "TokenB",
				TickIndex: 0,
				KeyToken:  "TokenA",
				Key:       0,
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "invalid key token",
			msg: MsgCancelLimitOrder{
				Creator:   sample.AccAddress(),
				Receiver:  sample.AccAddress(),
				TokenA:    "TokenA",
				TokenB:    "TokenB",
				TickIndex: 0,
				KeyToken:  "TokenC",
				Key:       0,
			},
			err: ErrInvalidKeyToken,
		}, {
			name: "valid msg",
			msg: MsgCancelLimitOrder{
				Creator:   sample.AccAddress(),
				Receiver:  sample.AccAddress(),
				TokenA:    "TokenA",
				TokenB:    "TokenB",
				TickIndex: 0,
				KeyToken:  "TokenA",
				Key:       0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
