import socket
import time

import requests
import json

dataStr = json.dumps({"name": "模拟设备","sn": "sm202305241716"})
print(dataStr)
resp = requests.post("https://smarthome.laoyinbi.cn/v1/device/register", dataStr, headers={
    "authorization":"Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2ODUzNTQ3MDYsImciOiJ7XCJJRFwiOjEsXCJOaWNlbmFtZVwiOlwi5a-G56CBXCIsXCJVcGRhdGV0aW1lXCI6MTY4NTA5Mzg0MzE5MSxcIkVtYWlsc1wiOlwiMjczMjI3NTQzQHFxLmNvbVwifSJ9.OolstGIYe_HUtLuN3tjljU-AJuibZYDgfnXbHr1SePE"
})

dataJson = resp.json()
print(dataJson)
if 'code' in dataJson :
    print(dataJson)
    exit(0)
token = dataJson['data']


client = socket.socket(socket.AF_INET, socket.SOCK_STREAM)

# client.connect(("127.0.0.1", 9000))
client.connect(("192.168.1.6", 9000))

def send(data):
    print("send ====>>>>>>>>>")
    data = data + '#lend1#'
    client.send(data.encode())
    dat = b''
    while True:
        data = client.recv(10240)
        print(len(data))
        if len(data) == 0:
            break
        dat = dat+data
        if data.find(b'#lend1#') > -1:
            dat = dat[:-7]
            print('end  ====>>>>>>>>>')
            break
    print(dat)


send("find")
send("info:273227543@qq.com:" + token)
client.close()
