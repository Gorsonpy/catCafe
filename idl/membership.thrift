namespace go membership

struct BaseRequest{
}

struct MembershipModel{
    1: i64 customerID,
    2: string username,
    3: string contactInfo,
    4: string points,
    5: string level,
    6: string registrationDate,
    7: string passwd
}
struct LoginRegisterReq{
    1: string username,
    2: string passwd
}
struct BaseResponse{
    1: i64 code,
    2: string msg
}
struct LoginResponse{
    1: BaseResponse base,
    2: map<string, string> data
}
service MembershipService{
    LoginResponse membershipLogin(1: LoginRegisterReq req)(api.post = "/membership/login")
    BaseResponse membershipRegister(1: LoginRegisterReq req)(api.post = "/membership/register")
}