package iInfrastructures

type ICacheHandler interface {
	// SET AND UPDATE CACHE
	SetCache(key string, value []byte) error
	// GET CACHE
	GetCache(key string) ([]byte, error)
}
