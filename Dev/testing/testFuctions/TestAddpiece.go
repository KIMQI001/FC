package testFuctions

// 构建虚拟密封过程，来进行硬件测试
import (
	"bytes"
	"context"
	"crypto/rand"
	"github.com/FC/FC/Dev/rpcClient/mySectorBuilder"
	"github.com/filecoin-project/go-filecoin/address"
	"github.com/filecoin-project/go-filecoin/proofs/sectorbuilder"
	"github.com/filecoin-project/go-filecoin/proofs/verification"
	"github.com/filecoin-project/go-filecoin/types"
	"github.com/ipfs/go-cid"
	dag "github.com/ipfs/go-merkledag"
	"github.com/pkg/errors"
	"io"
	"log"
)

func TestAddPiece()  error{
	miner:=""
	dir:=""
	lastId:=0
	sb,err:=mySectorBuilder.GenSectorBuilder(miner,dir,dir, uint64(lastId))
	if err != nil {
		log.Fatalln("Generate SectorBuilder failed~~~!!!!")
	}

	inputBytes := RequireRandomBytes(types.TwoHundredFiftySixMiBSectorSize.Uint64())
	ref, size, reader, err := CreateAddPieceArgs(inputBytes)
	sectorID, err := sb.AddPiece(context.Background(), ref, size, reader)
	if err!=nil{
		return errors.Wrap(err,"addPiece failed~~~!!!")
	}
	select {
	case val := <-sb.SectorSealResults():
		if val.SealingResult.SectorID!=sectorID{
			return errors.Wrap(err,"sectorID is not same~~~!!!")
		}
		minerAddr,err := address.NewFromString(miner)
		if err != nil {
			return errors.Wrap(err,"gen miner addr failed~~~!!!")
		}

		sres, serr := (&verification.RustVerifier{}).VerifySeal(verification.VerifySealRequest{
			CommD:      val.SealingResult.CommD,
			CommR:      val.SealingResult.CommR,
			CommRStar:  val.SealingResult.CommRStar,
			Proof:      val.SealingResult.Proof,
			ProverID:   sectorbuilder.AddressToProverID(minerAddr),
			SectorID:   val.SealingResult.SectorID,
			SectorSize: types.TwoHundredFiftySixMiBSectorSize,
		})
		if serr!=nil{
			return errors.Wrap(serr,"verify failed~~!!")
		}
		if sres.IsValid!=true{
			return errors.New("verify is not valid~~!!")
		}
		return nil
	}
}



// CreateAddPieceArgs creates a PieceInfo for the provided bytes
func CreateAddPieceArgs(pieceData []byte) (cid.Cid, uint64, io.Reader, error) {
	data := dag.NewRawNode(pieceData)

	return data.Cid(), uint64(len(pieceData)), bytes.NewReader(pieceData), nil
}

// RequireRandomBytes produces n-number of bytes
func RequireRandomBytes( n uint64) []byte { // nolint: deadcode
	slice := make([]byte, n)

	if _, err := io.ReadFull(rand.Reader, slice);err!=nil{
		log.Fatalln("RequireRandomBytes failed~~!!!")
	}

	return slice
}
