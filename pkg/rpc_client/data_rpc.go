package rpc_client

import (
	// "time"
	// "errors"
	// "database/sql"
	"log"
	// "net/rpc"
)

// Request structure to send to the RPC server.

type GetDataRequest struct {}

// The time-series data structure used to store all the data that will be
// returned by the `Mikapod Soil` Arduino device.
type GetDataResponse struct {
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

// RPC Calls

func (s *MikapodSoilReaderService) GetData() (*GetDataResponse, error) {
	var request *GetDataRequest
	var response *GetDataResponse
	rpcErr := s.Client.Call("RPC.GetData", &request, &response)
	if rpcErr != nil {
		log.Println("ERROR | MikapodSoilReaderService | GetData | rpcErr:", rpcErr)
		return nil, rpcErr
	}
	return response, nil
}
