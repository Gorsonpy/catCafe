package dal

import (
	"github.com/Gorsonpy/catCafe/biz/dal/es"
	"github.com/Gorsonpy/catCafe/biz/dal/mysql"
)

func Init() {
	mysql.Init()
	es.Init()
}
