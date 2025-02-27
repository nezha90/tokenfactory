package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/nezha90/tokenfactory/testutil/sample"
	"github.com/stretchr/testify/require"
)

func TestMsgMintAndSendTokens_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgMintAndSendTokens
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgMintAndSendTokens{
				Owner: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgMintAndSendTokens{
				Owner: sample.AccAddress(),
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
