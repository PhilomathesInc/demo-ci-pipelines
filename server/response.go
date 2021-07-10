package server

// Response is a string response the server returns
type Response string

const (
	// Server health is healthy
	Healthy Response = `{"status":"ok"}`
	// Server health is unhealthy
	UnHealthy Response = `{"status":"not ok"}`
	// Server welcome message
	WelcomeMsg Response = "welcome to the demo go app"
)
