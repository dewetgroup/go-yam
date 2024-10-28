package yam

import (
	"context"
	"net/http"
	"net/url"
)

// Response is a type alias for http.Response.
type Request struct {
	request *http.Request
	URL     url.URL
	Query   url.Values
}

func fromHttpRequest(request *http.Request) Request {
	return Request{
		request: request,
		URL:     *request.URL,
	}
}

func (r *Request) Context() context.Context {
	return r.request.Context()
}

func (r *Request) PathValue(key string) string {
	return r.request.PathValue(key)
}
