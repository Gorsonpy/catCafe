package service

import (
	"time"

	"github.com/Gorsonpy/catCafe/biz/dal/es"
	"github.com/Gorsonpy/catCafe/biz/dal/mysql"
	"github.com/Gorsonpy/catCafe/biz/model/cat"
	"github.com/Gorsonpy/catCafe/pkg/errno"
)

func UpdateCat(cat *cat.CatModel) (int64, string) {
	c, err := mysql.QueryCat(cat.CatId)
	if err != nil {
		return errno.UpdateErrorCode, err.Error()
	}
	if cat.Name != "" {
		c.Name = cat.Name
	}
	if cat.Breed != "" {
		c.Breed = cat.Breed
	}
	if cat.Gender != "" {
		c.Gender = cat.Gender
	}
	if cat.Age != 0 {
		c.Age = cat.Age
	}
	if cat.HealthStatus != "" {
		c.HealthStatus = cat.HealthStatus
	}
	if cat.PhotoUrl != "" {
		c.PhotoUrl = cat.PhotoUrl
	}
	if cat.CheckInDate != "" {
		c.CheckInDate, err = time.Parse(time.DateTime, cat.CheckInDate)
		if err != nil {
			return errno.UpdateErrorCode, err.Error()
		}
	}

	err = mysql.UpdateCat(&c)
	if err != nil {
		return errno.UpdateErrorCode, err.Error()
	}
	err = es.UpdateCat(&c)
	if err != nil {
		return errno.UpdateErrorCode, err.Error()
	}

	return errno.StatusSuccessCode, errno.SuccessMsg
}
func DelCat(catId int64) (int64, string) {
	err := mysql.DelCat(catId)
	if err != nil {
		return errno.DelErrorCode, err.Error()
	}
	err = es.DelCat(catId)
	if err != nil {
		return errno.DelErrorCode, err.Error()
	}
	return errno.StatusSuccessCode, errno.SuccessMsg
}
func QueryCats(req *cat.QueryCatsReq) (int64, string, []*cat.CatModel) {
	cats, err := es.QueryCats(req)
	if err != nil {
		return errno.GetErrorCode, err.Error(), nil
	}
	list := es.CatToModel(cats)
	return errno.StatusSuccessCode, errno.SuccessMsg, list
}
func QueryTopCats(k int64) (int64, string, []*cat.CatModel) {
	cats := mysql.QueryTopCats(k)
	list := es.CatToModel(cats)
	return errno.StatusSuccessCode, errno.SuccessMsg, list
}

func AddCat(cat *cat.CatModel) (int64, string, int64) {
	var err error
	t, err := time.Parse(time.DateTime, cat.CheckInDate)
	if err != nil {
		return errno.CreateErrorCode, err.Error(), -1
	}

	esC := &es.Cat{
		CatId:          cat.CatId,
		Name:           cat.Name,
		Breed:          cat.Breed,
		Gender:         cat.Gender,
		Age:            cat.Age,
		HealthStatus:   cat.HealthStatus,
		PhotoUrl:       cat.PhotoUrl,
		CheckInDate:    t,
		AppointmentNum: cat.AppointmentNum,
	}
	err = mysql.AddCat(esC)
	if err != nil {
		return errno.CreateErrorCode, err.Error(), -1
	}

	err = es.AddCat(esC)
	if err != nil {
		return errno.CreateErrorCode, err.Error(), -1
	}

	return errno.StatusSuccessCode, errno.SuccessMsg, esC.CatId
}
