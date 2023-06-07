package service

import (
	"context"
	"strconv"
	"test/service/myrpc/proto"
	"test/service/operate"
	"test/service/tools"
)

// 登录
func (s Service) LoginWeb(ctx context.Context, in *proto.LoginWebReq) (*proto.Loginresp, error) {
	// 初步校验账户密码
	if len(in.Password) != 32 || len(in.Time) != 13 {
		return nil, tools.ERROR(500, "账号密码错误")
	}
	if !tools.VerifyEmailFormat(in.Emails) {
		return nil, tools.ERROR(500, "账号密码错误")
	}

	// 判断是否超时
	times, err := strconv.ParseInt(in.Time, 10, 64)
	if err != nil {
		return nil, err
	}
	now := tools.GetTimes()
	if times+300*1000 < now {
		return nil, tools.ERROR(500, "登录已过期")
	}

	return operate.LoginWeb(ctx, in)
}

// 修改密码(登陆前)
func (s Service) LoginPwd(ctx context.Context, in *proto.LoginPwdReq) (*proto.Loginresp, error) {
	if !tools.VerifyEmailFormat(in.Emails) {
		return nil, tools.ERROR(500, "邮箱错误")
	}
	if len(in.Code) != 6 {
		return nil, tools.ERROR(500, "验证码错误")
	}
	if len(in.Password) != 32 {
		return nil, tools.ERROR(500, "修改失败")
	}
	return operate.LoginPwd(ctx, in)
}

// 修改密码(登陆后)
func (s Service) LoginPwdLogin(ctx context.Context, in *proto.LoginPwdLoginReq) (*proto.Loginresp, error) {
	if len(in.Password) != 32 || len(in.NewPassword) != 32 || len(in.Time) != 13 {
		return nil, tools.ERROR(500, "修改失败")
	}
	// 判断是否超时
	times, err := strconv.ParseInt(in.Time, 10, 64)
	if err != nil {
		return nil, err
	}
	now := tools.GetTimes()
	if times+300*1000 < now {
		return nil, tools.ERROR(500, "修改密码超时")
	}
	return operate.LoginPwdLogin(ctx, in)
}

// 修改邮箱/昵称
func (s Service) LoginInfo(ctx context.Context, in *proto.LoginInfoReq) (*proto.Loginresp, error) {
	if in.NickName == "" {
		return &proto.Loginresp{
			Message: "修改成功",
		}, nil
	}
	return operate.LoginInfo(ctx, in)
}

// 创建账号
func (s Service) LoginUser(ctx context.Context, in *proto.LoginUserReq) (*proto.Loginresp, error) {
	if len(in.Password) != 32 {
		return nil, tools.ERROR(500, "申请失败,密码有误")
	}
	if !tools.VerifyEmailFormat(in.Emails) {
		return nil, tools.ERROR(500, "邮箱错误")
	}
	if len(in.Code) != 6 {
		return nil, tools.ERROR(500, "验证码错误")
	}
	if in.NickName == "" {
		return nil, tools.ERROR(500, "昵称不能为空")
	}
	return operate.LoginUser(ctx, in)
}

// 发送创建账户验证码
func (s Service) LoginUserCode(ctx context.Context, in *proto.LoginUserCodeReq) (*proto.Loginresp, error) {
	if !tools.VerifyEmailFormat(in.Emails) {
		return nil, tools.ERROR(500, "邮箱错误")
	}
	if in.Types == "" || in.Name == "" {
		return nil, tools.ERROR(500, "发送验证码失败")
	}
	return operate.LoginUserCode(ctx, in)
}
