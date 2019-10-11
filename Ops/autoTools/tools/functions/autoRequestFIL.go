package functions

// 起一个命令行，重复执行申请FIL命令
/*
	申请FIL流程：
+ export WALLET_ADDR=`go-filecoin address ls --repodir=''`
+ curl -X POST -F "target=${WALLET_ADDR}" "http://user.kittyhawk.wtf:9797/tap" | cut -d" " -f4

*/

func StartRequestFILs() {

}