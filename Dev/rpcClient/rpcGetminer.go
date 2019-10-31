package rpcClient

import (
	"context"
	"github.com/filecoin-project/go-filecoin/Dev/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {}

func RegisterRpcListener()  {
	// listen a port
	lis,err := net.Listen("tcp",":1200")
	if err !=nil{
		log.Fatalf("faild to listen: %v",err)
	}
	s :=grpc.NewServer()
	proto.RegisterRegisterServiceServer(s,&server{})
	_ = s.Serve(lis)
}

func (s *server)GetMiner(ctx context.Context, in *proto.GetMinerReq) (*proto.GetMinerRes, error)  {
	// register as a subNode

	// return the node miner

	return &proto.GetMinerRes{}, nil
}
