package internal // github.com/mikaponics/mikapod-soil-reader/internal

import (
    "log"
	"net"
    "os"

	"google.golang.org/grpc"

	pb "github.com/mikaponics/mikapod-soil-reader/api"
	"github.com/mikaponics/mikapod-soil-reader/configs"
)

type MikapodSoilReader struct {
    isRunning bool
    arduinoReader *ArduinoReader
    grpcServer *grpc.Server
}

func InitMikapodSoilReader(arduinoDevicePath string) (*MikapodSoilReader) {
    ar := ArduinoReaderInit(arduinoDevicePath)
    return &MikapodSoilReader{
        isRunning: true,
        arduinoReader: ar,
        grpcServer: nil,
    }
}


// Function will consume the main runtime loop and run the business logic
// of the Mikapod Logger application.
func (app *MikapodSoilReader) RunMainRuntimeLoop() {
    // Open a TCP server to the specified localhost and environment variable
    // specified port number.
    lis, err := net.Listen("tcp", configs.MikapodSoilReaderServiceAddress)
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }

    // Initialize our gRPC server using our TCP server.
    grpcServer := grpc.NewServer()

    // Save reference to our application state.
    app.grpcServer = grpcServer

    // For debugging purposes only.
    log.Printf("READER: gRPC server ready and running.")

    // Block the main runtime loop for accepting and processing gRPC requests.
    pb.RegisterMikapodSoilReaderServer(grpcServer, &MikapodSoilReaderGRPC{
        // DEVELOPERS NOTE:
        // We want to attach to every gRPC call the following variables...
        arduinoReader: app.arduinoReader,
    })
    if err := grpcServer.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}

// Function will tell the application to stop the main runtime loop when
// the process has been finished.
func (app *MikapodSoilReader) StopMainRuntimeLoop() {
    log.Printf("READER: Starting graceful shutdown now...")

    // Finish any RPC communication taking place at the moment before
    // shutting down the gRPC server.
    app.grpcServer.GracefulStop()

	// app.storageCon.Close()
    os.Exit(1)
}
