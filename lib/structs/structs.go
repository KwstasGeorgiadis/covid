package structs

// ErrorMessage is being used across all controllers
type ErrorMessage struct {
	ErrorMessage string `json:"message"`
	Code         int    `json:"code"`
}
