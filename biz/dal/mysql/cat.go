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

func DelCat(catId int64) error {
	return DB.Table("cats").Delete(&es.Cat{}, catId).Error
}

func QueryCat(catId int64) (es.Cat, error) {
	cat := es.Cat{}
	err := DB.Table("cats").Where("cat_id = ?", catId).Find(&cat).Error
	return cat, err
}

func UpdateCat(cat *es.Cat) error {
	return DB.Table("cats").Updates(&cat).Error
}
