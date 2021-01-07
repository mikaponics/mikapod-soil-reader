package util // github.com/mikaponics/mikapod-soil-reader/internal/util

import (
    "log"
    "time"
    "encoding/json"

    "github.com/tarm/serial"
)

const RX_BYTE = "1";

// The time-series data structure used to store all the data that will be
// returned by the `Mikapod Soil` Arduino device.
type TimeSeriesData struct {
    Status string `json:"status,omitempty"`
    Runtime int `json:"runtime,omitempty"`
    Id int `json:"id,omitempty"`
    HumidityValue float32 `json:"humidity_value,omitempty"`
    HumidityUnit string `json:"humidity_unit,omitempty"`
    TemperatureValue float32 `json:"temperature_primary_value,omitempty"`
    TemperatureUnit string `json:"temperature_primary_unit,omitempty"`
    PressureValue float32 `json:"pressure_value,omitempty"`
    PressureUnit string `json:"pressure_unit,omitempty"`
    TemperatureBackupValue float32 `json:"temperature_secondary_value,omitempty"`
    TemperatureBackupUnit string `json:"temperature_secondary_unit,omitempty"`
    AltitudeValue float32 `json:"altitude_value,omitempty"`
    AltitudeUnit string `json:"altitude_unit,omitempty"`
    IlluminanceValue float32 `json:"illuminance_value,omitempty"`
    IlluminanceUnit string `json:"illuminance_unit,omitempty"`
    SoilMoistureValue float32 `json:"soil_moisture_value,omitempty"`
    SoilMoistureUnit string `json:"soil_moisture_unit,omitempty"`
    Timestamp int64 `json:"timestamp,omitempty"`
}

// The abstraction of the `Mikapod Soil` reader.
type ArduinoReader struct {
    serialPort *serial.Port
}

// Constructor used to intialize the serial reader designed to communicate
// with Arduino configured for the `Mikapod Soil` settings.
func ArduinoReaderInit(devicePath string) (* ArduinoReader)  {
    log.Printf("READER: Attempting to connect Arduino device...")
    c := &serial.Config{Name: devicePath, Baud: 9600}
    s, err := serial.OpenPort(c)
    if err != nil {
        log.Fatal(err)
    }

    // DEVELOPERS NOTE:
    // The following code will warm up the Arduino device before we are
    // able to make calls to the external sensors.
    log.Printf("READER: Waiting for Arduino external sensors to warm up")
    ar := &ArduinoReader{serialPort: s}
    ar.Read()
    time.Sleep(5 * time.Second)
    ar.Read()
    time.Sleep(5 * time.Second)
    return ar
}

// Function returns the JSON data of the instrument readings from our Arduino
// device configured for the `Mikapod Soil` settings.
func (ar* ArduinoReader) Read() (*TimeSeriesData) {
    // DEVELOPERS NOTE:
    // (1) The external device (Arduino) is setup to standby idle until it
    //     receives a poll request from this code, once a poll request has
    //     been submitted then all the sensors get polled and their data is
    //     returned.
    // (2) Please look at the following code to understand how the external
    //     device works in:
    //     src/instruments/mikapod-soil-arduino/mikapod-soil-arduino.ino
    // (3) The reason for design is as follows:
    //     (a) The external device does not have a real-time clock
    //     (b) We don't want to add any real-time clock shields because
    //         extra hardware means it costs more.
    //     (c) We don't want to write complicated code of synching time
    //         from this code because it will make the code complicated.
    //     (d) Therefore we chose to make sensor polling be event based
    //         and this code needs to send a "poll request".

    // STEP 1:
    // We need to send a single byte to the external device (Arduino) which
    // will trigger a polling event on all the sensors.
    n, err := ar.serialPort.Write([]byte(RX_BYTE))
    if err != nil {
        log.Fatal(err)
    }

    // STEP 2:
    // The external device will poll the device, we need to make our main
    // runtime loop to be blocked so we wait until the device finishes and
    // returns all the sensor measurements.
    buf := make([]byte, 1028)
    n, err = ar.serialPort.Read(buf)
    if err != nil {
        log.Fatal(err)
    }

    // STEP 3:
    // Check to see if ANY data was returned from the external device, if
    // there was then we load up the string into a JSON object.
    var tsd TimeSeriesData
    err = json.Unmarshal(buf[:n], &tsd)
    if err != nil {
        return nil
    }
    tsd.Timestamp = time.Now().Unix()
    return &tsd
}
