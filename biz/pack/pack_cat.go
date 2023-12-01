package pack

import "github.com/Gorsonpy/catCafe/biz/model/cat"

func PackAddCatResp(resp *cat.AddCatResp, code int64, msg string, catId int64) {
	resp.Base = cat.NewBaseResponse()
	resp.Base.Code = code
	resp.Base.Msg = msg
	resp.Data = make(map[string]int64)
	resp.Data["catId"] = catId
}

func PackQueryCatsResp(resp *cat.QueryCatsResp, code int64, msg string, list []*cat.CatModel) {
	resp.Base = cat.NewBaseResponse()
	resp.Base.Code = code
	resp.Base.Msg = msg
	resp.Data = make([]*cat.CatModel, 0)
	resp.Data = list
}
