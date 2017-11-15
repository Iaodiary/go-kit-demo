package adendpoint

import (
	"github.com/go-kit/kit/endpoint"
)

// Set collects all endpoints that compose an ad service. It's meant to
// be used as a helper struct, to collect all of the endpoints into a single
// parameter.
type Set struct {
	GetAdEndpoint endpoint.Endpoint
}

func New(svc adservice.Service) {

}
