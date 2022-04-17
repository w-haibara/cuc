package db

import (
	"path/filepath"

	"github.com/glebarez/sqlite"
	"github.com/w-haibara/cuc/config"
	"gorm.io/gorm"
)

var dbFileName = filepath.Join(config.Dir, "cuc.db")

func init() {
	db, err := gorm.Open(sqlite.Open(dbFileName), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	if err := db.AutoMigrate(&DB{}); err != nil {
		panic(err.Error())
	}
}

type DB struct {
	gorm.Model
	ClickUp ClickUp
}

type ClickUp struct {
	TeamID   string
	SpaceID  string
	FolderID string
}

func Updates(cu ClickUp) {
	db, err := gorm.Open(sqlite.Open(dbFileName), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	db.Model(&DB{}).Updates(DB{ClickUp: cu})
}

func FetchFolderID() string {
	db, err := gorm.Open(sqlite.Open(dbFileName), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	var v DB
	db.First(&v)

	return v.ClickUp.FolderID
}
