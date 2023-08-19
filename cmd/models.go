package cmd

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
