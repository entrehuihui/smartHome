package cache

import (
	"encoding/json"
	"test/service/databases"

	"github.com/gomodule/redigo/redis"
)

// 设备信息缓存
const deviceInfoKey = "deviceInfoKey"

func GetDeviceInfoKey(sn string) string {
	return deviceInfoKey + ":" + sn
}

func SetDeviceInfo(deviceInfo *databases.Device) error {
	b, err := json.Marshal(deviceInfo)
	if err != nil {
		return err
	}

	_, err = databases.GetRedis().Do("SET", GetDeviceInfoKey(deviceInfo.Sn), b, "EX", 259200)
	return err
}

func GetDeviceInfo(sn string) (*databases.Device, error) {
	dataBuf, err := redis.String(databases.GetRedis().Do("GET", GetDeviceInfoKey(sn)))
	if err != nil {
		deviceInfo := new(databases.Device)
		err = databases.GetDB().Where("sn = ? AND del = 0", sn).Find(deviceInfo).Limit(1).Error
		if err != nil {
			return nil, err
		}
		deviceInfo.Sn = sn
		go SetDeviceInfo(deviceInfo)
		return deviceInfo, err
	}
	deviceInfo := new(databases.Device)
	err = json.Unmarshal([]byte(dataBuf), deviceInfo)
	return deviceInfo, err
}

func DelUDeviceInfo(sn string) {
	redis.String(databases.GetRedis().Do("DEL", GetDeviceInfoKey(sn)))
}
