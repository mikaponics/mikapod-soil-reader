package rpc_server

import (
	"github.com/mikaponics/mikapod-soil-reader/internal/util"
)

type RPC struct{
    IsRunning bool
    ArduinoReader *util.ArduinoReader
}
