package cache

import (
	"bytes"
	"context"
	"database/sql/driver"
	"encoding/gob"
	"errors"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/golang/groupcache/lru"
)

type (

	// A Key defines a comparable Go value.
	// See http://golang.org/ref/spec#Comparison_operators
	Key any

	Entry struct {
		Columns []string
		Values  [][]driver.Value
	}
	// AddGetDeleter defines the interface for getting,
	// adding and deleting entries from the cache.
	AddGetDeleter interface {
		Clear(context.Context) error
		Del(context.Context, Key) error
		Add(context.Context, Key, Entry, time.Duration) error
		Get(context.Context, Key) (Entry, error)
	}
)

func init() {
	// Register non builtin driver.Values.
	gob.Register(time.Time{})
}

// MarshalBinary implements the encoding.BinaryMarshaler interface.
func MarshalBinary(data any) ([]byte, error) {
	var buf bytes.Buffer
	if err := gob.NewEncoder(&buf).Encode(data); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// UnmarshalBinary implements the encoding.BinaryUnmarshaler interface.
func UnmarshalBinary(buf []byte, out *Entry) error {
	type entry struct {
		Columns []string
		Values  [][]driver.Value
	}
	e := new(entry)
	if err := gob.NewDecoder(bytes.NewBuffer(buf)).Decode(&e); err != nil {
		return err
	}
	out.Columns = e.Columns
	out.Values = e.Values
	return nil
}

// ErrNotFound is returned by Get when and Entry does not exist in the cache.
var ErrNotFound = errors.New("entcache: entry was not found")

type (
	// LRU provides an LRU cache that implements the AddGetter interface.
	LRU struct {
		*lru.Cache
	}
	// entry wraps the Entry with additional expiry information.
	entry struct {
		buffer Entry
		expiry time.Time
	}
)

// New creates a new Cache.
// If maxEntries is zero, the cache has no limit.
func NewLRU(maxEntries int) *LRU {
	return &LRU{
		Cache: lru.New(maxEntries),
	}
}

// Add adds the entry to the cache.
func (l *LRU) Add(_ context.Context, k Key, data Entry, ttl time.Duration) error {
	defer func() {
		if p := recover(); p != nil {
			fmt.Printf("add cache error: %v", p)
		}
	}()
	if ttl == 0 {
		l.Cache.Add(k, data)
	} else {
		l.Cache.Add(k, &entry{buffer: data, expiry: time.Now().Add(ttl)})
	}
	return nil
}

// Get gets an entry from the cache.
func (l *LRU) Get(_ context.Context, k Key) (Entry, error) {
	defer func() {
		if p := recover(); p != nil {
			fmt.Printf("get cache error: %v", p)
		}
	}()
	e, ok := l.Cache.Get(k)
	if !ok {
		return Entry{}, ErrNotFound
	}
	switch e := e.(type) {
	case Entry:
		fmt.Println("load data from lrn cache.")
		return e, nil
	case *entry:
		if time.Now().Before(e.expiry) {
			fmt.Println("load data from lrn cache.")
			return e.buffer, nil
		}
		l.Cache.Remove(k)
		return Entry{}, ErrNotFound
	default:
		return Entry{}, fmt.Errorf("entcache: unexpected entry type: %T", e)
	}
}

// Del deletes an entry from the cache.
func (l *LRU) Del(_ context.Context, k Key) error {
	l.Cache.Remove(k)
	return nil
}

// clear all cache.
func (l *LRU) Clear(_ context.Context) error {
	fmt.Println("clear lrn sql cache.")
	l.Cache.Clear()
	return nil
}

// Redis provides a remote cache backed by Redis
// and implements the SetGetter interface.
type Redis struct {
	C redis.Cmdable
}

// NewRedis returns a new Redis cache level from the given Redis connection.
//
//	entcache.NewRedis(redis.NewClient(&redis.Options{
//		Addr: ":6379"
//	}))
//
//	entcache.NewRedis(redis.NewClusterClient(&redis.ClusterOptions{
//		Addrs: []string{":7000", ":7001", ":7002"},
//	}))
//
func NewRedis(c redis.Cmdable) *Redis {
	return &Redis{C: c}
}

// Add adds the entry to the cache.
func (r *Redis) Add(ctx context.Context, k Key, data Entry, ttl time.Duration) error {
	key := fmt.Sprint(k)
	if key == "" {
		return nil
	}
	buf, err := MarshalBinary(data)
	if err != nil {
		return err
	}
	if err := r.C.Set(ctx, key, buf, ttl).Err(); err != nil {
		return err
	}
	return nil
}

// Get gets an entry from the cache.
func (r *Redis) Get(ctx context.Context, k Key) (Entry, error) {
	key := fmt.Sprint(k)
	if key == "" {
		return Entry{}, ErrNotFound
	}
	buf, err := r.C.Get(ctx, key).Bytes()
	if err != nil || len(buf) == 0 {
		return Entry{}, ErrNotFound
	}
	out := new(Entry)
	if err := UnmarshalBinary(buf, out); err != nil {
		return Entry{}, err
	}
	fmt.Println("load data from redis cache.")
	return *out, nil
}

// Del deletes an entry from the cache.
func (r *Redis) Del(ctx context.Context, k Key) error {
	key := fmt.Sprint(k)
	if key == "" {
		return nil
	}
	return r.C.Del(ctx, key).Err()
}

// clear all cache.
func (r *Redis) Clear(ctx context.Context) error {
	fmt.Println("clear redis sql cache.")
	return r.C.FlushAll(ctx).Err()
}

// multiLevel provides a multi-level cache implementation.
type multiLevel struct {
	levels []AddGetDeleter
}

// Add adds the entry to the cache.
func (m *multiLevel) Add(ctx context.Context, k Key, e Entry, ttl time.Duration) error {
	for i := range m.levels {
		if err := m.levels[i].Add(ctx, k, e, ttl); err != nil {
			return err
		}
	}
	return nil
}

// Get gets an entry from the cache.
func (m *multiLevel) Get(ctx context.Context, k Key) (Entry, error) {
	for i := range m.levels {
		return m.levels[i].Get(ctx, k)
	}
	return Entry{}, ErrNotFound
}

// Del deletes an entry from the cache.
func (m *multiLevel) Del(ctx context.Context, k Key) error {
	for i := range m.levels {
		if err := m.levels[i].Del(ctx, k); err != nil {
			return err
		}
	}
	return nil
}

// Clear all cache.
func (m *multiLevel) Clear(ctx context.Context) error {
	for i := range m.levels {
		if err := m.levels[i].Clear(ctx); err != nil {
			return err
		}
	}
	return nil
}

// contextLevel provides a context/request level cache implementation.
type contextLevel struct{}

// Get gets an entry from the cache.
func (*contextLevel) Get(ctx context.Context, k Key) (Entry, error) {
	c, ok := FromContext(ctx)
	if !ok {
		return Entry{}, ErrNotFound
	}
	return c.Get(ctx, k)
}

// Add adds the entry to the cache.
func (*contextLevel) Add(ctx context.Context, k Key, e Entry, ttl time.Duration) error {
	c, ok := FromContext(ctx)
	if !ok {
		return nil
	}
	return c.Add(ctx, k, e, ttl)
}

// Del deletes an entry from the cache.
func (*contextLevel) Del(ctx context.Context, k Key) error {
	c, ok := FromContext(ctx)
	if !ok {
		return nil
	}
	return c.Del(ctx, k)
}

// clear all cache.
func (*contextLevel) Clear(ctx context.Context) error {
	c, ok := FromContext(ctx)
	if !ok {
		return nil
	}
	return c.Clear(ctx)
}
