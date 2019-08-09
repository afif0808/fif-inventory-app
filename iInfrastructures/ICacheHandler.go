package iInfrastructures

type ICacheHandler interface {
	// SET AND UPDATE CACHE
	SetCache(key string, value interface{}) error
	// GET CACHE
	GetCache(key string) (interface{}, error)
	// DELETE CACHE
	DeleteCache(key string) error
}
