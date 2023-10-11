package format

import (
	"fmt"
)

type FMDash struct {
	IsRaceOn bool

	PositionX float64
	PositionY float64
	PositionZ float64

	Speed  float64
	Power  float64
	Torque float64

	TireTempFL float64
	TireTempFR float64
	TireTempRL float64
	TireTempRR float64

	Boost    float64
	Fuel     float64
	Distance float64

	BestLapTime     float64
	LastLapTime     float64
	CurrentLapTime  float64
	CurrentRaceTime float64

	Lap          uint64
	RacePosition uint64

	Accelerator uint64
	Brake       uint64
	Clutch      uint64
	Handbrake   uint64

	Gear  uint64
	Steer int64

	NormalDrivingLine uint64
	NormalAIBrakeDiff uint64
}

type FMDashSerializer struct{}

func NewDashSerializer() *FMDashSerializer {
	return &FMDashSerializer{}
}

func (s *FMDashSerializer) Parse(data []byte, len int) (interface{}, error) {
	if len != 331 {
		return nil, ErrorInvalidLength
	}

	dash := FMDash{}

	dash.IsRaceOn = data[0] == 1

	fmt.Printf("%08b\n", data[1:8])
	// fmt.Printf("%8d\n", data[5])

	// dash.PositionX = float64(int32(data[0]) | int32(data[1])<<8 | int32(data[2])<<16 | int32(data[3])<<24)
	// dash.PositionY = float64(int32(data[4]) | int32(data[5])<<8 | int32(data[6])<<16 | int32(data[7])<<24)
	// dash.PositionZ = float64(int32(data[8]) | int32(data[9])<<8 | int32(data[10])<<16 | int32(data[11])<<24)
	// fmt.Printf("%08b\n", data[1:25])

	// fmt.Printf("Position: (%f, %f, %f)\n", dash.PositionX, dash.PositionY, dash.PositionZ)

	return &dash, nil
}
