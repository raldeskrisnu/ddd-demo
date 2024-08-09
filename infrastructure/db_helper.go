package infrastructure

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"os"
)

type DBHelper struct {
	CatRepository CatRepository
	db            *gorm.DB
}

func InitDbHelper() (*DBHelper, error) {
	env := os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASSWORD") + "@tcp(" + os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT") + ")/" + os.Getenv("DB_NAME")
	db, err := gorm.Open("mysql", env)
	if err != nil {
		panic(err)
	}
	db.AutoMigrate()
	return &DBHelper{
		CatRepository: &CatRepositoryImpl{db},
		db:            db,
	}, nil
}
