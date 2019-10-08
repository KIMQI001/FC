package mySectorBuilder

import "github.com/filecoin-project/go-filecoin/proofs/sectorbuilder"

func RpcPostListener(sb sectorbuilder.SectorBuilder)  {
	sb.GeneratePoSt()
}
