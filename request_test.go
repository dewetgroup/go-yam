package yam

import (
	"context"
	"net/http"
	"net/url"
	"testing"
)

func CreateMockHTTPRequest(method Method, path string) *http.Request {
	request := &http.Request{
		Method: string(method),
		URL:    &url.URL{Path: path},
	}
	return request.WithContext(context.WithValue(context.Background(), "key", "value"))
}

func TestRequestHasCorrectPath(t *testing.T) {
	httpRequest := CreateMockHTTPRequest(GET, "/mock/path")
	request := fromHttpRequest(httpRequest)
	if request.URL.Path != "/mock/path" {
		t.Errorf("Unexpected request URL '%s'", request.URL.Path)
	}
}

func TestRequestCanAccessContext(t *testing.T) {
	httpRequest := CreateMockHTTPRequest(GET, "/mock/path")
	request := fromHttpRequest(httpRequest)
	context := request.Context()
	value := context.Value("key")
	if !(value.(string) == "value") {
		t.Errorf("Expected context to contain key 'key' with value 'value', found '%s'", value.(string))
	}
}
