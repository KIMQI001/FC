package mySectorBuilder

import (
	"context"
	"github.com/filecoin-project/go-filecoin/Dev/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)



const (
	 POREP int = 1
	 PoST int =2
)
type server struct {}

func DMineRpcListener()  {
	// listen a port
	lis,err := net.Listen("tcp",":1200")
	if err !=nil{
		log.Fatalf("faild to listen: %v",err)
	}
	s :=grpc.NewServer()
	proto.RegisterDMineServiceServer(s,&server{})
	_ = s.Serve(lis)
}

func (s *server)GenProof(ctx context.Context, in *proto.ProofReq) (*proto.ProofRes, error)  {
	// judge the proof title
	ProofSchedule(in)
	return &proto.ProofRes{}, nil
}

func (s *server)GetMetaData(ctx context.Context, in *proto.ProofReq) (*proto.ProofRes, error)  {
	GetSbPtr().SectorSealResults()
	return &proto.ProofRes{}, nil
}

func (s *server)ReadPiece(ctx context.Context, in *proto.ReadPieceReq) (*proto.ReadPieceRes, error) {

	return &proto.ReadPieceRes{}, nil
}
