def do_connect():
    import network
    import time
    wlan = network.WLAN(network.STA_IF)
    wlan.active(False)
    print("wlan.isconnected()", wlan.isconnected())
    wlan.active(True)
    print("wlan.isconnected()", wlan.isconnected())
    print(wlan.scan())
    if not wlan.isconnected():
        print('connecting to network...')
        wlan.connect('20Fone', '12345678wcflss')
        while not wlan.isconnected():
            pass
    print("wlan.isconnected()", wlan.isconnected())
    print('network config:', wlan.ifconfig())


do_connect()
