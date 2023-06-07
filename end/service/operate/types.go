package operate

import (
	"context"
	"test/service/databases"
	"test/service/myrpc/proto"
	"test/service/tools"
)

// 设备类型列表
func TypesGet(ctx context.Context, in *proto.TypesGetReq) (*proto.TypesGetResp, error) {
	userInfo := tools.GetJwtInfo(ctx)
	reps := new(proto.TypesGetResp)
	db := databases.GetDB().Model(databases.Typesinfo{}).Where("userid = ? AND del = 0", userInfo.ID)
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

// 增加设备类型
func TypesPost(ctx context.Context, in *proto.TypesPostReq) (*proto.TypesPostResp, error) {
	userInfo := tools.GetJwtInfo(ctx)
	err := databases.GetDB().Create(&databases.Typesinfo{
		Name:       in.Name,
		Details:    in.Details,
		Createtime: tools.GetTimes(),
		Del:        0,
		Userid:     userInfo.ID,
	}).Error
	if err != nil {
		return nil, err
	}
	return &proto.TypesPostResp{
		Message: "创建成功",
	}, err
}

// 删除设备类型
func TypesDel(ctx context.Context, in *proto.TypesDelReq) (*proto.TypesDelResp, error) {
	userInfo := tools.GetJwtInfo(ctx)
	err := databases.GetDB().Model(databases.Typesinfo{}).Where("id = ? AND userid = ? AND del = 0", in.Id, userInfo.ID).Update("del", 1).Error
	if err != nil {
		return nil, err
	}
	return &proto.TypesDelResp{Message: "删除成功"}, err
}

// 修改设备类型
func TypesPut(ctx context.Context, in *proto.TypesPutReq) (*proto.TypesPutResp, error) {
	userInfo := tools.GetJwtInfo(ctx)
	err := databases.GetDB().Model(databases.Typesinfo{}).Where("id = ? AND userid = ? AND del = 0", in.Id, userInfo.ID).Updates(map[string]interface{}{
		"name":    in.Name,
		"details": in.Details,
	}).Error
	if err != nil {
		return nil, err
	}
	return &proto.TypesPutResp{
		Message: "修改成功",
	}, err
}
