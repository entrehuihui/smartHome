from umqtt.simple import MQTTClient
import time
import json
from machine import Pin

user = ""
password = ""
deviceSN = "sm202305241716"

def listTCP():
    global user
    global password
    import  socket
    sock = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
    sock.bind(("", 9000))
    sock.listen(5)
    dat = b''
    sendData = deviceSN + '#lend1#'
    while True:
        print("准备接收")
        connection,address = sock.accept()
        
        print('connect from ', address)
        # connection.timeout(50)
        try:
            while True:
                dat=b''
                while True:
                    data = connection.recv(10240)
                    if len(data) == 0:
                        break
                    dat = dat+data
                    if data.find(b'#lend1#')>-1:
                        dat=dat[:-7]
                        print('end')
                        break
                print("====>>", dat)
                connection.send(sendData.encode())
                if len(dat) > 4:
                    datString = str(dat, 'utf-8')
                    print(datString)
                    info = datString.split(':')
                    print(info)
                    if len(info) != 3:
                        break
                    if info[0]!= "info":
                        break
                    if info[1]=="" or info[2] == "":
                        break
                    user = info[1]
                    password = info[2]
                    return
                    
        except:
            print("发送异常 重新连接")


        connection.close()

listTCP()
print("退出tcp", user, password)
