package main

import (
	"fmt"
	tools2 "github.com/FC/FC/Ops/autoTools/tools"
	"log"
)
var (
	DeployMode 		string
	Start      		uint64
	End        		uint64
	CmdApiPort		uint64
	SwarmPort		uint64
	DeployPos  		string
	BinFilePos  	string
	TradeFilePos  	string


)
func main() {
	fmt.Printf("请输入mode，如果需要远程部署请输入Remote，否则跳过")
	_, err := fmt.Scanln(&DeployMode)
	if err!=nil{
		log.Fatalln("scan Error!")
	}
	switch DeployMode {
	default:
		fmt.Println("请输入部署起点数，结点数，apiPort起点数，swarmPort起点数，部署位置和交易文件位置（例：0 5 ～ 3455 6001 trade.log）")
		_, err = fmt.Scanln(&Start, &End, &DeployPos)
		if err!=nil{
			log.Fatalln("scan Error!")
		}
		tools2.LocalDeploy(Start,End, DeployPos)

	case "Remote":
		// 请输入远程部署IP列表（例：192.168.2.30 192.168.2.31）

		// 请输入部署起点数，结点数，和部署位置（例：0 5 ～）
		tools2.RemoteDeploy()
	}

}