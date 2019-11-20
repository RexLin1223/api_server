package protocol

// TokenResponse is format response to request client
type TokenResponse struct {
	Result  string `json:"result"`
	Message string `json:"message"`
	Token   string `json:"token"`
}
