package interfaces

type Hasher interface {
	Hash(key string) (string, error)
	Compare(key, toCompare string) bool
}