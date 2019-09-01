package internal // github.com/mikaponics/mikapod-soil-reader/internal

import (
	"context"
	"fmt"

	"github.com/golang/protobuf/ptypes"

	pb "github.com/mikaponics/mikapod-soil-reader/api"
)

type MikapodSoilReaderGRPC struct{
	arduinoReader *ArduinoReader
}

func (s *MikapodSoilReaderGRPC) GetData(ctx context.Context, in *pb.PollSensorsRequest) (*pb.SensorsDataResponse, error) {

	datum := s.arduinoReader.Read()
	fmt.Printf("%+v\n", datum)
	// fmt.Printf("%+v\n", s.arduinoReader)

	return &pb.SensorsDataResponse{
		Status: true,
		Timestamp: ptypes.TimestampNow(), // Note: https://godoc.org/github.com/golang/protobuf/ptypes#Timestamp
		// HumidityValue: datum.HumidityValue,
		// HumidityUnit: datum.HumidityUnit,
		// TemperatureValue: datum.TemperatureValue,
		// TemperatureUnit: datum.TemperatureUnit,
	}, nil
}


// float  = 5;
// float  = 6;
// float pressureValue = 7;
// string pressureUnit = 8;
// float temperatureBackupValue = 9;
// string temperatureBackupUnit = 10;
// float altitudeValue = 11;
// string altitudeUnit = 12;
// float illuminanceValue = 13;
// string illuminanceUnit = 14;
// float soilMoistureValue = 15;
// string soilMoistureUnit = 16;
