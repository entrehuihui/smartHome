package cache

import (
	"encoding/json"
	"test/service/databases"

	"github.com/gomodule/redigo/redis"
)

// 用户信息缓存
const userInfoKey = "userInfoKey"

func GetUserInfoKey(emails string) string {
	return userInfoKey + ":" + emails
}
func SetRedisUserInfo(userInfo *databases.Userinfo) error {
	b, err := json.Marshal(userInfo)
	if err != nil {
		return err
	}

	_, err = databases.GetRedis().Do("SET", GetUserInfoKey(userInfo.Emails), b, "EX", 259200)
	return err
}

func GetUserInfo(emails string) (*databases.Userinfo, error) {
	dataBuf, err := redis.String(databases.GetRedis().Do("GET", GetUserInfoKey(emails)))
	if err != nil {
		userInfo := new(databases.Userinfo)
		err = databases.GetDB().Where("emails = ? AND del = 0", emails).Find(userInfo).Limit(1).Error
		if err != nil {
			return nil, err
		}
		go SetRedisUserInfo(userInfo)
		return userInfo, err
	}
	userInfo := new(databases.Userinfo)
	err = json.Unmarshal([]byte(dataBuf), userInfo)
	return userInfo, err
}

func DelUserInfo(emails string) {
	redis.String(databases.GetRedis().Do("DEL", GetUserInfoKey(emails)))
}

// 用户code缓存
const userCodeKey = "userCodeKey"

func GetUserCodeKey(emails, types string) string {
	return userCodeKey + ":" + emails + ":" + types
}

func SetUserCode(emails, types, code string) error {
	_, err := databases.GetRedis().Do("SET", GetUserCodeKey(emails, types), code, "EX", 300)
	return err
}

func GetUserCode(emails, types string) (string, error) {
	dataBuf, err := redis.String(databases.GetRedis().Do("GET", GetUserCodeKey(emails, types)))
	return dataBuf, err
}

func DelUserCode(emails, types string) {
	redis.String(databases.GetRedis().Do("DEL", GetUserCodeKey(emails, types)))
}
