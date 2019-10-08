package mySectorBuilder
import "github.com/filecoin-project/go-filecoin/proofs/sectorbuilder"

func RpcAddPieceListener(sb sectorbuilder.SectorBuilder)  {

	// listen a port ,if there is a addPiece request,start AddPiece
	sb.AddPiece()
}

