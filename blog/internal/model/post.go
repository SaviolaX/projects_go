package model

type Post struct {
	Base
	Title      string
	Entry      string
	AuthorID   uint
	Author     User `gorm:"foreignKey:AuthorID"`
	CategoryID uint
	Category   Category
}
