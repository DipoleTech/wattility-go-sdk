# wattility-go-sdk

`wattility-go-sdk` lib, send data to `wattility` open api.



## 1.  使用

### 1.1 初始化SKD

初始化客户端，传入APP ID 和APP Secret

```go
client, err := wattility_go_sdk.NewClient(AppId, AppSecret)
if err != nil {
		panic(err)
}
```

### 1.2 Client设置

#### 1.2.1 设置接口地址

可以根据需要设置自定义的开发环境地址

```go
client.SetHost("")
```

#### 1.2.2 ~~设置代理转发~~

```
client.SetProxy()
```

#### 1.2.3 设置Debug

```go
client.SetDebug()
```

#### 1.2.4 设置返回消息处理

```go
client.SetRecHandle(handle *ReceiveHandle)
```

#### 1.2.5 开始连接

```go
go client.StartConn()
```



## 2. 接口

### 2.1 LoadBaseSummary

```go
client.LoadBaseSummary(body []LoadBaseSummaryBody): error
```

### 2.2 LoadBaseFactor

```go
client.LoadBaseFactor(body []LoadBaseFactorBody): error
```

