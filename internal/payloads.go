package internal

type TopSecretResponse struct {
	Position map[string]float64 `json:"position"`
	Message  string             `json:"message"`
}

type ErrorResponse struct {
	Detail interface{} `json:"detail"`
}
