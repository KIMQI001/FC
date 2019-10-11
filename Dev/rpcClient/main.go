package main

import (
	mySectorBuilder2 "github.com/FC/FC/Dev/rpcClient/mySectorBuilder"
	"log"
)

func main() {
	// Register

	// Get a genSectorBuilderRequest
	miner:=""
	dir:=""
	lastId:=0
	sb,err:= mySectorBuilder2.GenSectorBuilder(miner,dir,dir, uint64(lastId))

	if err != nil {
		log.Fatalln("Generate SectorBuilder failed~~~!!!!")
	}
	// SendBack a sb response

	// rpc addPiece task listener
	go mySectorBuilder2.RpcAddPieceListener(sb)

	// rpc Post task listener
	go mySectorBuilder2.RpcPostListener(sb)
}
