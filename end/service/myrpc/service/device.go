package service

import (
	"context"
	"test/service/myrpc/proto"
	"test/service/operate"
	"test/service/tools"
)

// 设备注册
func (s Service) DeviceReg(ctx context.Context, in *proto.DeviceRegReq) (*proto.DeviceResp, error) {
	if in.Sn == "" {
		return nil, tools.ERROR(500, "SN不能为空")
	}
	if in.Name == "" {
		return nil, tools.ERROR(500, "设备名称不能为空")
	}
	if in.Typeid < 1 {
		return nil, tools.ERROR(500, "设备型号不能为空")
	}
	return operate.DeviceReg(ctx, in)
}

// 设备列表
func (s Service) DeviceGet(ctx context.Context, in *proto.DeviceGetReq) (*proto.DeviceGetResp, error) {
	if in.Limit > 999 {
		in.Limit = 10
	}
	if len(in.Fuzzy) > 30 {
		return nil, tools.ERROR(500, "参数错误")
	}
	return operate.DeviceGet(ctx, in)
}

// 设备修改
func (s Service) DevicePut(ctx context.Context, in *proto.DevicePutReq) (*proto.DevicePutResp, error) {
	if in.Id < 1 {
		return nil, tools.ERROR(500, "参数错误 ID错误")
	}
	if in.Name == "" {
		return nil, tools.ERROR(500, "设备名称不能为空")
	}
	if in.Typeid < 1 {
		return nil, tools.ERROR(500, "设备型号不能为空")
	}
	if in.Groupid < -1 || in.Groupid == 0 {
		return nil, tools.ERROR(500, "参数错误 分组错误")
	}
	if in.Typeid < 1 {
		return nil, tools.ERROR(500, "参数错误 设备型号错误")
	}
	return operate.DevicePut(ctx, in)
}
