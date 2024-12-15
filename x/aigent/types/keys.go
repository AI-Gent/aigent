package types

const (
	// ModuleName defines the module name
	ModuleName = "aigent"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_aigent"
)

var (
	ParamsKey = []byte("p_aigent")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
