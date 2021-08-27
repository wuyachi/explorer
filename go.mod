module explorer

go 1.14

require (
	github.com/astaxie/beego v1.12.3
	github.com/ethereum/go-ethereum v1.9.24
	github.com/fsnotify/fsnotify v1.4.9 // indirect
	github.com/kylelemons/godebug v1.1.0 // indirect
	github.com/mitchellh/go-testing-interface v1.0.0 // indirect
	github.com/palettechain/palette-go-sdk v0.0.1
	github.com/shiena/ansicolor v0.0.0-20200904210342-c7312218db18 // indirect
	github.com/shopspring/decimal v1.2.0
	github.com/smartystreets/goconvey v1.6.4
	github.com/urfave/cli v1.22.4
	google.golang.org/genproto v0.0.0-20191108220845-16a3f7862a1a // indirect
	gopkg.in/yaml.v2 v2.3.0 // indirect
	gorm.io/driver/mysql v1.0.3
	gorm.io/gorm v1.20.7
)

replace (
	github.com/coreos/etcd v0.0.1 => github.com/polynetwork/coreos-etcd v0.0.1
	github.com/coreos/go-semver v0.0.1 => github.com/polynetwork/coreos-semver v0.0.1
	github.com/coreos/go-systemd v0.0.1 => github.com/polynetwork/coreos-systemd v0.0.1
	github.com/coreos/pkg v0.0.1 => github.com/polynetwork/coreos-pkg v0.0.1
	github.com/ethereum/go-ethereum v1.9.24 => ../palette
	github.com/palettechain/palette-go-sdk v0.0.1 => ../palette-go-sdk
)
