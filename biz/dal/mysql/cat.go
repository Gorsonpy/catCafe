package mysql

import (
	"github.com/Gorsonpy/catCafe/biz/dal/es"
	"github.com/cloudwego/kitex/pkg/klog"
)

func AddCat(cat *es.Cat) error {
	klog.Info(cat.CatId)
	return DB.Table("cats").Create(&cat).Error
}
