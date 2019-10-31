package main

import (
	"github.com/filecoin-project/go-filecoin/Dev/rpcServer/mySectorBuilder"
	"log"
)

func main() {
	// Register

	// Get a genSectorBuilderRequest
	miner:="t2c5rchhtyscadk7hduwma45jx44w4ercw7nzh3vy"
	dir:="/home/zoe/files/sectors"
	lastId:=0
	sb,err:= mySectorBuilder.NewSectorBuilderMod(miner,dir,dir, uint64(lastId))
	if err != nil {
		log.Fatalln("Generate SectorBuilder failed~~~!!!!")
	}
	mySectorBuilder.SetSbPtr(sb)
	// SendBack a sb response

	// rpc addPiece task listener
	go mySectorBuilder.DMineRpcListener()

}
