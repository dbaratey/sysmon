package main

import (
	"context"
	"encoding/json"
	"github.com/dbaratey/sysmon"
	"google.golang.org/grpc"
	"log"
	"os"
)

func main() {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	opts = append(opts, grpc.WithBlock())
	conn, err := grpc.Dial("localhost:10000", opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := sysmon.NewSysmonAgentClient(conn)
	aCl, err := client.GetStats(context.Background(), &sysmon.NullReq{})
	if err != nil {
		panic(err)
	}

	for {
		m, err := aCl.Recv()
		if err != nil {
			log.Fatal(err)
		}
		json.NewEncoder(os.Stdout).Encode(&m)
	}

}
