package operate

import (
	"context"
	"test/service/databases"
	"test/service/myrpc/proto"
	"test/service/operate/cache"
	"test/service/tools"

	"github.com/gomodule/redigo/redis"
	"gorm.io/gorm"
)

// 登录
func LoginWeb(ctx context.Context, in *proto.LoginWebReq) (*proto.Loginresp, error) {
	// 查找账号
	userInfo := new(databases.Userinfo)
	err := databases.GetDB().Where("emails = ? AND del = 0", in.Emails).Find(userInfo).Limit(1).Error
	if err != nil {
		return nil, err
	}
	if userInfo.ID == 0 {
		return nil, tools.ERROR(500, "账号密码错误")
	}

	// 校验密码
	if in.Password != tools.MD5(userInfo.Password+in.Time) {
		return nil, tools.ERROR(500, "账号密码错误")
	}

	// 查看账号状态
	if userInfo.State == 1 {
		return nil, tools.ERROR(500, "账号未激活,请查看激活邮件")
	} else if userInfo.State == 3 {
		return nil, tools.ERROR(500, "账号已被禁用,请联系管理员")
	}

	// 生成jwt
	token, err := tools.NewJWT(tools.Jwtinfo{
		ID:         userInfo.ID,
		Emails:     userInfo.Emails,
		Nicename:   userInfo.Nicename,
		Updatetime: userInfo.Updatetime,
	})
	if err != nil {
		return nil, err
	}
	// 保存到redis
	cache.SetRedisUserInfo(userInfo)
	return &proto.Loginresp{
		Code:    0,
		Message: "",
		Data: &proto.LoginInfo{
			Nicename: userInfo.Nicename,
			Emails:   userInfo.Emails,
			Token:    "Bearer " + token,
		},
	}, nil
}

// 修改密码(登陆前)
func LoginPwd(ctx context.Context, in *proto.LoginPwdReq) (*proto.Loginresp, error) {
	// 验证code
	code, err := cache.GetUserCode(in.Emails, "updatepassword")
	if err != nil {
		if err == redis.ErrNil {
			return nil, tools.ERROR(500, "验证码错误")
		}
		return nil, err
	}
	if in.Code != code {
		return nil, tools.ERROR(500, "验证码错误")
	}
	userInfo, err := cache.GetUserInfo(in.Emails)
	if err != nil {
		return nil, err
	}
	if userInfo.ID == 0 {
		return nil, tools.ERROR(500, "账户错误")
	}
	if userInfo.State != 2 {
		return nil, tools.ERROR(500, "账户未激活或已禁用")
	}
	err = databases.GetDB().Model(userInfo).Updates(map[string]interface{}{
		"password":   in.Password,
		"updatetime": tools.GetTimes(),
	}).Error
	if err != nil {
		return nil, err
	}
	cache.DelUserCode(in.Emails, "updatepassword")
	cache.DelUserInfo(in.Emails)
	return &proto.Loginresp{
		Message: "修改成功",
	}, err
}

// 修改密码(登陆后)
func LoginPwdLogin(ctx context.Context, in *proto.LoginPwdLoginReq) (*proto.Loginresp, error) {
	userInfo := tools.GetJwtInfo(ctx)
	// 校验密码
	if in.Password != tools.MD5(userInfo.Password+in.Time) {
		return nil, tools.ERROR(500, "密码错误")
	}
	err := databases.GetDB().Model(userInfo).Updates(map[string]interface{}{
		"password":   in.NewPassword,
		"updatetime": tools.GetTimes(),
	}).Error
	if err != nil {
		return nil, err
	}
	cache.DelUserInfo(userInfo.Emails)
	return &proto.Loginresp{
		Message: "修改成功",
	}, err
}

// 修改邮箱/昵称
func LoginInfo(ctx context.Context, in *proto.LoginInfoReq) (*proto.Loginresp, error) {
	return nil, nil
}

// 创建账号
func LoginUser(ctx context.Context, in *proto.LoginUserReq) (*proto.Loginresp, error) {
	// 验证code
	code, err := cache.GetUserCode(in.Emails, "create")
	if err != nil {
		if err == redis.ErrNil {
			return nil, tools.ERROR(500, "验证码错误")
		}
		return nil, err
	}
	if in.Code != code {
		return nil, tools.ERROR(500, "验证码错误")
	}

	userInfo := new(databases.Userinfo)
	err = databases.GetDB().Transaction(func(tx *gorm.DB) error {
		err := tx.Where("emails = ? AND del = 0", in.Emails).Find(userInfo).Limit(1).Error
		if err != nil {
			return err
		}
		if userInfo.ID != 0 {
			return tools.ERROR(500, "账号已存在")
		}
		err = tx.Create(&databases.Userinfo{
			Nicename:   in.NickName,
			Emails:     in.Emails,
			Password:   in.Password,
			Createtime: tools.GetTimes(),
			Updatetime: tools.GetTimes(),
			Del:        0,
			State:      2,
		}).Error
		if err != nil {
			return err
		}
		cache.DelUserCode(in.Emails, "create")
		return err
	})
	if err != nil {
		return nil, err
	}
	return &proto.Loginresp{
		Message: "创建成功",
	}, err
}

// 发送创建账户验证码
func LoginUserCode(ctx context.Context, in *proto.LoginUserCodeReq) (*proto.Loginresp, error) {
	// 查询账户是否存在
	count := int64(0)
	err := databases.GetDB().Model(databases.Userinfo{}).Where("emails = ? AND del = 0", in.Emails).Count(&count).Error
	if err != nil {
		return nil, err
	}
	if count != 0 {
		if in.Types == "create" {
			return nil, tools.ERROR(500, "账号已存在")
		}
	}
	// 检测是否存在code
	_, err = cache.GetUserCode(in.Emails, in.Types)
	if err == nil {
		return nil, tools.ERROR(500, "已发送,请稍后再试")
	}
	// 缓存code
	code := tools.RandString(6)
	err = cache.SetUserCode(in.Emails, in.Types, code)
	if err != nil {
		return nil, err
	}
	// 发送邮件
	err = tools.SetEmail([]string{in.Emails}, in.Name, code)
	if err != nil {
		return nil, err
	}
	return &proto.Loginresp{
		Message: "验证码发送成功",
		Data: &proto.LoginInfo{
			Token: code,
		},
	}, err
}
