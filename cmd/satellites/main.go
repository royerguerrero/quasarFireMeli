package satellites

import (
	"sort"
)

type Satellite struct {
	Name      string   `json:"name"`
	Distances string   `json:"distances"`
	Message   []string `json:"message"`
}

type Satellites [3]Satellite

func GetLocation(distances ...float32) (x, y float32) {
	return 0.0, 0.0
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

	var message string

	for _, k := range keys {
		message += decryptedMessage[k] + " "
	}

	return message
}
