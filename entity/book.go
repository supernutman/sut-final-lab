package entity

import (
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title 	string	 `valid:"length(3|100)~Title len 3 to 100"`
	Price 	float64	 `valid:"range(50|5000)~price must between 50 and 5000"`
	Code 	string	 `valid:"matches(^BK\\d{6}$)~code is invalid"` 	 
}
