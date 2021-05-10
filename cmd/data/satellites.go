package data

import (
	"math"
	"sort"
	"strings"
)

type Satellite struct {
	Name     string   `json:"name"`
	Distance float64  `json:"distance"`
	Message  []string `json:"message"`
}

type SatellitesRepo struct {
	Satellites [3]Satellite `json:"satellites"`
}

func (s SatellitesRepo) FindByName(name string) Satellite {
	for _, satellite := range s.Satellites {
		if strings.EqualFold(name, satellite.Name) {
			return satellite
		}
	}
	return Satellite{}
}

func (s SatellitesRepo) getSatellitesOnline() map[string][2]float64 {
	onlineSatellites := make(map[string][2]float64)
	onlineSatellites["Kenobi"] = [2]float64{-500.0, -200.0}
	onlineSatellites["Skywalker"] = [2]float64{100.0, -100.0}
	onlineSatellites["Sato"] = [2]float64{500.0, 100.0}

	return onlineSatellites
}

var Satellites = SatellitesRepo{}

// Obtains the position of the emisor.
func GetLocation(distances ...float64) (x, y float64) {

	satellites := Satellites.getSatellitesOnline()

	x1, y1 := satellites["Kenobi"][0], satellites["Kenobi"][1]
	x2, y2 := satellites["Skywalker"][0], satellites["Skywalker"][1]
	x3, y3 := satellites["Sato"][0], satellites["Sato"][1]

	r1 := distances[0]
	r2 := distances[1]
	r3 := distances[2]

	A := 2*x2 - 2*x1
	B := 2*y2 - 2*y1
	C := math.Pow(r1, 2) - math.Pow(r2, 2) - math.Pow(x1, 2) + math.Pow(x2, 2) - math.Pow(y1, 2) + math.Pow(y2, 2)
	D := 2*x3 - 2*x2
	E := 2*y3 - 2*y2
	F := math.Pow(r2, 2) - math.Pow(r3, 2) - math.Pow(x2, 2) + math.Pow(x3, 2) - math.Pow(y2, 2) + math.Pow(y3, 2)
	x = (C*E - F*B) / (E*A - B*D)
	y = (C*D - A*F) / (B*D - A*E)

	return x, y
}

// Deletes message noise
func GetMessage(messages ...[]string) (msg string) {
	decryptedMessage := make(map[int]string)

	for i := range messages {
		for j, word := range messages[i] {
			if word != "" {
				decryptedMessage[j] = word
			}
		}
	}

	keys := []int{}
	for k := range decryptedMessage {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	for _, k := range keys {
		msg += decryptedMessage[k] + " "
	}

	return strings.Trim(msg, " ")
}
