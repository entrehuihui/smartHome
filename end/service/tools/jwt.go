package tools

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"test/service/databases"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// const tokenkey key = "tokenkey"

// // GetTokenValue .
// func GetTokenValue(ctx context.Context) *string {
// 	token := ctx.Value(tokenkey).(string)
// 	return &token
// }

// // SetTokenValue .
// func SetTokenValue(ctx context.Context, token string) context.Context {
// 	return context.WithValue(ctx, tokenkey, token)
// }

// exp JWT过期时间 默认6小时=21600秒
var exp = 21600
var hmacSampleSecret = []byte("saleAuthenticationCenter")

// SetExp 设置过期和解密字符串时间
func SetExp(sec int, hmacSample string) {
	if sec > 0 {
		exp = sec
	}
	if hmacSample != "" {
		hmacSampleSecret = []byte(hmacSample)
	}
}

// Jwtinfo .
type Jwtinfo struct {
	ID         int
	Nicename   string
	Updatetime int64
	Emails     string
}

// NewJWT .
func NewJWT(jwtinfo Jwtinfo) (string, error) {

	body, err := json.Marshal(jwtinfo)
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"g":   string(body),
		"exp": time.Now().Add(time.Second * time.Duration(exp)).Unix(),
	})
	tokenString, err := token.SignedString(hmacSampleSecret)
	if err != nil {
		return "", err
	}

	// 缓存
	return tokenString, err
}

// ParseJWT .
func ParseJWT(tokenString string) (*Jwtinfo, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return hmacSampleSecret, nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	// fmt.Println(claims)
	if !ok || !token.Valid {
		return nil, errors.New("not authorized")
	}
	body, err := getKeyByteValue(&claims, "g")
	if err != nil {
		return nil, err
	}
	jwtinfo := new(Jwtinfo)
	err = json.Unmarshal([]byte(body), jwtinfo)
	if err != nil {
		return jwtinfo, err
	}
	if jwtinfo.ID == 0 {
		return nil, errors.New("登录信息已过期")
	}
	return jwtinfo, err
}

func getKeyByteValue(claims *jwt.MapClaims, key string) (string, error) {
	body, ok := (*claims)[key].(string)
	if !ok {
		return "", errors.New(key + " is error")
	}
	return body, nil
}

type key string

const (
	// GradeID .角色ID
	UserinfoKey key = "userinfo"
)

// GetJwtInfo .
func GetJwtInfo(ctx context.Context) *databases.Userinfo {
	userInfos := ctx.Value(UserinfoKey).(string)
	userInfo := new(databases.Userinfo)
	json.Unmarshal([]byte(userInfos), userInfo)
	return userInfo
}

func SetJwtInfo(ctx context.Context, userInfo *databases.Userinfo) (context.Context, error) {
	b, err := json.Marshal(userInfo)
	if err != nil {
		return ctx, err
	}

	newCtx := context.WithValue(ctx, UserinfoKey, string(b))
	return newCtx, err
}
