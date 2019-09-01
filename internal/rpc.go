package internal // github.com/mikaponics/mikapod-soil-reader/internal

import (
	"context"

	pb "github.com/mikaponics/mikapod-soil-reader/api"
)

type MikapodSoilReaderGRPC struct{
	arduinoReader *ArduinoReader
}

func (s *MikapodSoilReaderGRPC) GetData(ctx context.Context, in *pb.PollSensorsRequest) (*pb.SensorsDataResponse, error) {
	return &pb.SensorsDataResponse{
		Status: true,
	}, nil
}
