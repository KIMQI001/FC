package mySectorBuilder

import (
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {}

func DMineRpcListener()  {
	// listen a port ,if there is a addPiece request,start AddPiece
	//sb.AddPiece()
	lis,err := net.Listen("tcp",":1200")
	if err !=nil{
		log.Fatalf("faild to listen: %v",err)
	}
	s :=grpc.NewServer()

	_ = s.Serve(lis)
}

func (s *server)AddPiece()  {
	GetSbPtr()
}
func (s *server)PoSt()  {

}
func (s *server)ReadPiece()  {

}
