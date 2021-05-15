package errorutil

type errorBody struct {
	Error   string `json:"error,omitempty"`
	Message string `json:"message,omitempty"`
	Details string `json:"details,omitempty"`
	ErrorCode
}
