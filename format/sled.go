package format

import (
	"math"
)

type FMSLED struct {
	IsRaceOn    bool
	TimestampMS uint32

	EngineMaxRpm     float32
	EngineIdleRpm    float32
	CurrentEngineRpm float32

	AccelerationX float32
	AccelerationY float32
	AccelerationZ float32

	VelocityX float32
	VelocityY float32
	VelocityZ float32

	AngularVelocityX float32
	AngularVelocityY float32
	AngularVelocityZ float32

	Yaw   float32
	Pitch float32
	Roll  float32

	NormalizedSuspensionFL float32
	NormalizedSuspensionFR float32
	NormalizedSuspensionRL float32
	NormalizedSuspensionRR float32

	TireSlipRatioFL float32
	TireSlipRatioFR float32
	TireSlipRatioRL float32
	TireSlipRatioRR float32

	WheelRotationSpeedFL float32
	WheelRotationSpeedFR float32
	WheelRotationSpeedRL float32
	WheelRotationSpeedRR float32

	WheelOnRumbleStripFL bool
	WheelOnRumbleStripFR bool
	WheelOnRumbleStripRL bool
	WheelOnRumbleStripRR bool

	WheelInPuddleDepthFL float32
	WheelInPuddleDepthFR float32
	WheelInPuddleDepthRL float32
	WheelInPuddleDepthRR float32

	SurfaceRumbleFL float32
	SurfaceRumbleFR float32
	SurfaceRumbleRL float32
	SurfaceRumbleRR float32

	TireSlipAngleFL float32
	TireSlipAngleFR float32
	TireSlipAngleRL float32
	TireSlipAngleRR float32

	TireCombinedSlipFL float32
	TireCombinedSlipFR float32
	TireCombinedSlipRL float32
	TireCombinedSlipRR float32

	SuspensionTravelMetersFL float32
	SuspensionTravelMetersFR float32
	SuspensionTravelMetersRL float32
	SuspensionTravelMetersRR float32

	CarOrdinal          int32
	CarClass            int32
	CarPerformanceIndex int32
	DrivetrainType      int32
	NumCylinders        int32
}

type FMSLEDSerializer struct{}

func NewSLEDSerializer() *FMSLEDSerializer {
	return &FMSLEDSerializer{}
}

