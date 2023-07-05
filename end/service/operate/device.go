package operate

import (
	"context"
	"test/service/databases"
	"test/service/myrpc/proto"
	"test/service/operate/cache"
	"test/service/tools"

	"gorm.io/gorm"
)

// 设备注册
func DeviceReg(ctx context.Context, in *proto.DeviceRegReq) (*proto.DeviceResp, error) {
	// 查询设备是否存在
	deviceInfo, err := cache.GetDeviceInfo(in.Sn)
	if err != nil {
		return nil, err
	}
	if deviceInfo.ID == 0 {
		//不存在 创建设备信息
		jwtInfo := tools.GetJwtInfo(ctx)
		deviceKey := tools.MD5(tools.RandString(6))
		deviceMD5 := tools.MD5(deviceKey)
		deviceInfo := new(databases.Device)
		timeNow := tools.GetTimes()

		// 不存在就创建设备
		deviceInfoNew := &databases.Device{
			Name:       in.Name,
			Sn:         in.Sn,
			Createtime: timeNow,
			Updatetime: timeNow,
			Userid:     jwtInfo.ID,
			Del:        0,
			Groupid:    -1,
			Details:    "",
			Typeid:     int(in.Typeid),
			Devicekey:  deviceKey,
			Devicemd5:  deviceMD5,
		}
		count := int64(0)
		databases.GetDB().Transaction(func(tx *gorm.DB) error {
			// 查询设备型号
			err := tx.Model(databases.Typesinfo{}).Where("id = ? AND userid = ? AND del = 0", in.Typeid, jwtInfo.ID).Count(&count).Error
			if err != nil {
				return err
			}
			if count != 0 {
				return tools.ERROR(500, "设备型号错误")
			}
			// 先查询设备是否存在
			err = tx.Where("sn = ? AND del = 0", in.Sn).Find(deviceInfo).Error
			if err != nil {
				deviceKey = deviceInfo.Devicekey
				return err
			}
			if deviceInfo.ID != 0 {
				return err
			}
			err = tx.Create(deviceInfoNew).Error
			return err
		})
		if err != nil {
			return nil, err
		}
		cache.DelUDeviceInfo(in.Sn)
		return &proto.DeviceResp{
			Data: deviceKey,
		}, err

	}

	return &proto.DeviceResp{
		Data: deviceInfo.Devicekey,
	}, err
}

// 设备列表
func DeviceGet(ctx context.Context, in *proto.DeviceGetReq) (*proto.DeviceGetResp, error) {
	userInfo := tools.GetJwtInfo(ctx)
	resp := new(proto.DeviceGetResp)
	db := databases.GetDB().Model(databases.Device{}).Where("userid = ? AND del = 0", userInfo.ID)
	if in.Typeid != 0 {
		db = db.Where("typeid = ?", in.Typeid)
	}
	if in.Groupid != 0 {
		db = db.Where("groupid = ?", in.Groupid)
	}
	if in.Fuzzy != "" {
		db = db.Where("sn like ? OR name like ? OR details like ?", "%"+in.Fuzzy+"%", "%"+in.Fuzzy+"%", "%"+in.Fuzzy+"%")
	}
	err := db.Count(&resp.Total).Error
	if err != nil {
		return nil, err
	}
	err = db.Limit(int(in.Limit)).Offset(int(in.Offset)).Scan(&resp.Data).Error
	if err != nil {
		return nil, err
	}
	return resp, err
}

// 设备修改
func DevicePut(ctx context.Context, in *proto.DevicePutReq) (*proto.DevicePutResp, error) {
	userInfo := tools.GetJwtInfo(ctx)
	err := databases.GetDB().Model(databases.Device{}).Where("id = ? AND userid = ?", in.Id, userInfo.ID).Updates(
		map[string]interface{}{
			"name":    in.Name,
			"details": in.Details,
			"typeid":  in.Typeid,
			"groupid": in.Groupid,
		}).Error
	if err != nil {
		return nil, err
	}
	return &proto.DevicePutResp{
		Message: "修改成功",
	}, nil
}
