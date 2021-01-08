package constants

// The local server configuration constants for this service.
const (
	MikapodSoilReaderServicePort = ":50052"             // Please do not change!
	MikapodSoilReaderServiceAddress = "localhost:50052" // Please do not change!
)

// The unique identifiers used to describe the physical sensors in our device.
const (
	MikapodSoilReaderHumidityInstrumentId int32 = 1
	MikapodSoilReaderTemperatureInstrumentId int32 = 2
	MikapodSoilReaderPressureInstrumentId int32 = 3
	MikapodSoilReaderTemperatureBackupInstrumentId int32 = 4
	MikapodSoilReaderAltitudeInstrumentId int32 = 5
	MikapodSoilReaderIlluminanceInstrumentId int32 = 6
	MikapodSoilReaderSoilMoistureInstrumentId int32 = 7
)
