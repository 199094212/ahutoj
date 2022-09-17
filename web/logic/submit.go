package logic

import (
	"ahutoj/web/dao"
	mysqldao "ahutoj/web/dao/mysqlDao"
	"ahutoj/web/io/constanct"
	"ahutoj/web/io/request"
	"ahutoj/web/io/response"
	"ahutoj/web/models"
	"ahutoj/web/utils"

	"github.com/gin-gonic/gin"
)

func AddSubmit(ctx *gin.Context, req *request.AddSubmitReq) (interface{}, error) {
	submit := dao.Submit{
		PID:        req.PID,
		CID:        req.CID,
		UID:        req.UID,
		Source:     req.Source,
		Lang:       req.Lang,
		Result:     constanct.OJ_JUDGE,
		SubmitTime: req.SubmitTime,
	}
	err := models.CreateSubmit(ctx, submit)
	if err != nil {
		return nil, err
	}
	return response.CreateResponse(constanct.SuccessCode), nil
}

func RejudgeSubmit(ctx *gin.Context, req *request.RejudgeSubmitReq) (interface{}, error) {
	submit := dao.Submit{}
	if req.SID != nil {
		submit.SID = *req.SID
	}
	if req.CID != nil {
		submit.CID = *req.CID
	}
	if req.PID != nil {
		submit.PID = *req.PID
	}
	if req.UID != nil {
		submit.UID = *req.UID
	}
	err := models.RejudgeSubmit(ctx, submit)
	if err != nil {
		return nil, err
	}
	return response.CreateResponse(constanct.SuccessCode), nil
}

func GetSubmits(ctx *gin.Context, req *request.SubmitListReq) (interface{}, error) {
	logger := utils.GetLogInstance()
	resp := response.SubmitListResp{}
	submit := dao.Submit{}
	if req.CID != nil {
		submit.CID = *req.CID
	}
	if req.PID != nil {
		submit.PID = *req.PID
	}
	if req.UID != nil {
		submit.UID = *req.UID
	}
	offset, limit := utils.GetPageInfo(req.Page, req.Limit)
	submits, err := models.GetSubmitList(ctx, submit, offset, limit)
	if err != nil {
		logger.Errorf("call SelectSubmitList failed,req=%+v,err=%s", utils.Sdump(req), err.Error())
		return nil, err
	}
	resp.Count, err = models.GetSubmitListCount(ctx, submit)
	if err != nil {
		logger.Errorf("call GetSubmitListCount failed,req=%+v,err=%s", utils.Sdump(req), err.Error())
		return nil, err
	}
	resp.Response = response.CreateResponse(constanct.SuccessCode)
	resp.Data = make([]response.SubmitLIstItem, len(submits))
	for i, temp := range submits {
		resp.Data[i] = response.SubmitLIstItem{
			SID:        temp.SID,
			PID:        temp.PID,
			Lang:       temp.Lang,
			Result:     temp.Result,
			UseTime:    temp.Usetime,
			UseMemory:  temp.Memory,
			SubmitTime: temp.SubmitTime,
		}
	}
	return resp, nil
}

func GetSubmit(ctx *gin.Context, req *request.GetSubmitReq) (interface{}, error) {
	logger := utils.GetLogInstance()
	submit, err := mysqldao.SelectSubmitBySID(ctx, req.SID)
	if err != nil {
		logger.Errorf("Call SelectSubmitBySID failed, SID=%v, err=%s", req.SID, err.Error())
		return response.CreateResponse(constanct.MySQLErrorCode), err
	}
	return response.GetSubmitResp{
		Response:   response.CreateResponse(constanct.SuccessCode),
		SID:        submit.SID,
		PID:        submit.PID,
		Source:     submit.Source,
		Lang:       submit.Lang,
		Result:     submit.Result,
		UseTime:    submit.Usetime,
		UseMemory:  submit.Memory,
		SubmitTime: submit.SubmitTime,
	}, nil
}
