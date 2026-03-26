package model

type User struct {
	Base
	Username string
	Email    string
	Password string
	Posts    []Post `gorm:"foreignKey:AuthorID"`
}
