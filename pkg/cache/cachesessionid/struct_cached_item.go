package cachesessionid

// CachedItem Represents the cached item type for the cache.
type CachedItem struct {
	UserProfileID uint8  // ID of security profile
	UserID        uint64 // primary key as per user table.
	SessionID     string // KEY
}
