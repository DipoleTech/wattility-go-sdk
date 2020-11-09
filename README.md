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

#### 1.2.2 设置代理转发

设置转发地址

```
client.SetProxy()
```



## 2. 接口

### 2.1 检查签名

```go
# CheckSignApi
_, err :=client.CheckSignApi()
```

