package main

import (
    "fmt"
    "time"

    sr "github.com/mikaponics/mikapod-soil-reader/internal"
)


/**
 * Application used to interface with the hardware instruments (ex: humidity,
 * temperature, etc) and continously save the latest data.
 */
func main() {
    ar := sr.ArduinoReaderInit()
    for {
        data := ar.Read()
        fmt.Printf("%+v\n\n", data)
        time.Sleep(2 * time.Second)
    }
    fmt.Printf("Current Unix Time: %v\n", time.Now().Unix())
}
