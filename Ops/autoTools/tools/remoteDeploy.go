package tools

// 自动化远程批量部署服务器
/*
	部署流程：
	1.维护一张服务器表（切片）
	2.遍历这个切片，并发发送 bin 文件，并执行部署程序
	3.远程部署daemon放在后台
	
*/

type ServerMap struct {
	
}

func RemoteDeploy()  {
	
}