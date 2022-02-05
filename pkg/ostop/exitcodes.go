package ostop

/*
File contains app OS exit codes.
Codes after 100 are less important.
*/

const (
	// RUNTIME Error.
	RUNTIME = 10
	// CACHE Could not create cache.
	CACHE = 11

	// RDBMSConnection Could not create RDBMS connection.
	RDBMSConnection = 12
	// RDBMSPopulationOfSchema Could not create RDBMS connection.
	RDBMSPopulationOfSchema = 14

	/* Less Important */

	// SQLiteCleanUp could not be performed
	SQLiteCleanUp = 101
	// SQLiteFolderCreation Could not create folder for SQLite database.
	SQLiteFolderCreation = 102
)
