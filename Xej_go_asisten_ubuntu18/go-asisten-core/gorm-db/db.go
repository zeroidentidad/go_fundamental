package db

import (
	"github.com/jinzhu/gorm"

	// sqlite3 driver
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// OpenConnectionDB function
func OpenConnectionDB(folder string, lang string) (*gorm.DB, error) {
	return gorm.Open("sqlite3", folder+lang+".db")
}
