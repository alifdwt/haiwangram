package responses

type ErrorMessage struct {
	// StatusCode int    `json:"statusCode"`
	Message string `json:"message"`
	// Error      bool   `json:"error"`
}
