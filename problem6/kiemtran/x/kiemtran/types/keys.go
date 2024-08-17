package types

const (
	// ModuleName defines the module name
	ModuleName = "kiemtran"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_kiemtran"
)

var (
	ParamsKey = []byte("p_kiemtran")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
