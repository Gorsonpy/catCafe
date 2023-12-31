// Code generated by hertz generator.

package cat

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/Gorsonpy/catCafe/biz/dal/mysql"
	cat "github.com/Gorsonpy/catCafe/biz/model/cat"
	"github.com/Gorsonpy/catCafe/biz/model/membership"
	"github.com/Gorsonpy/catCafe/biz/pack"
	"github.com/Gorsonpy/catCafe/biz/service"
	"github.com/Gorsonpy/catCafe/pkg/errno"
	"github.com/Gorsonpy/catCafe/pkg/utils"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// UpdateCat .
// @router /cat [PUT]
func UpdateCat(ctx context.Context, c *app.RequestContext) {
	var err error
	var req cat.CatModel
	err = c.BindAndValidate(&req)

	resp := new(membership.BaseResponse)
	if err != nil {
		pack.PackBase(resp, errno.ParamErrorCode, errno.ParamErrorMsg)
		c.JSON(consts.StatusOK, resp)
		return
	}

	token_byte := c.GetHeader("token")
	claim, _ := utils.CheckToken(string(token_byte))
	if !mysql.IsAdmin(claim.UserId) {
		resp.Code = errno.AuthorizationFailedErrCode
		resp.Msg = errno.PermissionFailedMsg
		c.JSON(consts.StatusOK, resp)
		return
	}
	code, msg := service.UpdateCat(&req)
	pack.PackBase(resp, code, msg)
	c.JSON(consts.StatusOK, resp)
}

// QueryCats .
// @router /cat [GET]
func QueryCats(ctx context.Context, c *app.RequestContext) {
	var req cat.QueryCatsReq
	err := c.BindAndValidate(&req)
	resp := new(cat.QueryCatsResp)
	if err != nil {
		pack.PackQueryCatsResp(resp, errno.ParamErrorCode, errno.ParamErrorMsg, nil)
		return
	}

	code, msg, cats := service.QueryCats(&req)
	pack.PackQueryCatsResp(resp, code, msg, cats)
	c.JSON(consts.StatusOK, resp)
}

// AddCat .
// @router /cat [POST]
func AddCat(ctx context.Context, c *app.RequestContext) {
	var err error
	var req cat.CatModel
	err = c.BindAndValidate(&req)
	resp := new(cat.AddCatResp)

	token_byte := c.GetHeader("token")
	claim, _ := utils.CheckToken(string(token_byte))
	if !mysql.IsAdmin(claim.UserId) {
		pack.PackAddCatResp(resp, errno.AuthorizationFailedErrCode, errno.PermissionFailedMsg, -1)
		c.JSON(consts.StatusOK, resp)
		return
	}

	if err != nil {
		pack.PackAddCatResp(resp, errno.ParamErrorCode, err.Error(), -1)
		c.JSON(consts.StatusOK, resp)
		return
	}

	code, msg, catId := service.AddCat(&req)
	pack.PackAddCatResp(resp, code, msg, catId)
	c.JSON(consts.StatusOK, resp)
}

// DelCat .
// @router /cat [DELETE]
func DelCat(ctx context.Context, c *app.RequestContext) {
	resp := new(cat.BaseResponse)
	var req cat.DelCatsReq
	err := c.BindAndValidate(&req)
	if err != nil {
		resp.Code = errno.ParamErrorCode
		resp.Msg = err.Error()
		c.JSON(consts.StatusOK, resp)
		return
	}

	token_byte := c.GetHeader("token")
	claim, _ := utils.CheckToken(string(token_byte))
	if !mysql.IsAdmin(claim.UserId) {
		resp.Code = errno.AuthorizationFailedErrCode
		resp.Msg = errno.PermissionFailedMsg
		c.JSON(consts.StatusOK, resp)
		return
	}

	if err != nil {
		resp.Code = errno.ParamErrorCode
		resp.Msg = errno.ParamErrorMsg
		c.JSON(consts.StatusOK, resp)
		return
	}
	code, msg := service.DelCat(req.CatIds)
	resp.Code = code
	resp.Msg = msg
	c.JSON(consts.StatusOK, resp)
}

// QueryCatsByPop .
// @router /cat/limit [GET]
func QueryCatsByPop(ctx context.Context, c *app.RequestContext) {
	resp := new(cat.QueryCatsResp)
	limit, err := strconv.Atoi(c.Query("limit"))

	if err != nil {
		pack.PackQueryCatsResp(resp, errno.AuthorizationFailedErrCode, errno.UnLoginFailedMsg, nil)
		c.JSON(consts.StatusOK, resp)
		return
	}

	code, msg, cats := service.QueryTopCats(int64(limit))
	pack.PackQueryCatsResp(resp, code, msg, cats)
	c.JSON(consts.StatusOK, resp)
}

// UploadFile .
// @router /file [POST]
func UploadFile(ctx context.Context, c *app.RequestContext) {
	resp := new(cat.UploadResp)
	file, err := c.FormFile("file")
	if err != nil {
		pack.PackUploadResp(resp, errno.CreateErrorCode, err.Error(), "")
		c.JSON(consts.StatusOK, resp)
		return
	}

	// 设置阿里云 OSS 相关信息
	endpoint := "oss-cn-beijing.aliyuncs.com"
	accessKeyID := "LTAI5t974iHp9XRqvD1KsgDD"
	accessKeySecret := "7OCxGLqoSJEDYvulMtCMTTrGLEKFfV"
	bucketName := "gorsonpy-bucket"

	// 创建 OSS 客户端
	client, err := oss.New(endpoint, accessKeyID, accessKeySecret)
	if err != nil {
		pack.PackUploadResp(resp, errno.CreateErrorCode, err.Error(), "")
		c.JSON(consts.StatusOK, resp)
		return
	}

	// 获取存储空间（Bucket）
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		pack.PackUploadResp(resp, errno.CreateErrorCode, err.Error(), "")
		c.JSON(consts.StatusOK, resp)
		return
	}

	// 生成唯一文件名，使用时间戳
	timestamp := time.Now().UnixNano() / int64(time.Millisecond)
	objectKey := fmt.Sprintf("uploads/%d_%s", timestamp, file.Filename)

	f, err := file.Open()

	// 上传文件到 OSS
	err = bucket.PutObject(objectKey, f)
	if err != nil {
		pack.PackUploadResp(resp, errno.CreateErrorCode, err.Error(), "")
		c.JSON(consts.StatusOK, resp)
		return
	}

	// 构建上传成功后的 OSS URL
	ossURL := fmt.Sprintf("https://%s.%s/%s", bucketName, endpoint, objectKey)
	pack.PackUploadResp(resp, errno.StatusSuccessCode, errno.SuccessMsg, ossURL)
	c.JSON(consts.StatusOK, resp)
}
