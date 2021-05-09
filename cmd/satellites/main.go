package satellites

import (
	"math"
	"sort"
)

type Satellite struct {
	Name      string   `json:"name"`
	Distances string   `json:"distances"`
	Message   []string `json:"message"`
}

type Satellites [3]Satellite

// Obtains the position of the emisor.
func GetLocation(distances ...float64) (x, y float64) {
	x1, y1 := -500.0, -200.0
	x2, y2 := 100.0, -100.0
	x3, y3 := 500.0, 100.0

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

	return msg
}
