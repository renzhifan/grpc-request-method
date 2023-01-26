# grpc-example
## 1.go-grpc的四种请求方式

### 1-1.一元RPC方法：

客户端发起一个请求，从服务端得到一个响应，和普通函数调用相似。

```go
rpc SayHello (HelloRequest) returns (HelloResponse) {}
```

### 1-2.服务端流：

客户端发起一个请求，从服务端得到一个消息流，客户端可以一直从流中读取数据直到读取完服务端返回的所有数据。

```go
 rpc GetStream (HelloRequest) returns (stream HelloResponse) {}
```

### 1-3.客户端流：

客户端发起一个消息流到服务端，服务端收到数据流后执行响应。

```go
rpc SetStream (stream HelloRequest) returns (HelloResponse) {}
```

### 1-4.双向流：

客户端发起一个消息流到服端，服务端收到流数据后返回另外一个数据流给客户端。

```go
rpc AllStream (stream HelloRequest) returns (stream HelloResponse) {}
```