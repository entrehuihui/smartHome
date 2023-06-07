
x = 50


def fun():
    global x
    print("---1", x)
    x = 999
    print("----2", x)


fun()
print("====3", x)
