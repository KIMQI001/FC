package functions

// 起一个命令行，重复执行发起交易命令
/*
	发起交易机制：

+ export CID=`go-filecoin client import ./hello.txt`
+ go-filecoin client list-asks --enc=json | jq
+ go-filecoin client propose-storage-deal <miner> $CID <ask> <duration> --allow
*/

// 初始化交易配置
func InitTradeCfg()  {

}

// 执行交易命令
func StartTrade()  {

}