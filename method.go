package yam

type Method string

// Should you require a method not listed here,
// you can simply provide the string value when routing.

const (
	GET     Method = "GET"
	POST    Method = "POST"
	DELETE  Method = "DELETE"
	PATCH   Method = "PATCH"
	PUT     Method = "PUT"
	HEAD    Method = "HEAD"
	OPTIONS Method = "OPTIONS"
	CONNECT Method = "CONNECT"
	TRACE   Method = "TRACE"
	LOCK    Method = "LOCK"
	UNLOCK  Method = "UNLOCK"
)
