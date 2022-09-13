package mysqldao

import (
	"ahutoj/web/dao"
	"context"
)

func SelectSubmitList(ctx context.Context, submit dao.Submit, offset, limit int) (ans []dao.Submit, err error) {
	db := GetDB(ctx)
	err = db.Table(submit.TableName()).Where(&submit).Limit(limit).Offset(offset).Find(&ans).Error
	return ans, err
}

func SelectSubmitBySid(ctx context.Context, sid int64) (ans dao.Submit, err error) {
	db := GetDB(ctx)
	err = db.Table(dao.Submit{}.TableName()).Where("sid=?", sid).Find(&ans).Error
	return ans, err
}

func InsertSubmit(ctx context.Context, submit dao.Submit) (err error) {
	db := GetDB(ctx)
	err = db.Table(submit.TableName()).Create(&submit).Error
	return err
}

func RejudgeSubmit(ctx context.Context, submit dao.Submit) (err error) {
	db := GetDB(ctx)
	temp := dao.Submit{
		Result: "rejudge",
	}
	err = db.Table(temp.TableName()).Where(&submit).Updates(&temp).Error
	return err
}
func SelectCountSubmit(ctx context.Context, submit dao.Submit) (int64, error) {
	db := GetDB(ctx)
	var ans int64
	err = db.Table(submit.TableName()).Where(&submit).Count(&ans).Error
	return ans, err
}
