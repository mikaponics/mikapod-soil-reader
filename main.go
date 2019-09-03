package main

import (
	"os"
	"os/signal"
	"syscall"

    sr "github.com/mikaponics/mikapod-soil-reader/internal"
)

// DEVELOPERS NOTE:
// - Don't forget to run the following in the console:
//       export MIKAPOD_SOIL_READER_DEVICE_PATH=/dev/cu.usbmodem1D132101
// - Please replace the device path value with your devices value!

/**
 * Application used to interface with the hardware instruments (ex: humidity,
 * temperature, etc) and continously save the latest data.
 */
func main() {
    arduinoDevicePath := os.Getenv("MIKAPOD_SOIL_READER_DEVICE_PATH")
    app := sr.InitMikapodSoilReader(arduinoDevicePath)

    // DEVELOPERS CODE:
	// The following code will create an anonymous goroutine which will have a
	// blocking chan `sigs`. This blocking chan will only unblock when the
	// golang app receives a termination command; therfore the anyomous
	// goroutine will run and terminate our running application.
	//
	// Special Thanks:
	// (1) https://gobyexample.com/signals
	// (2) https://guzalexander.com/2017/05/31/gracefully-exit-server-in-go.html
	//
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
        <-sigs // Block execution until signal from terminal gets triggered here.
        app.StopMainRuntimeLoop()
    }()

    app.RunMainRuntimeLoop()
}
