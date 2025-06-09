package nowpaymentsresponse

type InvalidRequestPayResponse struct {
	Status     bool   `json:"status"`
	StatusCode int    `json:"statusCode"`
	Code       string `json:"code"`
	Message    string `json:"message"`
	Error      string `json:"error"`
}
