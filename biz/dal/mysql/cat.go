package mysql

import (
	"github.com/Gorsonpy/catCafe/biz/dal/es"
)

func AddCat(cat *es.Cat) error {
	return DB.Table("cats").Create(&cat).Error
}

func QueryTopCats(k int64) []*es.Cat {
	cats := make([]*es.Cat, 0)
	DB.Table("cats").Limit(int(k)).Order("appointment_num DESC").Find(&cats)
	return cats
}
