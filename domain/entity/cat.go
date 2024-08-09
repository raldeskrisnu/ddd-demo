package entity

type Cat struct {
	Id      string  `gorm:"NOT NULL;COLUMN:id"`
	CatName string  `gorm:"NOT NULL;COLUMN:cat_name"`
	Price   float32 `gorm:"NOT NULL;COLUMN:price"`
}

func (Cat) TableName() string {
	return "cat"
}
