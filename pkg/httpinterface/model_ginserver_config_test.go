package httpinterface

import (
	"os"
	"testing"

	"github.com/TudorHulban/log"
	"github.com/stretchr/testify/assert"
)

func TestCreateConfig(t *testing.T) {
	const vers = "0.1"

	co, err := CreateConfig("0.0.0.0:80", vers, log.New(log.DEBUG, os.Stderr, true), 3)
	assert.NoError(t, err)

	assert.Equal(t, co.IPV4Address, "0.0.0.0", "testing IPV4")
	assert.Equal(t, co.Port, uint16(80), "testing port")
	assert.Equal(t, co.BinaryVersion, vers)
}
