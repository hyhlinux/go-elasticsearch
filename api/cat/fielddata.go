// generated by github.com/elastic/go-elasticsearch/cmd/generator; DO NOT EDIT

package cat

import (
	"net/http"
	"time"

	"github.com/elastic/go-elasticsearch/transport"
	"github.com/elastic/go-elasticsearch/util"
)

// FielddataOption is a non-required Fielddata option that gets applied to an HTTP request.
type FielddataOption func(r *transport.Request)

// WithFielddataFields - a comma-separated list of fields to return the fielddata size.
func WithFielddataFields(fields []string) FielddataOption {
	return func(r *transport.Request) {
	}
}

// FielddataBytes - the unit in which to display byte values.
type FielddataBytes int

const (
	// FielddataBytesB can be used to set FielddataBytes to "b"
	FielddataBytesB = iota
	// FielddataBytesK can be used to set FielddataBytes to "k"
	FielddataBytesK = iota
	// FielddataBytesKb can be used to set FielddataBytes to "kb"
	FielddataBytesKb = iota
	// FielddataBytesM can be used to set FielddataBytes to "m"
	FielddataBytesM = iota
	// FielddataBytesMb can be used to set FielddataBytes to "mb"
	FielddataBytesMb = iota
	// FielddataBytesG can be used to set FielddataBytes to "g"
	FielddataBytesG = iota
	// FielddataBytesGb can be used to set FielddataBytes to "gb"
	FielddataBytesGb = iota
	// FielddataBytesT can be used to set FielddataBytes to "t"
	FielddataBytesT = iota
	// FielddataBytesTb can be used to set FielddataBytes to "tb"
	FielddataBytesTb = iota
	// FielddataBytesP can be used to set FielddataBytes to "p"
	FielddataBytesP = iota
	// FielddataBytesPb can be used to set FielddataBytes to "pb"
	FielddataBytesPb = iota
)

// WithFielddataBytes - the unit in which to display byte values.
func WithFielddataBytes(bytes FielddataBytes) FielddataOption {
	return func(r *transport.Request) {
	}
}

// WithFielddataFieldsParam - a comma-separated list of fields to return in the output.
func WithFielddataFieldsParam(fieldsParam []string) FielddataOption {
	return func(r *transport.Request) {
	}
}

// WithFielddataFormat - a short version of the Accept header, e.g. json, yaml.
func WithFielddataFormat(format string) FielddataOption {
	return func(r *transport.Request) {
	}
}

// WithFielddataH - comma-separated list of column names to display.
func WithFielddataH(h []string) FielddataOption {
	return func(r *transport.Request) {
	}
}

// WithFielddataHelp - return help information.
func WithFielddataHelp(help bool) FielddataOption {
	return func(r *transport.Request) {
	}
}

// WithFielddataLocal - return local information, do not retrieve the state from master node (default: false).
func WithFielddataLocal(local bool) FielddataOption {
	return func(r *transport.Request) {
	}
}

// WithFielddataMasterTimeout - explicit operation timeout for connection to master node.
func WithFielddataMasterTimeout(masterTimeout time.Duration) FielddataOption {
	return func(r *transport.Request) {
	}
}

// WithFielddataS - comma-separated list of column names or column aliases to sort by.
func WithFielddataS(s []string) FielddataOption {
	return func(r *transport.Request) {
	}
}

// WithFielddataV - verbose mode. Display column headers.
func WithFielddataV(v bool) FielddataOption {
	return func(r *transport.Request) {
	}
}

// WithFielddataErrorTrace - include the stack trace of returned errors.
func WithFielddataErrorTrace(errorTrace bool) FielddataOption {
	return func(r *transport.Request) {
	}
}

// WithFielddataFilterPath - a comma-separated list of filters used to reduce the respone.
func WithFielddataFilterPath(filterPath []string) FielddataOption {
	return func(r *transport.Request) {
	}
}

// WithFielddataHuman - return human readable values for statistics.
func WithFielddataHuman(human bool) FielddataOption {
	return func(r *transport.Request) {
	}
}

// WithFielddataIgnore - ignores the specified HTTP status codes.
func WithFielddataIgnore(ignore []int) FielddataOption {
	return func(r *transport.Request) {
	}
}

// WithFielddataPretty - pretty format the returned JSON response.
func WithFielddataPretty(pretty bool) FielddataOption {
	return func(r *transport.Request) {
	}
}

// WithFielddataSourceParam - the URL-encoded request definition. Useful for libraries that do not accept a request body for non-POST requests.
func WithFielddataSourceParam(sourceParam string) FielddataOption {
	return func(r *transport.Request) {
	}
}

// Fielddata - see https://www.elastic.co/guide/en/elasticsearch/reference/5.x/cat-fielddata.html for more info.
//
// options: optional parameters.
func (c *Cat) Fielddata(options ...FielddataOption) (*FielddataResponse, error) {
	req := c.transport.NewRequest("GET")
	for _, option := range options {
		option(req)
	}
	resp, err := c.transport.Do(req)
	return &FielddataResponse{resp}, err
}

// FielddataResponse is the response for Fielddata.
type FielddataResponse struct {
	Response *http.Response
	// TODO: fill in structured response
}

// DecodeBody decodes the JSON body of the HTTP response.
func (r *FielddataResponse) DecodeBody() (util.MapStr, error) {
	return transport.DecodeResponseBody(r.Response)
}
