package rpc_client

import (
	// "errors"
	"log"
	"net/rpc"
)

type MikapodSoilReaderService struct {
	Client *rpc.Client
}

func New(addr string) *MikapodSoilReaderService {
	if addr == "" {
		log.Fatal("LOG | MikapodSoilReaderService | New | No address set.")
	}

	client, err := rpc.DialHTTP("tcp", addr)
	if err != nil {
		log.Fatal("ERROR | MikapodSoilReaderService | New | Dialing TCP Error:", err)
	}

	return &MikapodSoilReaderService{
		Client: client,
	}
}
