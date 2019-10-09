package functions

// 自动初始化配置并运行filecoin
/*
	初始化流程：
+ rm -rf ~/.filecoin --repodir=''
+ go-filecoin init --devnet-user --genesisfile=https://genesis.user.kittyhawk.wtf/genesis.car --repodir=''
+ go-filecoin daemon --repodir=''
+ go-filecoin config heartbeat.nickname "hkiCloud" --repodir=''
+ go-filecoin config heartbeat.beatTarget "/dns4/backend-stats.kittyhawk.wtf/tcp/8080/ipfs/QmUWmZnpZb6xFryNDeNU7KcJ1Af5oHy7fB9npU67sseEjR" --repodir=''

*/

// 设置初始化数量和初始化配置参数
func StartInitAndRuning(nums uint64,pos string)  {

}

// 自动初始化并运行filecoin
func AutoInit()  {

}
