package req

// Request define requests in .req.json
type Request struct {
	Name    string      `json:"name"`
	URL     string      `json:"url"`
	Method  string      `json:"method"`
	Headers []string    `json:"headers"`
	Body    interface{} `json:"body"`
}
