package service

import (
	"context"
	"test/service/myrpc/proto"
	"test/service/operate"
	"test/service/tools"
)

// 设备类型列表
func (s Service) TypesGet(ctx context.Context, in *proto.TypesGetReq) (*proto.TypesGetResp, error) {
	if in.Limit > 999 {
		in.Limit = 10
	}
	if len(in.Fuzzy) > 30 {
		return nil, tools.ERROR(500, "模糊检索过长")
	}
	return operate.TypesGet(ctx, in)
}

// 增加设备类型
func (s Service) TypesPost(ctx context.Context, in *proto.TypesPostReq) (*proto.TypesPostResp, error) {
	if in.Name == "" {
		return nil, tools.ERROR(500, "名称不能为空")
	}
	if len(in.Name) > 100 {
		return nil, tools.ERROR(500, "名称过长")
	}
	if len(in.Details) > 1000 {
		in.Details = in.Details[:1000]
	}
	return operate.TypesPost(ctx, in)
}

// 删除设备类型
func (s Service) TypesDel(ctx context.Context, in *proto.TypesDelReq) (*proto.TypesDelResp, error) {
	if in.Id <= 0 {
		return &proto.TypesDelResp{
			Message: "删除成功",
		}, nil
	}
	return operate.TypesDel(ctx, in)
}

// 修改设备类型
func (s Service) TypesPut(ctx context.Context, in *proto.TypesPutReq) (*proto.TypesPutResp, error) {
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
		return &proto.TypesPutResp{
			Message: "修改成功",
		}, nil
	}
	return operate.TypesPut(ctx, in)
}
