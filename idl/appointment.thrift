namespace go appointment

struct BaseRequest{
}

struct BaseResponse{
    1: i64 code
    2: string msg
}

struct AppointmentModel{
    1: i64 appointmentId
    2: i64 customerId
    3: i64 catId
    4: string startTime
    5: string endTime
    6: string contactInfo
    7: bool status
}
struct CreateAppointmentReq{
    1: i64 catId
    2: string startTime
    3: string endTime
    4: string contactInfo
}

struct QueryAppointmentResp{
    1: BaseResponse base
    2: list<AppointmentModel> data
}

struct QueryAppointmentReq{
    1: string startTime
    2: string endTime
}

struct UpdateAppointmentReq{
    1: i64 appointmentId
    2: i64 customerId
    3: i64 catId
    4: string startTime
    5: string endTime
    6: string contactInfo 
}

service AppointmentService{
    BaseResponse CreateAppointment(1: CreateAppointmentReq req)(api.post = "/appointment")
    QueryAppointmentResp QueryAppointment(1:QueryAppointmentReq req)(api.get = "/appointment/current")
    BaseResponse UpdateAppointment(1: UpdateAppointmentReq req)(api.put = "/appointment")
    BaseResponse ConfirmStatus(1: BaseRequest req)(api.put = "/appointment/confirm")
    BaseResponse DelAppointment(1: BaseRequest req)(api.delete = "/appointment")
}