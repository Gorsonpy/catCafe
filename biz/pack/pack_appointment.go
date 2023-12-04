package pack

import "github.com/Gorsonpy/catCafe/biz/model/appointment"

func PackAppBase(resp *appointment.BaseResponse, code int64, msg string) {
	resp.Code = code
	resp.Msg = msg
}

func PackQueryAppointmentResp(resp *appointment.QueryAppointmentResp, code int64, msg string, list []*appointment.AppointmentModel) {
	resp.Base = appointment.NewBaseResponse()
	resp.Base.Code = code
	resp.Base.Msg = msg
	resp.Data = list
}
