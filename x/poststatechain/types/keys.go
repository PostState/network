package types

const (
	// ModuleName defines the module name
	ModuleName = "poststatechain"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_poststatechain"
)

var (
	ParamsKey = []byte("p_poststatechain")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
