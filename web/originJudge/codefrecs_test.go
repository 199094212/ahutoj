package originjudge_test

import (
	"ahutoj/web/dao"
	mysqldao "ahutoj/web/dao/mysqlDao"
	"ahutoj/web/io/constanct"
	originjudge "ahutoj/web/originJudge"
	"ahutoj/web/utils"
	"context"
	"testing"
	"time"
)

func TestSubmitAndGetResult(t *testing.T) {
	utils.ConfigInit("../../config.yaml")
	cfJudge := originjudge.CodeForceJudge{}
	cfJudge.Submit = dao.Submit{
		Lang:       constanct.CPP11,
		SubmitTime: time.Now().UnixMilli(),
		Source: `#include<iostream>
		using namespace std;
		int main(){
			//111222..q
			int a,b;
			cin>>a>>b;
			cout<<a+b<<endl;
		return 0;
		}`,
		Result: constanct.OJ_JUDGE,
	}
	cfJudge.PID = "103446I"
	cfJudge.Judge(context.Background(), cfJudge.Submit, cfJudge.PID)
}

func TestAtcoderLogin(t *testing.T) {
	utils.ConfigInit("../../config.yaml")
	utils.GetConfInstance().MySQLConfig.Host = "116.205.190.37"
	mysqldao.InitMysql()
	utils.LogInit()
	submit, _ := mysqldao.SelectSubmitBySID(context.Background(), 1024)
	submit.Result = constanct.OJ_JUDGE
	for {
		for i := 0; i < 5; i++ {
			originJudge := originjudge.GetOriginJudgeFunc(originjudge.OJPlatform(submit.OJPlatform))
			originJudge.Judge(context.Background(), submit, "abc272_c")
		}
		time.Sleep(20 * time.Second)
	}

}
