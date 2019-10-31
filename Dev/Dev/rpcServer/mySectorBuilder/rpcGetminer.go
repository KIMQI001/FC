package mySectorBuilder

import (
	"context"
	"github.com/filecoin-project/go-filecoin/Dev/proto"
	"google.golang.org/grpc"
	"log"
)

func Getminer() string {
	// Set up a connection to the server.
	conn, err := grpc.Dial("127.0.0.1:1200", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := proto.NewRegisterServiceClient(conn)

	// Contact the server and print out its response.
	//name := "SoWhat"
	//if len(os.Args) > 1 {
	//	name = os.Args[1]
	//}
	r, err := c.GetMiner(context.Background(), &proto.GetMinerReq{})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Miner is : %s", r.Miner)
	return r.Miner
}