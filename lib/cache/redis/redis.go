package redis

import "hiframe/lib/cache"

type Redis struct {
	cache.BaseInterface
}

func (r *Redis) Set(key string, value interface{}, ttl_second int) (err error) {

	return
}

func (r *Redis) Has(key string) (err error) {

	return
}

func (r *Redis) Get(key string) (err error) {

	return
}

func (r *Redis) Del(key string) (err error) {

	return
}
