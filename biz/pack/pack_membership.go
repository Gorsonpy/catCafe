package pack

import "github.com/Gorsonpy/catCafe/biz/model/membership"

func PackLoginResp(resp *membership.LoginResponse, code int64, msg string, token string, isAdmin bool) {
	resp.Base = membership.NewBaseResponse()
	resp.Base.Code = code
	resp.Base.Msg = msg
	resp.Data = make(map[string]string)
	resp.Data["token"] = token

	s := ""
	if isAdmin {
		s = "1"
	} else {
		s = "0"
	}
	resp.Data["isAdmin"] = s
}

func PackQueryMem(resp *membership.QueryMemResp, code int64, msg string, list []*membership.MembershipModel) {
	resp.Base = membership.NewBaseResponse()
	resp.Base.Code = code
	resp.Base.Msg = msg
	resp.Data = make([]*membership.MembershipModel, 0)
	resp.Data = list
}
