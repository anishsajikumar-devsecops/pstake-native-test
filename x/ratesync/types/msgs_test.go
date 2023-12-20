package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	ibcexported "github.com/cosmos/ibc-go/v7/modules/core/exported"
	"github.com/persistenceOne/pstake-native/v2/x/liquidstakeibc/types"
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/stretchr/testify/require"
)

var ValidHostChainInMsg = func(id uint64) HostChain {
	return HostChain{
		ID:           id,
		ChainID:      "test-1",
		ConnectionID: ibcexported.LocalhostConnectionID,
		ICAAccount: types.ICAAccount{
			Address:      "",
			Balance:      sdk.Coin{},
			Owner:        "",
			ChannelState: 0,
		},
		Features: Feature{
			LiquidStakeIBC: LiquidStake{
				FeatureType:     0,
				CodeID:          0,
				Instantiation:   0,
				ContractAddress: "",
				Denoms:          []string{},
				Enabled:         false,
			},
			LiquidStake: LiquidStake{
				FeatureType:     1,
				CodeID:          0,
				Instantiation:   0,
				ContractAddress: "",
				Denoms:          nil,
				Enabled:         false,
			}},
	}
}

func TestMsgUpdateParams_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgUpdateParams
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgUpdateParams{
				Authority: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgUpdateParams{
				Authority: authtypes.NewModuleAddress("addr").String(),
				Params:    DefaultParams(),
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

func TestMsgCreateHostChain_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgCreateHostChain
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgCreateHostChain{
				Authority: "invalid_address",
				HostChain: ValidHostChainInMsg(0),
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgCreateHostChain{
				Authority: authtypes.NewModuleAddress("addr1").String(),
				HostChain: ValidHostChainInMsg(0),
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

func TestMsgUpdateHostChain_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgUpdateHostChain
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgUpdateHostChain{
				Authority: "invalid_address",
				HostChain: ValidHostChainInMsg(1),
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgUpdateHostChain{
				Authority: authtypes.NewModuleAddress("addr1").String(),
				HostChain: ValidHostChainInMsg(1),
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

func TestMsgDeleteHostChain_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgDeleteHostChain
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgDeleteHostChain{
				Authority: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgDeleteHostChain{
				Authority: authtypes.NewModuleAddress("addr1").String(),
				ID:        1,
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
