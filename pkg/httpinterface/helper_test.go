package httpinterface

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsIPv4GoodValues(t *testing.T) {
	var testCases = []struct {
		passedValue string
	}{
		{"0.0.0.0"},
		{"1.1.1.1"},
		{"127.0.0.1"},
	}

	t.Run("testing good values", func(t *testing.T) {
		for _, tc := range testCases {
			assert.Nil(t, isIpv4(tc.passedValue))
		}
	})
}

func TestIsIPv4BadValues(t *testing.T) {
	var testCases = []struct {
		passedValue string
		description string
	}{
		{"a", "malformed"},
		{"a.b.c.d", "not an IP"},
		{"0.1.0.1", "starting with zero"},
	}

	t.Run("testing malformed values", func(t *testing.T) {
		for _, tc := range testCases {
			assert.Error(t, isIpv4(tc.passedValue), tc.description)
		}
	})
}

// BenchmarkIsIPv4-4   	 2461770	       514 ns/op	   1.95 MB/s	      64 B/op	       1 allocs/op
func BenchmarkIsIPv4(b *testing.B) {
	b.ReportAllocs()
	b.SetBytes(1)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		isIpv4("127.0.0.1")
	}
}
