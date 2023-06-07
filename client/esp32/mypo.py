import requests
import json

dataStr = json.dumps({"name": "模拟设备","sn": "sm202305241716","emails": "273227543@qq.com","password": "aa13639bd9a58d63fe59b1c81eb3ff49","time": "1685065246000"})
print(dataStr)
resp = requests.post("https://smarthome.laoyinbi.cn/v1/device/register", dataStr)

dataJson = resp.json()
print(dataJson['data'])