func (s *FMSLEDSerializer) Parse(data []byte, len int) (interface{}, error) {
	if len != 232 {
		return nil, ErrorInvalidLength
	}

	sled := FMSLED{
		IsRaceOn: data[0] == 1,
		// data[1:4]? padding?
		TimestampMS: Uint32AtOffset(data, 4),

		EngineMaxRpm:     math.Float32frombits(Uint32AtOffset(data, 8)),
		EngineIdleRpm:    math.Float32frombits(Uint32AtOffset(data, 12)),
		CurrentEngineRpm: math.Float32frombits(Uint32AtOffset(data, 16)),

		AccelerationX: math.Float32frombits(Uint32AtOffset(data, 20)),
		AccelerationY: math.Float32frombits(Uint32AtOffset(data, 24)),
		AccelerationZ: math.Float32frombits(Uint32AtOffset(data, 28)),

		VelocityX: math.Float32frombits(Uint32AtOffset(data, 32)),
		VelocityY: math.Float32frombits(Uint32AtOffset(data, 36)),
		VelocityZ: math.Float32frombits(Uint32AtOffset(data, 40)),

		AngularVelocityX: math.Float32frombits(Uint32AtOffset(data, 44)),
		AngularVelocityY: math.Float32frombits(Uint32AtOffset(data, 48)),
		AngularVelocityZ: math.Float32frombits(Uint32AtOffset(data, 52)),

		Yaw:   math.Float32frombits(Uint32AtOffset(data, 56)),
		Pitch: math.Float32frombits(Uint32AtOffset(data, 60)),
		Roll:  math.Float32frombits(Uint32AtOffset(data, 64)),

		NormalizedSuspensionFL: math.Float32frombits(Uint32AtOffset(data, 68)),
		NormalizedSuspensionFR: math.Float32frombits(Uint32AtOffset(data, 72)),
		NormalizedSuspensionRL: math.Float32frombits(Uint32AtOffset(data, 76)),
		NormalizedSuspensionRR: math.Float32frombits(Uint32AtOffset(data, 80)),

		TireSlipRatioFL: math.Float32frombits(Uint32AtOffset(data, 84)),
		TireSlipRatioFR: math.Float32frombits(Uint32AtOffset(data, 88)),
		TireSlipRatioRL: math.Float32frombits(Uint32AtOffset(data, 92)),
		TireSlipRatioRR: math.Float32frombits(Uint32AtOffset(data, 96)),

		WheelRotationSpeedFL: math.Float32frombits(Uint32AtOffset(data, 100)),
		WheelRotationSpeedFR: math.Float32frombits(Uint32AtOffset(data, 104)),
		WheelRotationSpeedRL: math.Float32frombits(Uint32AtOffset(data, 108)),
		WheelRotationSpeedRR: math.Float32frombits(Uint32AtOffset(data, 112)),

		WheelOnRumbleStripFL: data[116] == 1,
		WheelOnRumbleStripFR: data[120] == 1,
		WheelOnRumbleStripRL: data[124] == 1,
		WheelOnRumbleStripRR: data[128] == 1,

		WheelInPuddleDepthFL: math.Float32frombits(Uint32AtOffset(data, 132)),
		WheelInPuddleDepthFR: math.Float32frombits(Uint32AtOffset(data, 136)),
		WheelInPuddleDepthRL: math.Float32frombits(Uint32AtOffset(data, 140)),
		WheelInPuddleDepthRR: math.Float32frombits(Uint32AtOffset(data, 144)),

		SurfaceRumbleFL: math.Float32frombits(Uint32AtOffset(data, 148)),
		SurfaceRumbleFR: math.Float32frombits(Uint32AtOffset(data, 152)),
		SurfaceRumbleRL: math.Float32frombits(Uint32AtOffset(data, 156)),
		SurfaceRumbleRR: math.Float32frombits(Uint32AtOffset(data, 160)),

		TireSlipAngleFL: math.Float32frombits(Uint32AtOffset(data, 164)),
		TireSlipAngleFR: math.Float32frombits(Uint32AtOffset(data, 168)),
		TireSlipAngleRL: math.Float32frombits(Uint32AtOffset(data, 172)),
		TireSlipAngleRR: math.Float32frombits(Uint32AtOffset(data, 176)),

		TireCombinedSlipFL: math.Float32frombits(Uint32AtOffset(data, 180)),
		TireCombinedSlipFR: math.Float32frombits(Uint32AtOffset(data, 184)),
		TireCombinedSlipRL: math.Float32frombits(Uint32AtOffset(data, 188)),
		TireCombinedSlipRR: math.Float32frombits(Uint32AtOffset(data, 192)),

		SuspensionTravelMetersFL: math.Float32frombits(Uint32AtOffset(data, 196)),
		SuspensionTravelMetersFR: math.Float32frombits(Uint32AtOffset(data, 200)),
		SuspensionTravelMetersRL: math.Float32frombits(Uint32AtOffset(data, 204)),
		SuspensionTravelMetersRR: math.Float32frombits(Uint32AtOffset(data, 208)),

		CarOrdinal:          int32(Uint32AtOffset(data, 212)),
		CarClass:            int32(Uint32AtOffset(data, 216)),
		CarPerformanceIndex: int32(Uint32AtOffset(data, 220)),
		DrivetrainType:      int32(Uint32AtOffset(data, 224)),
		NumCylinders:        int32(Uint32AtOffset(data, 228)),
	}

	return &sled, nil
}
