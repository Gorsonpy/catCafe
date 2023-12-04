package mysql

import (
	"time"

	"github.com/Gorsonpy/catCafe/biz/model/appointment"
)

type Appointment struct {
	AppointmentId int64 `gorm:"primary_key"`
	CustomerId    int64
	CatId         int64
	ContactInfo   string
	Status        bool
	StartTime     time.Time
	EndTime       time.Time
}

func AppToModel(app *Appointment) *appointment.AppointmentModel {
	tmp := appointment.NewAppointmentModel()
	tmp.AppointmentId = app.AppointmentId
	tmp.CustomerId = app.CustomerId
	tmp.CatId = app.CatId
	tmp.ContactInfo = app.ContactInfo
	tmp.Status = app.Status
	tmp.StartTime = app.StartTime.Format(time.DateTime)
	tmp.EndTime = app.EndTime.Format(time.DateTime)
	return tmp
}

func UpdateApp(app *Appointment) error {
	return DB.Table("appointments").Updates(&app).Error
}

func AddAppointment(app *Appointment) error {
	return DB.Table("appointments").Create(&app).Error
}

func QueryApp(customerId int64, start string, end string) ([]*Appointment, error) {
	list := make([]*Appointment, 0)
	err := DB.Table("appointments").Where("customer_id = ? and start_time > ? and end_time < ?", customerId, start, end).Find(&list).Error
	return list, err
}

func ConfirmStatus(id int64, f bool) error {
	return DB.Table("appointments").Where("appointment_id = ?", id).Update("status", f).Error
}

func DelApp(id int64) error {
	return DB.Table("appointments").Where("appointment_id = ?", id).Delete(&Appointment{}).Error
}

func QueryAppByAppId(id int64) (*Appointment, error) {
	app := Appointment{}
	err := DB.Table("appointments").Where("appointment_id = ?", id).First(&app).Error
	return &app, err
}
