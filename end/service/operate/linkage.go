package operate

import (
	"context"
	"test/service/databases"
	"test/service/myrpc/proto"
	"test/service/tools"
)

// 设备组列表
func LinkageGet(ctx context.Context, in *proto.LinkageGetReq) (*proto.LinkageGetResp, error) {
	userInfo := tools.GetJwtInfo(ctx)
	reps := new(proto.LinkageGetResp)
	db := databases.GetDB().Model(databases.Linkageinfo{}).Where("userid = ? AND del = 0", userInfo.ID)
	if in.Fuzzy != "" {
		db = db.Where("name like ? OR details like ?", "%"+in.Fuzzy+"%", "%"+in.Fuzzy+"%")
	}
	err := db.Count(&reps.Total).Error
	if err != nil {
		return nil, err
	}
	err = db.Limit(int(in.Limit)).Offset(int(in.Offset)).Scan(&reps.Data).Error
	if err != nil {
		return nil, err
	}
	return reps, err
}

// 增加设备组
func LinkagePost(ctx context.Context, in *proto.LinkagePostReq) (*proto.LinkagePostResp, error) {
	userInfo := tools.GetJwtInfo(ctx)
	err := databases.GetDB().Create(&databases.Linkageinfo{
		Name:       in.Name,
		Details:    in.Details,
		Createtime: tools.GetTimes(),
		Del:        0,
		Userid:     userInfo.ID,
	}).Error
	if err != nil {
		return nil, err
	}
	return &proto.LinkagePostResp{
		Message: "创建成功",
	}, err
}

// 删除设备组
func LinkageDel(ctx context.Context, in *proto.LinkageDelReq) (*proto.LinkageDelResp, error) {
	userInfo := tools.GetJwtInfo(ctx)
	err := databases.GetDB().Model(databases.Linkageinfo{}).Where("id = ? AND userid = ? AND del = 0", in.Id, userInfo.ID).Update("del", 1).Error
	if err != nil {
		return nil, err
	}
	return &proto.LinkageDelResp{Message: "删除成功"}, err
}

// 修改设备组
func LinkagePut(ctx context.Context, in *proto.LinkagePutReq) (*proto.LinkagePutResp, error) {
	userInfo := tools.GetJwtInfo(ctx)
	err := databases.GetDB().Model(databases.Linkageinfo{}).Where("id = ? AND userid = ? AND del = 0", in.Id, userInfo.ID).Updates(map[string]interface{}{
		"name":    in.Name,
		"details": in.Details,
	}).Error
	if err != nil {
		return nil, err
	}
	return &proto.LinkagePutResp{
		Message: "修改成功",
	}, err
}
