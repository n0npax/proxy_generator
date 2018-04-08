package dbreader

import (
	"github.com/jinzhu/gorm"
	"github.com/n0npax/proxy_generator/parser"
	// load sqlite files
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func SqliteReadNginxRedirection(dbName string) []parser.NginxRedirection {
	db, err := gorm.Open("sqlite3", dbName)
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	db.AutoMigrate(&parser.NginxRedirection{})

	var redirections []parser.NginxRedirection
	db.Find(&redirections)

	return redirections
}
