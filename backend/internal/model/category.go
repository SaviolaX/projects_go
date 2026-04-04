package model

type Category struct {
	Base
	Name  string
	Posts []Post
}
