package cachecredentials

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/TudorHulban/log"
)

func TestCacheNoLog(t *testing.T) {
	_, errCache := GetCache()
	require.Error(t, errCache)
}

func TestCacheWithLog(t *testing.T) {
	c, errCache := GetCache(log.New(log.DEBUG, os.Stderr, true))
	require.NoError(t, errCache)
	assert.NoError(t, c.Close())
}
