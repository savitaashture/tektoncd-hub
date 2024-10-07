// Code generated by goa v3.19.1, DO NOT EDIT.
//
// status service
//
// Command:
// $ goa gen github.com/tektoncd/hub/api/design

package status

import (
	"context"
)

// Describes the status of each service
type Service interface {
	// Return status of the services
	Status(context.Context) (res *StatusResult, err error)
}

// APIName is the name of the API as defined in the design.
const APIName = "hub"

// APIVersion is the version of the API as defined in the design.
const APIVersion = "1.0"

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "status"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [1]string{"Status"}

// Describes the services and their status
type HubService struct {
	// Name of the service
	Name string
	// Status of the service
	Status string
	// Details of the error if any
	Error *string
}

// StatusResult is the result type of the status service Status method.
type StatusResult struct {
	// List of services and their status
	Services []*HubService
}