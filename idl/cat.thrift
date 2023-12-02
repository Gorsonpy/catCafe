namespace go cat


struct BaseRequest{
}

struct BaseResponse{
    1: i64 code,
    2: string msg
}

struct CatModel{
    1: i64 catId,
    2: string name,
    3: string breed,
    4: string gender,
    5: i64 age,
    6: string healthStatus,
    7: string photoUrl,
    8: string checkInDate,
    9: i64 appointmentNum,
}

struct AddCatResp{
    1: BaseResponse base,
    2: map<string, i64> data
}

struct QueryCatsReq{
    1: string name
    2: string breed
    3: string gender
    4: string lAge
    5: string rAge
    6: string healthStatus
    7: i64 limit
}
struct QueryCatsResp{
    1: BaseResponse base
    2: list<CatModel> data
}

service CatService{
    QueryCatsResp queryCatsByPop(1:BaseRequest req)(api.get = "/cat/limit")
    BaseResponse updateCat(1:CatModel req)(api.put = "/cat")
    QueryCatsResp queryCats(1:QueryCatsReq req)(api.get = "/cat") 
    AddCatResp addCat(1:CatModel req)(api.post = "/cat")
    BaseResponse delCat(1:BaseRequest req)(api.delete = "/cat")
}
