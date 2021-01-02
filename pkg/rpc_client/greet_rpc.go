package rpc_client

import (
	// "errors"
	"log"
)

// Perform a synchronous call on our remote service for the `Greet` RPC.
func (s *MikapodSoilReaderService) Greet(name string) (*string, error) {
	var reply string
	err := s.Client.Call("RPC.Greet", &name, &reply)
	if err != nil {
		log.Println("ERROR | MikapodSoilReaderService | Greet | err", err)
		return nil, err
	}
	return &reply, nil
}
