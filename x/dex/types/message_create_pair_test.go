package types

import (
	"testing"

	"github.com/NicholasDotSol/duality/testutil/sample"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestMsgCreatePair_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgCreatePair
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgCreatePair{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgCreatePair{
				Creator: sample.AccAddress(),
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