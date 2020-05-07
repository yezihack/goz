package goz

import "time"

// Options object
type Options struct {
	BaseURI    string
	Timeout    time.Duration
	timeout    time.Duration
	Query      interface{}
	Headers    map[string]interface{}
	Cookies    interface{}
	FormParams map[string]interface{}
	JSON       interface{}
	Proxy      string
}
