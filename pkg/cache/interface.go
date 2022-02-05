package cache

import "github.com/TudorHulban/badgerwrap"

// IKV is the interface proposed for interacting with a key value based cache.
type IKV interface {
	// Inserts or updates KV in store.
	Set(badgerwrap.KV) error
	// Inserts or updates KV in store. Time To Live in seconds.
	SetTTL(badgerwrap.KV, uint) error
	// Inserts or updates KV in store.
	// Key is byte slice, value is to be serialized structure.
	SetAny([]byte, interface{}) error
	// Inserts or updates KV in store. Value is to be serialized structure.
	SetAnyTTL([]byte, interface{}, uint) error
	// Returns value for passed key if found. If not found it returns empty slice and an error not nil.
	GetVByK([]byte) ([]byte, error)
	// Fills up the passed pointer value for passed key if found. If not found it returns an error not nil.
	GetAnyByK([]byte, interface{}) error
	// Returns a slice of KV if prefix found.
	// If not found it returns empty slice.
	GetKVByPrefix([]byte) ([]badgerwrap.KV, error)
	// Deletes KV bassed on key.
	DeleteKVByK([]byte) error
	// Close closes the opened KV store.
	Close() error
}
