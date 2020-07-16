package merror

// ErrorMessage is being used across all controllers
type ErrorMessage struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}
