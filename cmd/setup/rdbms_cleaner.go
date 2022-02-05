package setup

import (
	"log"
	"os"

	"github.com/TudorHulban/GinCRM/pkg/ostop"
)

/*
File contains RDBMS cleaner.
Implementation for SQLite included.
*/

// CleanRDBMS Checks database used and performs cleaning activities.
// Exported so it can be used in testing in other packages.
//
// Do not use for testing.
func CleanRDBMS() {
	switch RDBMSType {
	case RDBMSSQLite:
		go log.Println("Cleaning SQLite.")

		sqliteCleaner()
	}
}

func sqliteCleaner() {
	if _, errExists := os.Stat(SQLiteFilePath); !os.IsNotExist(errExists) {
		if errCleanSQLite := os.Remove(SQLiteFilePath); errCleanSQLite != nil {
			log.Println("Error removing previous SQLite database file, error: ", errCleanSQLite)
			os.Exit(ostop.SQLiteCleanUp)
		}

		go log.Println("Cleaned file:", SQLiteFilePath)
		return
	}

	go log.Println("Did not find to clean file:", SQLiteFilePath)
}
