package cache

import (
	"github.com/bradfitz/gomemcache/memcache"
)

type Memcache struct {
	mc *memcache.Client
}

func NewMemcache(servers []string) *Memcache {
	return &Memcache{
		mc: memcache.New(servers...),
	}
}

func (m *Memcache) Set(key string, value []byte, expireTime int) error {
	i := &memcache.Item{
		Key:        key,
		Value:      value,
		Expiration: int32(expireTime),
	}

	return m.mc.Set(i)
}

func (m *Memcache) Get(key string) ([]byte, error) {
	i, err := m.mc.Get(key)

	if err != nil {
		if err == memcache.ErrCacheMiss {
			return nil, nil
		}

		return nil, err
	}

	return i.Value, nil
}

func (m *Memcache) Ping() error {
	_, err := m.mc.Get("whatever")
	if err == memcache.ErrCacheMiss {
		return nil
	}

	return err
}
