package service

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
	return errno.StatusSuccessCode, errno.SuccessMsg
}

func MembershipLogin(username string, passwd string) (int64, string, string, bool) {
	if !mysql.ExistUsername(username) {
		return errno.UserExistedError.ErrorCode, "用户名不存在", "", false
	}

	m := mysql.GetMembershipByUsername(username)
	if !utils.CheckPasswordHash(passwd, m.Passwd) {
		return errno.PWDError.ErrorCode, "密码错误", "", false
	}
	token, err := utils.CreateToken(m.CustomerID)
	if err != nil {
		return -1, "token生成失败", "", false
	}
	return errno.StatusSuccessCode, errno.SuccessMsg, token, mysql.IsAdmin(m.CustomerID)
}

func UpdatePoints(id int64, delta int64) (int64, string) {
	mem, err := mysql.GetMembershipById(id)
	if err != nil {
		return errno.UpdateErrorCode, err.Error()
	}
	mem.Points = mem.Points + delta
	err = mysql.UpdateMemPoint(id, mem.Points)
	if err != nil {
		return errno.UpdateErrorCode, err.Error()
	}
	return errno.StatusSuccessCode, errno.SuccessMsg
}

func ListMem() (int64, string, []*membership.MembershipModel) {
	mems, err := mysql.ListMem()
	if err != nil {
		return errno.GetErrorCode, err.Error(), nil
	}

	list := make([]*membership.MembershipModel, 0)
	for _, mem := range mems {
		tmp := mysql.MemToModel(mem)
		list = append(list, tmp)
	}
	return errno.StatusSuccessCode, errno.SuccessMsg, list
}
