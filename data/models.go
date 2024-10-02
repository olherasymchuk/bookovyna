package data

import "gorm.io/gorm"

type Author struct {
	gorm.Model
	Name    string
	Surname string
}
type Publisher struct {
	gorm.Model
	Name string
}
type Book struct {
	gorm.Model
	Title        string
	Author_ID    uint32
	Price        float32
	Publisher_ID uint32
	Published    uint32
	ISBN         string
}
