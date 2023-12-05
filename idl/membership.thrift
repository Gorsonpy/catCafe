namespace go membership

struct BaseRequest{
}

struct MembershipModel{
    1: i64 customerID,
    2: string username,
    3: string contactInfo,
    4: i64 points,
    5: string level,
    6: string registrationDate,
    7: string passwd
    8: bool isAdmin
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

struct QueryMemResp{
    1: BaseResponse base
    2: list<MembershipModel> data
}

service MembershipService{
    QueryMemResp QueryMem(1:BaseRequest req)(api.get = "/membership")
    BaseResponse UpdatePoint(1:BaseRequest req)(api.put = "/membership")
    LoginResponse MembershipLogin(1: LoginRegisterReq req)(api.post = "/membership/login")
    BaseResponse MembershipRegister(1: LoginRegisterReq req)(api.post = "/membership/register")
}