package types

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"cosmossdk.io/errors"
	"github.com/cosmos/cosmos-sdk/types/bech32"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	icatypes "github.com/cosmos/ibc-go/v7/modules/apps/27-interchain-accounts/types"
	host "github.com/cosmos/ibc-go/v7/modules/core/24-host"
	ibcexported "github.com/cosmos/ibc-go/v7/modules/core/exported"
	liquidstakeibctypes "github.com/persistenceOne/pstake-native/v2/x/liquidstakeibc/types"
)

func (hc HostChain) ValidateBasic() error {
	err := host.ConnectionIdentifierValidator(hc.ConnectionId)
	if !(err == nil || hc.ConnectionId == ibcexported.LocalhostConnectionID) {
		return errors.Wrapf(sdkerrors.ErrInvalidRequest, "hostchain connectionID invalid")
	}

	if hc.IcaAccount.Owner != "" {
		portID, err := icatypes.NewControllerPortID(hc.IcaAccount.Owner)
		if err != nil {
			return err
		}
		err = host.PortIdentifierValidator(portID)
		if err != nil {
			return err
		}
	}

	switch hc.IcaAccount.ChannelState {
	case liquidstakeibctypes.ICAAccount_ICA_CHANNEL_CREATING:
		if hc.IcaAccount.Address != "" {
			return fmt.Errorf("ica account address for ICAAccount_ICA_CHANNEL_CREATING should be empty")
		}
		// No features allowed without ICA account.
		if hc.Features.LiquidStake.Enabled == true || hc.Features.LiquidStakeIBC.Enabled == true {
			return fmt.Errorf("no features should be enabled without a valid ICA account")
		}
	case liquidstakeibctypes.ICAAccount_ICA_CHANNEL_CREATED:
		if hc.IcaAccount.Address == "" {
			return fmt.Errorf("ica account address for ICAAccount_ICA_CHANNEL_CREATED should not be empty")
		}
		_, _, err = bech32.DecodeAndConvert(hc.IcaAccount.Address)
		if err != nil {
			return err
		}
	}

	err = hc.Features.ValdidateBasic()
	if err != nil {
		return err
	}

	return nil
}

func (hc HostChain) IsActive() bool {
	if hc.Features.LiquidStakeIBC.Enabled == true ||
		hc.Features.LiquidStake.Enabled == true {
		return true
	}
	return false
}

func (f Feature) ValdidateBasic() error {
	if f.LiquidStakeIBC.FeatureType != FeatureType_LIQUID_STAKE_IBC {
		return fmt.Errorf("invalid feature type expected %s, got %s", FeatureType_LIQUID_STAKE_IBC, f.LiquidStakeIBC.FeatureType)
	}
	err := f.LiquidStakeIBC.ValdidateBasic()
	if err != nil {
		return err
	}

	if f.LiquidStake.FeatureType != FeatureType_LIQUID_STAKE {
		return fmt.Errorf("invalid feature type expected %s, got %s", FeatureType_LIQUID_STAKE, f.LiquidStake.FeatureType)
	}
	err = f.LiquidStake.ValdidateBasic()
	if err != nil {
		return err
	}
	return nil
}

func (lsConfig LiquidStake) ValdidateBasic() error {
	if lsConfig.CodeID == 0 {
		if lsConfig.Instantiation != InstantiationState_INSTANTIATION_NOT_INITIATED {
			return fmt.Errorf("config with 0 code id should not have been initiated")
		}
	}
	switch lsConfig.Instantiation {
	case InstantiationState_INSTANTIATION_NOT_INITIATED:
		if lsConfig.ContractAddress != "" {
			return fmt.Errorf("InstantiationState_INSTANTIATION_NOT_INITIATED cannot have contract address")
		}
		if lsConfig.Enabled == true {
			return fmt.Errorf("feature cannot be turned on without instantiation complete")
		}
	case InstantiationState_INSTANTIATION_INITIATED:
		if lsConfig.ContractAddress != "" {
			return fmt.Errorf("InstantiationState_INSTANTIATION_INITIATED cannot have contract address")
		}
		if lsConfig.CodeID == 0 {
			return fmt.Errorf("InstantiationState_INSTANTIATION_INITIATED cannot have 0 codeID")
		}
		if lsConfig.Enabled == true {
			return fmt.Errorf("feature cannot be turned on without instantiation complete")
		}
	case InstantiationState_INSTANTIATION_COMPLETED:
		if lsConfig.ContractAddress == "" {
			return fmt.Errorf("InstantiationState_INSTANTIATION_COMPLETED cannot have empty contract address")
		}
		_, _, err := bech32.DecodeAndConvert(lsConfig.ContractAddress)
		if err != nil {
			return err
		}
		if lsConfig.CodeID == 0 {
			return fmt.Errorf("InstantiationState_INSTANTIATION_COMPLETED cannot have 0 codeID")
		}
	}
	err := ValidateLiquidStakeDenoms(lsConfig.Denoms)
	if err != nil {
		return err
	}
	return nil
}
func (lsConfig LiquidStake) AllowsAllDenoms() bool {
	if len(lsConfig.Denoms) == 1 && lsConfig.Denoms[0] == LiquidStakeAllowAllDenoms {
		return true
	}
	return false
}
func (lsConfig LiquidStake) AllowsDenom(denom string) bool {
	if lsConfig.AllowsAllDenoms() {
		return true
	}
	return slices.Contains(lsConfig.Denoms, denom)
}
func (lsConfig LiquidStake) Equals(l2 LiquidStake) bool {
	if lsConfig.CodeID != l2.CodeID {
		return false
	}
	if lsConfig.Instantiation != l2.Instantiation {
		return false
	}
	if lsConfig.ContractAddress != l2.ContractAddress {
		return false
	}
	if !slices.Equal(lsConfig.Denoms, l2.Denoms) {
		return false
	}
	if lsConfig.FeatureType != l2.FeatureType {
		return false
	}
	if lsConfig.Enabled != l2.Enabled {
		return false
	}
	return true
}

func MustICAPortIDfromOwner(owner string) string {
	id, err := icatypes.NewControllerPortID(owner)
	if err != nil {
		panic(err)
	}
	return id

}

func DefaultPortOwner(id uint64) string {
	return fmt.Sprintf("%s%v", DefaultPortOwnerPrefix, id)
}
func OwnerfromPortID(portID string) (string, error) {
	prefix := fmt.Sprintf("%s", icatypes.ControllerPortPrefix)
	idStr, found := strings.CutPrefix(portID, prefix)
	if !found {
		return "", fmt.Errorf("invalid portID, expect prefix %s", prefix)
	}

	return idStr, nil
}

func IDfromPortID(portID string) (uint64, error) {
	prefix := fmt.Sprintf("%s%s", icatypes.ControllerPortPrefix, DefaultPortOwnerPrefix)
	idStr, found := strings.CutPrefix(portID, prefix)
	if !found {
		return 0, fmt.Errorf("invalid portID, expect prefix %s", prefix)
	}
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func ValidateLiquidStakeDenoms(denoms []string) error {
	if len(denoms) == 1 && denoms[0] == LiquidStakeAllowAllDenoms {
		return nil
	}
	for _, denom := range denoms {
		if !liquidstakeibctypes.IsLiquidStakingDenom(denom) {
			return fmt.Errorf("invalid denom, expected a liquidstaking denom got %s", denom)
		}
	}
	return nil
}
