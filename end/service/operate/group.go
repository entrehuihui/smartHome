package operate

import (
	"context"
	"test/service/databases"
	"test/service/myrpc/proto"
	"test/service/tools"

	"gorm.io/gorm"
)

// 设备组列表
func GroupGet(ctx context.Context, in *proto.GroupGetReq) (*proto.GroupGetResp, error) {
	userInfo := tools.GetJwtInfo(ctx)
	reps := new(proto.GroupGetResp)
	db := databases.GetDB().Model(databases.Groupinfo{}).Where("userid = ? AND del = 0", userInfo.ID)
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
func GroupPost(ctx context.Context, in *proto.GroupPostReq) (*proto.GroupPostResp, error) {
	userInfo := tools.GetJwtInfo(ctx)
	err := databases.GetDB().Create(&databases.Groupinfo{
		Name:       in.Name,
		Details:    in.Details,
		Createtime: tools.GetTimes(),
		Del:        0,
		Userid:     userInfo.ID,
	}).Error
	if err != nil {
		return nil, err
	}
	return &proto.GroupPostResp{
		Message: "创建成功",
	}, err
}

// 删除设备组
func GroupDel(ctx context.Context, in *proto.GroupDelReq) (*proto.GroupDelResp, error) {
	userInfo := tools.GetJwtInfo(ctx)
	// 先检测有没有设备 然后再删除
	count := int64(0)
	err := databases.GetDB().Transaction(func(tx *gorm.DB) error {
		err := tx.Model(databases.Device{}).Where("groupid = ? AND del = 0", in.Id).Count(&count).Error
		if err != nil {
			return err
		}
		if count != 0 {
			return tools.ERROR(500, "设备组无法删除")
		}
		err = tx.Model(databases.Groupinfo{}).Where("id = ? AND userid = ? AND del = 0", in.Id, userInfo.ID).Update("del", 1).Error
		if err != nil {
			return err
		}
		return err
	})
	if err != nil {
		return nil, err
	}

	return &proto.GroupDelResp{Message: "删除成功"}, err
}

// 修改设备组
func GroupPut(ctx context.Context, in *proto.GroupPutReq) (*proto.GroupPutResp, error) {
	userInfo := tools.GetJwtInfo(ctx)
	err := databases.GetDB().Transaction(func(tx *gorm.DB) error {
		// 修改设备
		err := tx.Model(databases.Device{}).Where("groupid = ?", in.Id).Update("groupid", -1).Error
		if err != nil {
			return err
		}
		err = tx.Model(databases.Groupinfo{}).Where("id = ? AND userid = ? AND del = 0", in.Id, userInfo.ID).Updates(map[string]interface{}{
			"name":    in.Name,
			"details": in.Details,
		}).Error
		if err != nil {
			return err
		}
		return err
	})
	if err != nil {
		return nil, err
	}
	return &proto.GroupPutResp{
		Message: "修改成功",
	}, err
}

// 设备组添加设备
func GroupDevicePost(ctx context.Context, in *proto.GroupDevicePostReq) (*proto.GroupDevicePostResp, error) {
	// 先校验全部设备是否属于该用户
	userInfo := tools.GetJwtInfo(ctx)
	count := int64(0)
	err := databases.GetDB().Model(databases.Device{}).Where("sn = ? AND userid = ?", in.Sn, userInfo.ID).Count(&count).Error
	if err != nil {
		return nil, err
	}
	if len(in.Sn) != int(count) {
		return nil, tools.ERROR(500, "SN错误或设备不存在")
	}
	// 修改设备组
	err = databases.GetDB().Model(databases.Device{}).Where("sn = ? AND userid = ?", in.Sn, userInfo.ID).Update("groupid", in.GroupID).Error
	if err != nil {
		return nil, err
	}
	return &proto.GroupDevicePostResp{
		Message: "成功",
	}, nil
}

// 设备组删除设备
func GroupDevicePut(ctx context.Context, in *proto.GroupDevicePutReq) (*proto.GroupDevicePutResp, error) {
	err := databases.GetDB().Model(databases.Device{}).Where("sn = ? AND groupid = ?", in.Sn, in.GroupID).Update("groupid", -1).Error
	if err != nil {
		return nil, err
	}
	return &proto.GroupDevicePutResp{
		Message: "删除成功",
	}, nil
}
