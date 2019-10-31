package mySectorBuilder

import (
	"github.com/filecoin-project/go-filecoin/Dev/proto"
	"sync"
)

//梳理与sb功能相关数据结构

/*
	梳理出打造一个sectorBuilder及使用其相应功能所需的“部件”（即相关的数据结构）
*/

var PoStLock sync.Locker

func ProofSchedule(in *proto.ProofReq)  {
	// judge the title , arrange the task

	// judge the if PoSt process on LOCK mode

}
func AddPiece()  {
	GetSbPtr()
}

func PoSt()  {
	// process LOCK

}
