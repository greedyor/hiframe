package cache

type HICache struct {
}

type BaseInterface interface {
	Set(key string, value interface{}, ttl_second int64) error
	Has(key string) (bool, error)
	Get(key string) (interface{}, error)
	Del(key string) error
}
