package data

import (
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var Base *gorm.DB

func ConnectDB() {
	var err error
	Base, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}
	Base.AutoMigrate(&Author{})
	Base.AutoMigrate(&Publisher{})
	Base.AutoMigrate(&Book{})
}
func LoadMock() {
	Base.Create(Authors)
	Base.Create(Publishers)
	Base.Create(Books)
}
