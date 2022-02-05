package setup

// Possible persistence databases.
const (
	RDBMSSQLite = iota
	RDBMSPostgres
)

const (
	// RDBMSType type used.
	RDBMSType = 0

	// SQLiteFilePath Points to the OS SQLite file.
	SQLiteFilePath = "/home/tudi/ram/sqlite/gorm.db"

	// RDBMS Mode
	//
	// During test mode a UNIX nano timestamp is added to SQLite file name.
	RDBMSModeTesting = true
)
