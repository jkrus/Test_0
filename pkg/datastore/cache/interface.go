package cache

type (
	Cache interface {
		Set(k []byte, v []byte)
		Get(dst []byte, k []byte) []byte
	}
)
