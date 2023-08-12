package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"os"
)

func main() {
	db := provideDB()
	err := db.Save(&Book{Name: "Harry Potter", Author: "J.K. Rowling", ID: 1}).Error
	if err != nil {
		panic(err)
	}
	err = db.Save(&Book{Name: "The Lord of the Rings", Author: "J.R.R. Tolkien", ID: 2}).Error
	if err != nil {
		panic(err)
	}
	err = db.Save(&User{Name: "John Doe", Username: "johndoe", Password: "1234", ID: 1}).Error
	if err != nil {
		panic(err)
	}
	err = db.Save(&User{Name: "Cap", Username: "admin", Password: "123", ID: 2}).Error
	if err != nil {
		panic(err)
	}
}

func provideDB() *gorm.DB {
	driver := os.Getenv("DB_DRIVER")
	dsn := os.Getenv("DB_DSN")
	var db *gorm.DB
	var err error
	switch driver {
	case "mysql":
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			panic("failed to connect database mysql")
		}
	case "postgres":
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			panic("failed to connect database postgres")
		}
	case "mssql":
		db, err = gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
		if err != nil {
			panic("failed to connect database sqlserver")
		}
	default:
		panic("invalid db driver")
	}
	err = db.AutoMigrate(Book{}, User{})
	if err != nil {
		panic("failed to migrate database")
	}
	return db
}

type Book struct {
	ID     int `gorm:"primaryKey"`
	Name   string
	Author string
}

type User struct {
	ID       int `gorm:"primaryKey"`
	Name     string
	Username string `gorm:"unique"`
	Password string
}
