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

struct DelCatsReq{
    1: list<i64> catIds
}

struct QueryCatsReq{
    1: string searchContent
    2: string breed
    3: string gender
    4: i64 lAge
    5: i64 rAge
    6: i64 limit
}
struct QueryCatsResp{
    1: BaseResponse base
    2: list<CatModel> data
}

struct UploadResp{
    1: BaseResponse base
    2: map<string, string> data
}

service CatService{
    UploadResp UploadFile(1:BaseRequest req)(api.post = "/file")
    QueryCatsResp queryCatsByPop(1:BaseRequest req)(api.get = "/cat/limit")
    BaseResponse updateCat(1:CatModel req)(api.put = "/cat")
    QueryCatsResp queryCats(1:QueryCatsReq req)(api.post = "/cat/search") 
    AddCatResp addCat(1:CatModel req)(api.post = "/cat")
    BaseResponse delCat(1:DelCatsReq req)(api.delete = "/cat")
}
