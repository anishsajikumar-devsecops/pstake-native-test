package types

// Event types for the liquidstake module.
const (
	EventTypeMsgLiquidStake             = MsgTypeLiquidStake
	EventTypeMsgLiquidUnstake           = MsgTypeLiquidUnstake
	EventTypeMsgStakeToLP               = MsgTypeStakeToLP
	EventTypeMsgUpdateParams            = MsgTypeUpdateParams
	EventTypeAddLiquidValidator         = "add_liquid_validator"
	EventTypeRemoveLiquidValidator      = "remove_liquid_validator"
	EventTypeBeginRebalancing           = "begin_rebalancing"
	EventTypeAutocompound               = "autocompound"
	EventTypeUnbondInactiveLiquidTokens = "unbond_inactive_liquid_tokens"

	AttributeKeyDelegator             = "delegator"
	AttributeKeyNewShares             = "new_shares"
	AttributeKeyStkXPRTMintedAmount   = "stkxprt_minted_amount"
	AttributeKeyCompletionTime        = "completion_time"
	AttributeKeyUnbondingAmount       = "unbonding_amount"
	AttributeKeyUnbondedAmount        = "unbonded_amount"
	AttributeKeyLiquidValidator       = "liquid_validator"
	AttributeKeyRedelegationCount     = "redelegation_count"
	AttributeKeyRedelegationFailCount = "redelegation_fail_count"
	AttributeKeyLiquidAmount          = "liquid_amount"
	AttributeKeyStakedAmount          = "staked_amount"
	AttributeKeyPstakeAutocompoundFee = "pstake_autocompound_fee"

	AttributeKeyAuthority     = "authority"
	AttributeKeyUpdatedParams = "updated_params"

	AttributeValueCategory = ModuleName
)
