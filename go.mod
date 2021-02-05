module github.com/assetsadapterstore/tncd-adapter

go 1.14

require (
	github.com/astaxie/beego v1.12.0
	github.com/blocktree/futurepia-adapter v1.2.6
	github.com/blocktree/go-owcrypt v1.1.1
	github.com/blocktree/openwallet/v2 v2.0.10
	github.com/emirpasic/gods v1.12.0
	github.com/eoscanada/eos-go v0.8.10
	github.com/imroc/req v0.2.4
	github.com/pkg/errors v0.9.1
	github.com/shopspring/decimal v0.0.0-20200105231215-408a2507e114
	github.com/tidwall/gjson v1.3.5
)

//replace github.com/blocktree/openwallet => ../../openwallet
