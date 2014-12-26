package cache

type Cacher interface {
	Set(key string, value []byte, expireTime int) error
	Get(key string) ([]byte, error)
}
