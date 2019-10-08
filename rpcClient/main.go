package main

import (
	"github.com/FC/FC/rpcClient/mySectorBuilder"
	"log"
)

func main() {
	// Register

	// Get a genSectorBuilderRequest
	miner:=""
	dir:=""
	lastId:=0
	sb,err:=mySectorBuilder.GenSectorBuilder(miner,dir,dir, uint64(lastId))

	if err != nil {
		log.Fatalln("Generate SectorBuilder failed~~~!!!!")
	}
	// SendBack a sb response

	// rpc addPiece task listener
	go mySectorBuilder.RpcAddPieceListener(sb)

	// rpc Post task listener
	go mySectorBuilder.RpcPostListener(sb)
}
