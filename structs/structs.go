package structs

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	UserName string
	Parent int
}