package functions

// 起一个命令行，重复执行发起交易命令
/*
	发起交易机制：

+ export CID=`go-filecoin client import ./TradeFile`
+ go-filecoin swarm connect <peerID>
+ go-filecoin client list-asks --enc=json | jq
+ go-filecoin client propose-storage-deal <miner> $CID <ask> <duration> --allow
*/

// 初始化交易配置
func InitTradeCfg()  {
	// 检测交易文件是否存在，如果不存在，执行`dd if=/dev/urandom of=TradeFile bs=1M count=256`

	// export fileCID

	// 请输入<peerID> <miner> <ask>
}

// 执行交易命令
func StartTrade()  {

}