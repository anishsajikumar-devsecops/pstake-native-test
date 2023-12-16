package types

import "encoding/binary"

const (
	// ModuleName defines the module name
	ModuleName = "ratesync"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_ratesync"

	HostChainIDKeyPrefix = "host_chain_id"
	HostChainKeyPrefix   = "host_chain"
	ParamsKeyPrefix      = "params"

	LiquidStakeAllowAllDenoms = "*"
	LiquidStakeEpoch          = "day"
	DefaultPortOwnerPrefix    = "pstake_ratesync_"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

func HostChainIDKey() []byte {
	return KeyPrefix(HostChainIDKeyPrefix)
}

// HostChainKey returns the store key to retrieve a Chain from the index fields
func HostChainKey(
	id uint64,
) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return append(KeyPrefix(HostChainKeyPrefix), bz...)
}
