package data

import (
	"math"
	"math/rand"
	"time"
)

type StatusData struct {
	Tyield         float64 `json:"Tyield"`
	Dyield         float64 `json:"Dyield"`
	PF             float64 `json:"PF"`
	Pmax           float64 `json:"Pmax"`
	Pac            float64 `json:"Pac"`
	Sac            float64 `json:"Sac"`
	Uab            float64 `json:"Uab"`
	Ubc            float64 `json:"Ubc"`
	Uca            float64 `json:"Uca"`
	Ia             float64 `json:"Ia"`
	Ib             float64 `json:"Ib"`
	Ic             float64 `json:"Ic"`
	Freq           float64 `json:"Freq"`
	Tmod           float64 `json:"Tmod"`
	Tamb           float64 `json:"Tamb"`
	Mode           string  `json:"Mode"`
	Qac            int     `json:"Qac"`
	BusCapacitance float64 `json:"BusCapacitance"`
	AcCapacitance  float64 `json:"AcCapacitance"`
	Pdc            float64 `json:"Pdc"`
	PmaxLim        float64 `json:"PmaxLim"`
	SmaxLim        float64 `json:"SmaxLim"`
}

type SensorData struct {
	Device    string     `json:"Device"`
	Timestamp string     `json:"Timestamp"`
	ProVer    int        `json:"ProVer"`
	MinorVer  int        `json:"MinorVer"`
	SN        int64      `json:"SN"`
	Model     string     `json:"model"`
	Status    StatusData `json:"Status"`
}

const (
	proVerVal   = 12345
	minorVerVal = 12345
	snVal       = 123456789012345
)

func roundToOneDecimal(value float64) float64 {
	return math.Round(value*10) / 10
}

func generateMode() string {
	r := rand.Float64() * 100
	switch {
	case r < 1:
		return "Fault"
	case r < 3:
		return "Check"
	case r < 13:
		return "Standby"
	case r < 99:
		return "Running"
	default:
		return "Derate"
	}
}

func generateRandomFloat(min, max float64) float64 {
	return roundToOneDecimal(min + rand.Float64()*(max-min))
}

func GenerateData() SensorData {
	return SensorData{
		Device:    "IoT_002",
		Timestamp: time.Now().Format("2006-01-02 15:04:05"),
		ProVer:    proVerVal,
		MinorVer:  minorVerVal,
		SN:        snVal,
		Model:     "mo-2093",
		Status: StatusData{
			Tyield:         generateRandomFloat(0, 5000),
			Dyield:         generateRandomFloat(0, 500),
			PF:             generateRandomFloat(0, 100),
			Pmax:           generateRandomFloat(0.1, 1000),
			Pac:            generateRandomFloat(0.1, 1000),
			Sac:            generateRandomFloat(0.1, 1000),
			Uab:            generateRandomFloat(0.1, 1000),
			Ubc:            generateRandomFloat(0.1, 1000),
			Uca:            generateRandomFloat(0.1, 1000),
			Ia:             generateRandomFloat(0.1, 100),
			Ib:             generateRandomFloat(0.1, 100),
			Ic:             generateRandomFloat(0.1, 100),
			Freq:           generateRandomFloat(0.1, 1000),
			Tmod:           generateRandomFloat(0.1, 1000),
			Tamb:           generateRandomFloat(0.1, 1000),
			Mode:           generateMode(),
			Qac:            rand.Intn(2000),
			BusCapacitance: generateRandomFloat(0, 20),
			AcCapacitance:  generateRandomFloat(0, 20),
			Pdc:            generateRandomFloat(0, 50),
			PmaxLim:        generateRandomFloat(0, 50),
			SmaxLim:        generateRandomFloat(0, 50),
		},
	}
}
