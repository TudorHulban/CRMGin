package persistenceconn_test

import (
	"testing"

	"github.com/TudorHulban/GinCRM/cmd/setup"
	"github.com/TudorHulban/GinCRM/pkg/persistenceconn"
	"github.com/stretchr/testify/assert"
)

type TableA struct {
	ID          uint8  `gorm:"primaryKey"`
	Description string `validate:"required"`
}

type TableB struct {
	ID          uint8  `gorm:"primaryKey"`
	Description string `validate:"required"`
}

// TestTableCreation Targets the capacity of persisting one structure at a time.
func TestTableCreation(t *testing.T) {
	setup.CleanRDBMS()

	db := persistenceconn.GetRDBMSConn()
	assert.NoError(t, db.AutoMigrate(TableA{}))
	assert.True(t, db.Migrator().HasTable(TableA{}))
}

// TestTablesCreation Targets the capacity of persisting several structures at a time.
func TestTablesCreation(t *testing.T) {
	setup.CleanRDBMS()

	db := persistenceconn.GetRDBMSConn()
	assert.NoError(t, db.AutoMigrate(TableA{}, TableB{}))
	assert.True(t, db.Migrator().HasTable(TableA{}))
	assert.True(t, db.Migrator().HasTable(TableB{}))
}
