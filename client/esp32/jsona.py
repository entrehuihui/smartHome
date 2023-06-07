import json

a = '{"key":"connect",  "msg":"hello world!"}'


ac = json.loads(a)

print(ac)
print(ac['key'])
