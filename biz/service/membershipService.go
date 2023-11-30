package userService

import (
	"github.com/Gorsonpy/catCafe/biz/dal/mysql"
	"github.com/Gorsonpy/catCafe/biz/model/membership"
	"github.com/Gorsonpy/catCafe/pkg/errno"
	"github.com/Gorsonpy/catCafe/pkg/utils"
)

func MembershipRegister(username string, passwd string) (int64, string) {
	if mysql.ExistUsername(username) {
		return errno.CreateError.ErrorCode, "用户名已存在"
	}
	hash, err := utils.HashPassword(passwd)
	if err != nil {
		return errno.CreateError.ErrorCode, "密码加密错误"
	}

	m := &membership.MembershipModel{}
	m.Username = username
	m.Passwd = hash

	err = mysql.AddMembership(m)
	if err != nil {
		return errno.CreateError.ErrorCode, "user create error"
	}
	return errno.SuccessCode, errno.SuccessMsg
}

func MembershipLogin(username string, passwd string) (int64, string, string) {
	if !mysql.ExistUsername(username) {
		return errno.UserExistedError.ErrorCode, "用户名不存在", ""
	}

	m := mysql.GetMembershipByUsername(username)
	if !utils.CheckPasswordHash(passwd, m.Passwd) {
		return errno.PWDError.ErrorCode, "密码错误", ""
	}
	token, err := utils.CreateToken(m.CustomerID)
	if err != nil {
		return -1, "token生成失败", ""
	}
	return errno.SuccessCode, errno.SuccessMsg, token
}
