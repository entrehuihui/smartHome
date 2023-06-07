from umqtt.simple import MQTTClient
import time
import json
from machine import Pin

user = ""
password = ""
deviceSN = "sm202305241716"

topic = b'msg/'+user+'/'+deviceSN+'/pub'
topicPub = b'msg/'+user+'/'+deviceSN+'/pub'
client = MQTTClient('21312', 'webrtc.laoyinbi.cn', 1883,
                    deviceSN, password, keepalive=0, ssl=False, ssl_params={})

p2 = Pin(2, Pin.OUT)
p2.off()

def sub(topic, msg):
    # 在回调函数打印主题和消息
    msginfo = json.loads(msg)
    print(msginfo)
    if msginfo["type"] == '23':
        print("打开开关")
        p2.on()
        time.sleep(2)
        p2.off()
        print("发送执行成功信息")
        client.publish(topicPub, '{"type":"eply", "data":"'+msginfo["id"] +'"}', qos=0)

def mqttConnect():
    print("------------------------------")

    client.set_callback(sub)
    client.connect()
    print("-------------连接成功-----------------")
    # 订阅
    client.subscribe(topic)
    while True:
        client.wait_msg()

mqttConnect()
