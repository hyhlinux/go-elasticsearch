// generated by github.com/elastic/go-elasticsearch/cmd/generator; DO NOT EDIT

package api

import (
	"net/http"

	"github.com/elastic/go-elasticsearch/transport"
	"github.com/elastic/go-elasticsearch/util"
)

// PutScriptOption is a non-required PutScript option that gets applied to an HTTP request.
type PutScriptOption func(r *transport.Request)

// WithPutScriptErrorTrace - include the stack trace of returned errors.
func WithPutScriptErrorTrace(errorTrace bool) PutScriptOption {
	return func(r *transport.Request) {
	}
}

// WithPutScriptFilterPath - a comma-separated list of filters used to reduce the respone.
func WithPutScriptFilterPath(filterPath []string) PutScriptOption {
	return func(r *transport.Request) {
	}
}

// WithPutScriptHuman - return human readable values for statistics.
func WithPutScriptHuman(human bool) PutScriptOption {
	return func(r *transport.Request) {
	}
}

// WithPutScriptIgnore - ignores the specified HTTP status codes.
func WithPutScriptIgnore(ignore []int) PutScriptOption {
	return func(r *transport.Request) {
	}
}

// WithPutScriptPretty - pretty format the returned JSON response.
func WithPutScriptPretty(pretty bool) PutScriptOption {
	return func(r *transport.Request) {
	}
}

// WithPutScriptSourceParam - the URL-encoded request definition. Useful for libraries that do not accept a request body for non-POST requests.
func WithPutScriptSourceParam(sourceParam string) PutScriptOption {
	return func(r *transport.Request) {
	}
}

// PutScript - the scripting module enables you to use scripts to evaluate custom expressions. See https://www.elastic.co/guide/en/elasticsearch/reference/5.x/modules-scripting.html for more info.
//
// lang: script language.
//
// body: the document.
//
// options: optional parameters.
func (a *API) PutScript(lang string, body map[string]interface{}, options ...PutScriptOption) (*PutScriptResponse, error) {
	req := a.transport.NewRequest("PUT")
	for _, option := range options {
		option(req)
	}
	resp, err := a.transport.Do(req)
	return &PutScriptResponse{resp}, err
}

// PutScriptResponse is the response for PutScript.
type PutScriptResponse struct {
	Response *http.Response
	// TODO: fill in structured response
}

// DecodeBody decodes the JSON body of the HTTP response.
func (r *PutScriptResponse) DecodeBody() (util.MapStr, error) {
	return transport.DecodeResponseBody(r.Response)
}
