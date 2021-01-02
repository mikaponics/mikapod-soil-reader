package app // github.com/mikaponics/mikapod-storage/internal

import (
    "log"
    "net"
    "net/rpc"
    "net/http"

	"github.com/mikaponics/mikapod-soil-reader/configs"
    "github.com/mikaponics/mikapod-soil-reader/internal/rpc_server"
    "github.com/mikaponics/mikapod-soil-reader/internal/util"
)

type MikapodSoilReader struct {
    tcpAddr *net.TCPAddr
    listener *net.TCPListener
    rpcServer *rpc_server.RPC
}

func InitMikapodSoilReader(arduinoDevicePath string) (*MikapodSoilReader) {
    tcpAddr, err := net.ResolveTCPAddr("tcp", configs.MikapodSoilReaderServiceAddress)
	if err != nil {
		log.Fatal(err)
	}

    ar := util.ArduinoReaderInit(arduinoDevicePath)

    r := &rpc_server.RPC{
        ArduinoReader: ar,
        IsRunning: true,
	}

    log.Println("RPC API was initialized.")
    return &MikapodSoilReader{
        tcpAddr: tcpAddr,
        listener: nil,
        rpcServer : r,
    }
}


// Function will consume the main runtime loop and run the business logic
// of the Mikapod Logger application.
func (app *MikapodSoilReader) RunMainRuntimeLoop() {
    rpc.Register(app.rpcServer)
	rpc.HandleHTTP()
	log.Println("RPC was initialized.")
	l, e := net.ListenTCP("tcp", app.tcpAddr)
    app.listener = l // Track the `listener` so we can gracefully shutdown later.
	if e != nil {
		log.Fatal("listen error:", e.Error())
	}
	log.Println("Started soil reader service.")
	http.Serve(l, nil)
}

// Function will tell the application to stop the main runtime loop when
// the process has been finished.
func (app *MikapodSoilReader) StopMainRuntimeLoop() {
    log.Printf("Starting graceful shutdown now...")

    // Finish any RPC communication taking place at the moment before
    // shutting down the RPC server.
    app.listener.Close()
}
