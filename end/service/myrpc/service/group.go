package service

import (
	"context"
	"test/service/myrpc/proto"
	"test/service/operate"
	"test/service/tools"
)

// 设备组列表
func (s Service) GroupGet(ctx context.Context, in *proto.GroupGetReq) (*proto.GroupGetResp, error) {
	if in.Limit > 999 {
		in.Limit = 10
	}
	if len(in.Fuzzy) > 30 {
		return nil, tools.ERROR(500, "模糊检索过长")
	}
	return operate.GroupGet(ctx, in)
}

// 增加设备组
func (s Service) GroupPost(ctx context.Context, in *proto.GroupPostReq) (*proto.GroupPostResp, error) {
	if in.Name == "" {
		return nil, tools.ERROR(500, "名称不能为空")
	}
	if len(in.Name) > 100 {
		return nil, tools.ERROR(500, "名称过长")
	}
	if len(in.Details) > 1000 {
		in.Details = in.Details[:1000]
	}
	return operate.GroupPost(ctx, in)
}

// 删除设备组
func (s Service) GroupDel(ctx context.Context, in *proto.GroupDelReq) (*proto.GroupDelResp, error) {
	if in.Id <= 0 {
		return &proto.GroupDelResp{
			Message: "删除成功",
		}, nil
	}
	return operate.GroupDel(ctx, in)
}

// 修改设备组
func (s Service) GroupPut(ctx context.Context, in *proto.GroupPutReq) (*proto.GroupPutResp, error) {
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
		return &proto.GroupPutResp{
			Message: "修改成功",
		}, nil
	}
	return operate.GroupPut(ctx, in)
}
