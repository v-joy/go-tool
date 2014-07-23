所有数据包都有2个字节的包头，紧跟着是包体，格式如下
| head | body |
其中包头为网络字节序

客户端先注册，不允许重名，注册成功后可以向服务端广播，然后每个注册的客户端都可以收到这个消息（除了发送者）。
要求实现 客户端和服务端

客户端注册
C -> S: {“action”:”register”, “data”:{“name”:”xxx”}}
S -> C: {“action”:”ack”, “data”:{“status”:200, “msg”:”ok”}}

客户端广播消息
C -> S: {“action”:”broadcast”, “data”:{“msg”:”xxx”}}
S -> C: {“status”:200, “msg”:”ok”}（注1）

服务端转发消息
S -> C[0..n](除了发送者): {“action”:”chat", “data”:{“name”:”xxx”, “msg”:”xxx”}}
C[0..n] -> S: {“status”:200, “msg”:”ok”}

注1:服务端不必等待所有客户端都收到后再发这条数据
