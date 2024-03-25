package db

import (
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Database *gorm.DB

type Group struct {
	gorm.Model
	Code string
}

type Subject struct {
	gorm.Model
	Name string
}

type Pair struct {
	gorm.Model
	StartTime time.Time
	Group     Group
	GroupId   uint
	Subject   string
}

func Init() {
	db, err := gorm.Open(sqlite.Open("db.sqlite"), &gorm.Config{})

	if err != nil {
		panic("Cant open database")
	}

	db.AutoMigrate(&Group{}, &Subject{}, &Pair{})
	Database = db
}
