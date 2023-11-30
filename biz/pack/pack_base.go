package pack

import "github.com/Gorsonpy/catCafe/biz/model/membership"

func PackBase(resp *membership.BaseResponse, code int64, msg string) {
	resp.Code = code
	resp.Msg = msg
}
