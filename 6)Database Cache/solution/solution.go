package cache

import (
	"sync"
	"time"
)

// У нас есть key-value база данных, в которой хранятся пользователь.
// Эта база данных находится достаточно далеко от пользователей, из-за мы чего получаем.
// дополнительный latency в сотню миллисекунд. Мы хотим минимизировать.
// это время и начали думать в сторону кэширования...
//

// Нужно написать кэш для key-value базы данных, при этом важно учесть:
// – чтобы получился максимально эффективный и понятный код без багов.
// – чтобы, пользовательский код ничего не знал про кэширование.

var ()

type KVDatabase interface {
	Get(key string) (string, error)       // получить значение по ключу (используется пользователями очень часто)
	Keys() ([]string, error)              // получить все ключи (используется пользователями очень редко)
	MGet(keys []string) ([]string, error) // получить значения по ключам (используется пользователями очень редко)
}

type Cache struct {
	db          KVDatabase
	mu          sync.RWMutex
	data        map[string]string
	dataCreated time.Time
}

const invalidationTTL = time.Minute

func NewCache(db KVDatabase) *Cache {
	c := &Cache{
		db:          db,
		data:        make(map[string]string),
		dataCreated: time.Now(),
	}
	go c.startInvalidator()
	return c
}
func (c *Cache) startInvalidator() {
	ticker := time.NewTicker(invalidationTTL)
	defer ticker.Stop()

	for range ticker.C {
		c.mu.Lock()
		c.data = make(map[string]string)
		c.mu.Unlock()
	}

}
func (c *Cache) Get(key string) (string, error) {
	c.mu.RLock()
	v, ok := c.data[key]
	c.mu.RUnlock()
	if ok {
		return v, nil
	}
	v, err := c.db.Get(key)
	if err != nil {
		return "", err
	}

	c.mu.Lock()
	c.data[key] = v
	c.mu.Unlock()

	return v, nil

}
func (c *Cache) Keys() ([]string, error) {
	return c.db.Keys()
}
func (c *Cache) MGet(keys []string) ([]string, error) {
	return c.db.MGet(keys)
}
