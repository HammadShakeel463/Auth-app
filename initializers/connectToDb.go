package initializers

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDb() {
	//postgres://hbljtwvj:Omjba1yCMAK0aqYXUBpomI5bC956LCrq@rain.db.elephantsql.com/hbljtwvj
	var err error
	dsn := "host=rain.db.elephantsql.com user=hbljtwvj password=Omjba1yCMAK0aqYXUBpomI5bC956LCrq dbname=hbljtwvj port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect to db")
	}
}
