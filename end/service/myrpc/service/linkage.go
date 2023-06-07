package service

import (
	"context"
	"test/service/myrpc/proto"
	"test/service/operate"
	"test/service/tools"
)

// 设备组列表
func (s Service) LinkageGet(ctx context.Context, in *proto.LinkageGetReq) (*proto.LinkageGetResp, error) {
	if in.Limit > 999 {
		in.Limit = 10
	}
	if len(in.Fuzzy) > 30 {
		return nil, tools.ERROR(500, "模糊检索过长")
	}
	return operate.LinkageGet(ctx, in)
}

// 增加设备组
func (s Service) LinkagePost(ctx context.Context, in *proto.LinkagePostReq) (*proto.LinkagePostResp, error) {
	if in.Name == "" {
		return nil, tools.ERROR(500, "名称不能为空")
	}
	if len(in.Name) > 100 {
		return nil, tools.ERROR(500, "名称过长")
	}
	if len(in.Details) > 1000 {
		in.Details = in.Details[:1000]
	}
	return operate.LinkagePost(ctx, in)
}

// 删除设备组
func (s Service) LinkageDel(ctx context.Context, in *proto.LinkageDelReq) (*proto.LinkageDelResp, error) {
	if in.Id <= 0 {
		return &proto.LinkageDelResp{
			Message: "删除成功",
		}, nil
	}
	return operate.LinkageDel(ctx, in)
}

// 修改设备组
func (s Service) LinkagePut(ctx context.Context, in *proto.LinkagePutReq) (*proto.LinkagePutResp, error) {
	if in.Name == "" {
		return nil, tools.ERROR(500, "名称不能为空")
	}
	if len(in.Name) > 100 {
		return nil, tools.ERROR(500, "名称过长")
	}
	if len(in.Details) > 1000 {
		in.Details = in.Details[:1000]
	}

	if in.Id <= 0 {
		return &proto.LinkagePutResp{
			Message: "修改成功",
		}, nil
	}
	return operate.LinkagePut(ctx, in)
}
