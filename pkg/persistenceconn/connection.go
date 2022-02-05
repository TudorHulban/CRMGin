package persistenceconn

import (
	"log"
	"os"
	"strings"

	"github.com/TudorHulban/GinCRM/cmd/setup"
	"github.com/TudorHulban/GinCRM/pkg/logic/authentication"
	"github.com/TudorHulban/GinCRM/pkg/ostop"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var theRDBMSConn *gorm.DB

// GetRDBMSConn Returns RDBMS connection handler.
func GetRDBMSConn() *gorm.DB {
	if theRDBMSConn == nil {
		if errMake := os.MkdirAll(setup.SQLiteFilePath, os.ModePerm); errMake != nil {
			log.Println("Could not create folder for SQLite database: ", errMake)
			os.Exit(ostop.SQLiteFolderCreation)
		}

		var pathSQLite string
		if setup.RDBMSModeTesting {
			path := strings.Split(setup.SQLiteFilePath, ".")
			pathSQLite = strings.Join([]string{path[0], authentication.UXNano(), ".", path[1]}, "")
		} else {
			pathSQLite = setup.SQLiteFilePath
		}

		var errCo error
		theRDBMSConn, errCo = gorm.Open(sqlite.Open(pathSQLite), &gorm.Config{
			DisableAutomaticPing: true,
		})
		if errCo != nil {
			log.Println("Could not create RDBMS Connection: ", errCo)
			os.Exit(ostop.RDBMSConnection)
		}
	}

	return theRDBMSConn
}
