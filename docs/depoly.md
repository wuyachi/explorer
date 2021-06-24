# palette explorer deploy

## build from source code

### code

[source code](https://github.com/palettechain/explorer)

### build

#### build palette_server
```
cd ./cmd
go build main.go
```

#### build palette_http

```
cd .
go build main.go
```

## get the package

[palette explorer package](https://github.com/palettechain/explorer/releases/tag/v0.1.1)

## install palette_server

### config

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

NodeConfig.RestURL: the node url of palette chain

NodeConfig.Admin: the admin address of palette chain

Statistic: 1: start the statistic of transactions  2: don't start the statistic of transactions

DBConfig: the config of database

### start

```
./palette_server --cliconfig config.json
```


## install palette_http

### config

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

httpport: the port

mysql: the config of database

please put the configuration file in the directory ./conf


### start

```
./palette_http
```

