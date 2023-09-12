package cancun

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

var (
	// EIP 4844
	GAS_PER_BLOB = uint64(0x20000)

	MIN_DATA_GASPRICE         = uint64(1)
	MAX_BLOB_GAS_PER_BLOCK    = uint64(786432)
	TARGET_BLOB_GAS_PER_BLOCK = uint64(393216)

	TARGET_BLOBS_PER_BLOCK = uint64(TARGET_BLOB_GAS_PER_BLOCK / GAS_PER_BLOB)
	MAX_BLOBS_PER_BLOCK    = uint64(MAX_BLOB_GAS_PER_BLOCK / GAS_PER_BLOB)

	BLOB_GASPRICE_UPDATE_FRACTION = uint64(3338477)

	BLOB_COMMITMENT_VERSION_KZG = byte(0x01)

	// EIP 4788
	HISTORY_STORAGE_ADDRESS  = common.HexToAddress("0x000000000000000000000000000000000000000b")
	HISTORICAL_ROOTS_MODULUS = uint64(98304)

	// Test constants
	DATAHASH_START_ADDRESS = big.NewInt(0x20000)
	DATAHASH_ADDRESS_COUNT = 1000
)
