package pack

import "github.com/Gorsonpy/catCafe/biz/model/membership"

func PackLoginResp(resp *membership.LoginResponse, code int64, msg string, token string) {
	resp.Base = membership.NewBaseResponse()
	resp.Base.Code = code
	resp.Base.Msg = msg
	resp.Data = make(map[string]string)
	resp.Data["token"] = token
}
