package auth

import (
	"sync"
	"time"
)

type BlockListEntry struct {
	ExpiresAt time.Time
}

type TokenBlocklist struct {
	mu      sync.RWMutex
	entries map[string]BlockListEntry
}

var Blocklist = &TokenBlocklist{
	entries: make(map[string]BlockListEntry),
}

func (tb *TokenBlocklist) Add(tokenString string, expiresAt time.Time) {
	tb.mu.RLock()
	defer tb.mu.RUnlock()
	tb.entries[tokenString] = BlockListEntry{ExpiresAt: expiresAt}
}

func (tb *TokenBlocklist) IsBlocked(tokenString string) bool {
	tb.mu.RLock()
	defer tb.mu.RUnlock()

	entry, exists := tb.entries[tokenString]
	return exists && time.Now().Before(entry.ExpiresAt)
}

func (tb *TokenBlocklist) Cleanup() {
	tb.mu.Lock()
	defer tb.mu.Unlock()
	for token, entry := range tb.entries {
		if time.Now().After(entry.ExpiresAt) {
			delete(tb.entries, token)
		}
	}
}
