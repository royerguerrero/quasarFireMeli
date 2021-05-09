package satellites

type Satellite struct {
	Name      string   `json:"name"`
	Distances string   `json:"distances"`
	Message   []string `json:"message"`
}

type Satellites [3]Satellite

func (s *Satellites) GetLocation(distances ...float32) (x, y float32) {
	return 0.0, 0.0
}

func (s *Satellites) GetMessage(messages ...[]string) (msg string) {
	return "Test"
}
