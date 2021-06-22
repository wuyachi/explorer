# palette explorer部署文档

## 代码

[github](https://github.com/palettechain/explorer)

## build

### 编译扫描程序palette_server
```
cd ./cmd
go build main.go
```

### 编译服务程序palette_http

```
cd .
go build main.go
```


## 启动palette_server

### 配置

```
{
  "NodeConfig": {
    "RestURL": "http://106.75.250.160:22000",
    "Admin": "0xf3A9d42C01635A585f1721463842F8936075105F",
    "WalletFile": "./wallet/wallet.dat",
    "WalletPwd": "4cUYqGj2yib718E7ZmGQc"
  },
  "Statistic": 1,
  "DBConfig": {
    "URL": "localhost:3306",
    "User": "root",
    "Password": "root",
    "Scheme": "palette"
  }
}
```

NodeConfig.RestURL: palette链的节点

NodeConfig.Admin: palette链的管理员地址

Statistic: 1: 做交易统计  2: 不做交易统计

DBConfig: 数据库配置

### 启动

```
./palette_server --cliconfig config.json
```


## 启动palette_http

### 配置

```
appname = explorer
httpport = 22000
runmode = dev
copyrequestbody = true
mysqluser = "root"
mysqlpass = "root"
mysqlurls = "127.0.0.1:3306"
mysqldb   = "palette"
mysqldebug = true
```

httpport: 服务端口

mysql: 数据库配置

配置文件置于palette_http同目录下的conf子目录中


### 启动

```
./palette_http
```

