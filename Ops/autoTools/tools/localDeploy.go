package tools

import (
	functions2 "github.com/FC/FC/Ops/autoTools/tools/functions"
	"log"
)

// 本地自动化部署
/*
	参数：起点数字，结束点数字，部署文件位置
*/
func LocalDeploy(start,end uint64,pos string)  {
	for start<=end{
		go functions2.StartDeploy(start,pos)
		start++
	}
	log.Fatalln("Deploy Over ~~")
}
