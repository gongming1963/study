import socket

server = 'members.3322.org'
port = 80
PROXY_ADDR = ("192.168.2.223", 8088)
CONNECT = "CONNECT %s:%s HTTP/1.0\r\nConnection: close\r\n\r\n" % (server, port)

s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
s.connect(PROXY_ADDR)
s.send(CONNECT)
print s.recv(4096)      

host = "members.3322.org"
s.send("GET /dyndns/getip HTTP/1.1\n")
s.send("Accept:text/html,application/xhtml+xml,*/*;q=0.8\n")
s.send("Accept-Language:zh-CN,zh;q=0.8,en;q=0.6\n")
s.send("Cache-Control:max-age=0\n")
s.send("Connection:keep-alive\n")
s.send("Host:"+host+"\r\n")
s.send("Referer:http://www.baidu.com/\n")
s.send("user-agent: Googlebot\n\n")
while True:
    buf = s.recv(1024)
    if not len(buf):
        break
    print buf
