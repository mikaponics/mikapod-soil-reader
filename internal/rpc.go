package internal // github.com/mikaponics/mikapod-soil-reader/internal

import (
	"context"

	"github.com/golang/protobuf/ptypes"

	pb "github.com/mikaponics/mikapod-soil-reader/api"
)

type MikapodSoilReaderGRPC struct{
	arduinoReader *ArduinoReader
}

func (s *MikapodSoilReaderGRPC) GetData(ctx context.Context, in *pb.GetTimeSeriesData) (*pb.TimeSeriesDataResponse, error) {
	datum := s.arduinoReader.Read()
	return &pb.TimeSeriesDataResponse{
		Status: true,
		Timestamp: ptypes.TimestampNow(), // Note: https://godoc.org/github.com/golang/protobuf/ptypes#Timestamp
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
	}, nil
}
