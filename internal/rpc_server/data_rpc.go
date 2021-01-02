package rpc_server

import (
	"log"

	c "github.com/mikaponics/mikapod-soil-reader/pkg/rpc_client"
)

func (rpc *RPC) GetData(request *c.GetDataRequest, response *c.GetDataResponse) (error) {
    log.Printf("Reading")
	datum := rpc.ArduinoReader.Read()

	log.Printf("Polled: \n%+v\n", datum)

	*response = c.GetDataResponse{
		Status: "OK",
		Timestamp: datum.Timestamp,
		HumidityValue: datum.HumidityValue,
		HumidityUnit: datum.HumidityUnit,
		TemperatureValue: datum.TemperatureValue,
		TemperatureUnit: datum.TemperatureUnit,
		PressureValue: datum.PressureValue,
		PressureUnit: datum.PressureUnit,
		TemperatureBackupValue: datum.TemperatureBackupValue,
		TemperatureBackupUnit: datum.TemperatureBackupUnit,
		AltitudeValue: datum.AltitudeValue,
		AltitudeUnit: datum.AltitudeUnit,
		IlluminanceValue: datum.IlluminanceValue,
		IlluminanceUnit: datum.IlluminanceUnit,
		SoilMoistureValue: datum.SoilMoistureValue,
		SoilMoistureUnit: datum.SoilMoistureUnit,
	}
	return nil //err
}
