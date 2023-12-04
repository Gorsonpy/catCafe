package service

import (
	"time"

	"github.com/Gorsonpy/catCafe/biz/dal/mysql"
	"github.com/Gorsonpy/catCafe/biz/model/appointment"
	"github.com/Gorsonpy/catCafe/pkg/errno"
	"gorm.io/gorm"
)

func QueryApp(customerId int64, req *appointment.QueryAppointmentReq) (int64, string, []*appointment.AppointmentModel) {
	list, err := mysql.QueryApp(customerId, req.StartTime, req.EndTime)
	if err != gorm.ErrRecordNotFound && err != nil {
		return errno.GetErrorCode, err.Error(), nil
	}

	models := make([]*appointment.AppointmentModel, 0)
	for _, app := range list {
		model := mysql.AppToModel(app)
		models = append(models, model)
	}
	return errno.StatusSuccessCode, errno.SuccessMsg, models
}
func AddAppointment(customerId int64, req *appointment.CreateAppointmentReq) (int64, string) {
	var err error
	app := mysql.Appointment{}
	app.CustomerId = customerId
	app.CatId = req.CatId
	app.Status = false
	app.StartTime, err = time.Parse(time.DateTime, req.StartTime)
	if err != nil {
		return errno.CreateErrorCode, err.Error()
	}
	app.EndTime, err = time.Parse(time.DateTime, req.EndTime)
	if err != nil {
		return errno.CreateErrorCode, err.Error()
	}

	app.ContactInfo = req.ContactInfo
	err = mysql.AddAppointment(&app)
	if err != nil {
		return errno.CreateErrorCode, err.Error()
	}
	return errno.StatusSuccessCode, errno.SuccessMsg
}

func UpdateApp(req *appointment.UpdateAppointmentReq) (int64, string) {
	var err error
	app := mysql.Appointment{}
	app.AppointmentId = req.CustomerId
	app.CatId = req.CatId
	app.CustomerId = req.CustomerId
	app.StartTime, err = time.Parse(time.DateTime, req.StartTime)
	if err != nil {
		return errno.UpdateErrorCode, err.Error()
	}
	app.EndTime, err = time.Parse(time.DateTime, req.EndTime)
	if err != nil {
		return errno.UpdateErrorCode, err.Error()
	}
	app.ContactInfo = req.ContactInfo
	app.Status = false
	err = mysql.UpdateApp(&app)
	if err != nil {
		return errno.UpdateErrorCode, err.Error()
	}
	return errno.StatusSuccessCode, errno.SuccessMsg
}

func ConfirmStatus(id int64, st int64) (int64, string) {
	f := false
	if st == 1 {
		f = true
	}
	err := mysql.ConfirmStatus(id, f)
	if err != nil {
		return errno.UpdateErrorCode, err.Error()
	}
	return errno.StatusSuccessCode, errno.SuccessMsg
}

func DelApp(appId int64, userId int64, isAdmin bool) (int64, string) {
	app, err := mysql.QueryAppByAppId(appId)
	if err != nil {
		return errno.DelErrorCode, err.Error()
	}

	if app.CustomerId != userId && !isAdmin {
		return errno.DelErrorCode, errno.PermissionFailedMsg
	}
	err = mysql.DelApp(appId)
	if err != nil {
		return errno.DelErrorCode, err.Error()
	}
	return errno.StatusSuccessCode, errno.SuccessMsg
}
