package yam

import (
	"errors"
	"fmt"
	"net/http"
)

// Handler is any function that handles an HTTP Request and returns an (Response, error) tuple.
type Handler = func(Request) (Response, error)

// The default handler, simply returning not found.
func DefaultHandler(Request) (Response, error) {
	return NotFound{}, nil
}

// This handler can be used as a placeholder for not-implemented endpoints
func NotImplementedHandler(Request) (Response, error) {
	return nil, errors.New("not implemented")
}

// adapt adapts a handler to the http.HandlerFunc interface by ensuring the the return response is written to the http.ResponseWriter.
// It recovers from panics and any errors returned by the handler.
// Both cases result in an 500 response.
func adapt(handler Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				InternalServerError{fmt.Errorf("internal server error")}.Write(writer)
			}
		}()
		response, err := handler(fromHttpRequest(request))
		if err != nil {
			response = InternalServerError{fmt.Errorf("internal server error")}
		}
		response.Write(writer)
	})
}
