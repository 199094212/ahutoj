package logic

import (
	"ahutoj/web/dao"
	"ahutoj/web/io/constanct"
	"ahutoj/web/io/request"
	"ahutoj/web/io/response"
	"ahutoj/web/models"
	"ahutoj/web/utils"

	"github.com/gin-gonic/gin"
)

func AddProblem(req *request.Problem, c *gin.Context) (interface{}, error) {
	problem := dao.Problem{
		Pid:           req.Pid,
		Title:         req.Title,
		Description:   req.Description,
		Input:         req.Input,
		Output:        req.Output,
		Sample_input:  req.Sample_input,
		Sample_output: req.Sample_output,
		Hit:           req.Hit,
		LimitTime:     req.LimitTime,
		LimitMemory:   req.LimitMemory,
	}
	//查看题目是否存在
	exist := models.IsProblemExistByPid(c, &problem)
	if exist {
		return response.CreateResponse(constanct.PIDExistCode), nil
	}
	//题目不存在 添加题目
	err := models.CreateProblem(c, &problem)
	if err != nil {
		//日志报错
		utils.GetLogInstance().Errorf("call CreateProblem failed,err=%s", err.Error())
		return response.CreateResponse(constanct.MySQLErrorCode), err
	}
	//成功返回
	return response.CreateResponse(constanct.SuccessCode), nil
}
