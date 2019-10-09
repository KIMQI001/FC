package functions

// 自动化部署
/*
	部署流程：1.初始化并运行
			2.申请FIL
			3.发起自动化交易

+ go-filecoin wallet balance ${WALLET_ADDR} --repodir=''

*/

func StartDeploy(nums uint64,pos string)  {
	StartInitAndRuning(nums,pos)
	go StartRequestFILs()

	InitTradeCfg()
	for {
		// 查询余额，判断余额如果大于10,那么开始发起交易
		if FILs>10{
			StartTrade()

		}
	}

}