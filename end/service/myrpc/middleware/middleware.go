package middleware

import (
	"context"
	"fmt"
	"log"
	"test/service/operate/cache"
	"test/service/tools"

	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/auth"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// ZapLogger .
func ZapLogger() *zap.Logger {
	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatalf("failed to initialize zap logger: %v", err)
	}
	grpc_zap.ReplaceGrpcLogger(logger)
	return logger
}

// MyAuthFunction .
func MyAuthFunction(ctx context.Context) (context.Context, error) {

	method, ok := grpc.Method(ctx)
	if !ok {
		return nil, status.Errorf(301, "请求错误, 无法获取请求数据")
	}
	log.Println("method =====>>>> ", method)
	// // 过滤全部LoginService服务 /proto.SmsService/SmsPost
	if method == "/proto.Login/loginWeb" || method == "/proto.Login/loginWeChat" || method == "/proto.Login/loginUserCode" || method == "/proto.Login/loginUser" || method == "/proto.Login/loginPwd" {
		return ctx, nil
	}

	token, err := auth.AuthFromMD(ctx, "Bearer")
	if err != nil {
		return nil, status.Errorf(301, "请求错误, 无权限请求:"+err.Error())
	}
	if token == "" {
		return nil, status.Errorf(301, "登录已过期3")
	}

	jwtInfo, err := tools.ParseJWT(token)
	if err != nil {
		return nil, status.Errorf(301, "请求错误, 鉴权失败:"+err.Error())
	}

	// 对比updateTime 是否一样
	userInfo, err := cache.GetUserInfo(jwtInfo.Emails)
	if err != nil {
		return nil, err
	}
	if userInfo.Updatetime != jwtInfo.Updatetime {
		return nil, status.Errorf(301, "登录超时")
	}

	newCtx, err := tools.SetJwtInfo(ctx, userInfo)
	if err != nil {
		return nil, err
	}
	return newCtx, nil
}

func InterceptorLogger() logging.Logger {
	return logging.LoggerFunc(func(ctx context.Context, lvl logging.Level, msg string, fields ...any) {
		switch lvl {
		case logging.LevelDebug:
			msg = fmt.Sprintf("DEBUG :%v", msg)
		case logging.LevelInfo:
			msg = fmt.Sprintf("INFO :%v", msg)
		case logging.LevelWarn:
			msg = fmt.Sprintf("WARN :%v", msg)
		case logging.LevelError:
			msg = fmt.Sprintf("ERROR :%v", msg)
		default:
			panic(fmt.Sprintf("unknown level %v", lvl))
		}

		md, _ := metadata.FromIncomingContext(ctx)
		fields = append([]any{"msg", msg, "realAddress", md["x-forwarded-for"]}, fields...)
		log.Println(fields)
	})
}

func AllButHealthZfunc(ctx context.Context, callMeta interceptors.CallMeta) bool {
	return healthpb.Health_ServiceDesc.ServiceName != callMeta.Service
}
func CustomFunc(p any) (err error) {
	return status.Errorf(codes.Unknown, "panic triggered: ====>> %v", p)
}
