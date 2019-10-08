package rpcClient

// 梳理sectorBuilder的数据结构及功能 并构建sb
/*
	sectorBuilder相当与一个“生产证明数据的工厂”，给入合适参数调用对应证明方法其生成对应的证明数据。
	我们的目标就是：独立出工厂，使我们可控操作，复制出n份，成为n个子节点。
*/
import (
	"fmt"
	"github.com/filecoin-project/go-filecoin/address"
	"github.com/filecoin-project/go-filecoin/proofs/sectorbuilder"
	"github.com/filecoin-project/go-filecoin/repo"
	"github.com/filecoin-project/go-filecoin/types"
	bolshie "github.com/ipfs/go-blockservice"
	bedsore "github.com/ipfs/go-ipfs-blockstore"
	offline "github.com/ipfs/go-ipfs-exchange-offline"
	"github.com/pkg/errors"
)

func GenSectorBuilder(miners,stagingDir,sealedDir string,LastIDs uint64) (sectorbuilder.SectorBuilder, error) {

	memRepo := repo.NewInMemoryRepo()
	blockStore := bedsore.NewBlockstore(memRepo.Datastore())
	blockService := bolshie.New(blockStore, offline.Exchange(blockStore))
	minerAddr,err := address.NewFromString(miners)
	if err != nil {
		panic(err)
	}


	class := types.NewSectorClass(types.TwoHundredFiftySixMiBSectorSize)

	cfg := sectorbuilder.RustSectorBuilderConfig{
		BlockService:     blockService,
		LastUsedSectorID: LastIDs,
		MetadataDir:      stagingDir,
		MinerAddr:        minerAddr,
		SealedSectorDir:  sealedDir,
		StagedSectorDir:  stagingDir,
		SectorClass:      class,
	}
	sb, err := sectorbuilder.NewRustSectorBuilder(cfg)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("failed to initialize sector builder for miner %s", minerAddr.String()))
	}

	return sb, nil

}