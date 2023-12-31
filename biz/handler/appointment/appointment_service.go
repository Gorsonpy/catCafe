// Code generated by hertz generator.

package appointment

import (
	"context"
	"strconv"

	"github.com/Gorsonpy/catCafe/biz/dal/mysql"
	appointment "github.com/Gorsonpy/catCafe/biz/model/appointment"
	"github.com/Gorsonpy/catCafe/biz/pack"
	"github.com/Gorsonpy/catCafe/biz/service"
	"github.com/Gorsonpy/catCafe/pkg/errno"
	"github.com/Gorsonpy/catCafe/pkg/utils"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// CreateAppointment .
// @router /appointment [POST]
func CreateAppointment(ctx context.Context, c *app.RequestContext) {
	var err error
	var req appointment.CreateAppointmentReq
	err = c.BindAndValidate(&req)

	resp := new(appointment.BaseResponse)
	if err != nil {
		pack.PackAppBase(resp, errno.PwdErrorCode, errno.ParamErrorMsg)
		c.JSON(consts.StatusOK, resp)
		return
	}

	token_byte := c.GetHeader("token")
	claim, _ := utils.CheckToken(string(token_byte))
	code, msg := service.AddAppointment(claim.UserId, &req)
	pack.PackAppBase(resp, code, msg)
	c.JSON(consts.StatusOK, resp)
}

// QueryAppointment .
// @router /appointment [GET]
func QueryAppointment(ctx context.Context, c *app.RequestContext) {
	var err error
	var req appointment.QueryAppointmentReq
	err = c.BindAndValidate(&req)

	resp := new(appointment.QueryAppointmentResp)
	if err != nil {
		pack.PackQueryAppointmentResp(resp, errno.PwdErrorCode, errno.ParamErrorMsg, nil)
		c.JSON(consts.StatusOK, resp)
		return
	}

	token_byte := c.GetHeader("token")
	claim, _ := utils.CheckToken(string(token_byte))
	code, msg, list := service.QueryApp(claim.UserId, &req)

	pack.PackQueryAppointmentResp(resp, code, msg, list)
	c.JSON(consts.StatusOK, resp)
}

// UpdateAppointment .
// @router /appointment [PUT]
func UpdateAppointment(ctx context.Context, c *app.RequestContext) {
	var err error
	var req appointment.UpdateAppointmentReq
	err = c.BindAndValidate(&req)
	resp := new(appointment.BaseResponse)

	if err != nil {
		pack.PackAppBase(resp, errno.ParamErrorCode, errno.ParamErrorMsg)
		c.JSON(consts.StatusOK, resp)
		return
	}

	token_byte := c.GetHeader("token")
	claim, _ := utils.CheckToken(string(token_byte))
	req.CustomerId = claim.UserId

	code, msg := service.UpdateApp(&req)
	pack.PackAppBase(resp, code, msg)
	c.JSON(consts.StatusOK, resp)
}

// ConfirmStatus .
// @router /appointment/confirm [PUT]
func ConfirmStatus(ctx context.Context, c *app.RequestContext) {
	resp := new(appointment.BaseResponse)
	id, err := strconv.Atoi(c.Query("appointmentId"))
	if err != nil {
		pack.PackAppBase(resp, errno.ParamErrorCode, err.Error())
	}
	st, err := strconv.Atoi(c.Query("status"))
	if err != nil {
		pack.PackAppBase(resp, errno.ParamErrorCode, err.Error())
	}

	token_byte := c.GetHeader("token")
	claim, _ := utils.CheckToken(string(token_byte))
	if !mysql.IsAdmin(claim.UserId) {
		pack.PackAppBase(resp, errno.AuthorizationFailedErrCode, errno.PermissionFailedMsg)
		c.JSON(consts.StatusOK, resp)
		return
	}
	code, msg := service.ConfirmStatus(int64(id), int64(st))
	pack.PackAppBase(resp, code, msg)
	c.JSON(consts.StatusOK, resp)
}

// DelAppointment .
// @router /appointment [DELETE]
func DelAppointment(ctx context.Context, c *app.RequestContext) {
	resp := new(appointment.BaseResponse)
	id, err := strconv.Atoi(c.Query("appointmentId"))
	if err != nil {
		pack.PackAppBase(resp, errno.ParamErrorCode, err.Error())
		c.JSON(consts.StatusOK, resp)
		return
	}
	token_byte := c.GetHeader("token")
	claim, _ := utils.CheckToken(string(token_byte))
	code, msg := service.DelApp(int64(id), claim.UserId, mysql.IsAdmin(claim.UserId))
	pack.PackAppBase(resp, code, msg)
	c.JSON(consts.StatusOK, resp)
}
