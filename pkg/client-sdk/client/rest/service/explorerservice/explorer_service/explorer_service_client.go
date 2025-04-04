// Code generated by go-swagger; DO NOT EDIT.

package explorer_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// New creates a new explorer service API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

// New creates a new explorer service API client with basic auth credentials.
// It takes the following parameters:
// - host: http host (github.com).
// - basePath: any base path for the API client ("/v1", "/v3").
// - scheme: http scheme ("http", "https").
// - user: user for basic authentication header.
// - password: password for basic authentication header.
func NewClientWithBasicAuth(host, basePath, scheme, user, password string) ClientService {
	transport := httptransport.New(host, basePath, []string{scheme})
	transport.DefaultAuthentication = httptransport.BasicAuth(user, password)
	return &Client{transport: transport, formats: strfmt.Default}
}

// New creates a new explorer service API client with a bearer token for authentication.
// It takes the following parameters:
// - host: http host (github.com).
// - basePath: any base path for the API client ("/v1", "/v3").
// - scheme: http scheme ("http", "https").
// - bearerToken: bearer token for Bearer authentication header.
func NewClientWithBearerToken(host, basePath, scheme, bearerToken string) ClientService {
	transport := httptransport.New(host, basePath, []string{scheme})
	transport.DefaultAuthentication = httptransport.BearerToken(bearerToken)
	return &Client{transport: transport, formats: strfmt.Default}
}

/*
Client for explorer service API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption may be used to customize the behavior of Client methods.
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	ExplorerServiceGetRound(params *ExplorerServiceGetRoundParams, opts ...ClientOption) (*ExplorerServiceGetRoundOK, error)

	ExplorerServiceGetRoundByID(params *ExplorerServiceGetRoundByIDParams, opts ...ClientOption) (*ExplorerServiceGetRoundByIDOK, error)

	ExplorerServiceListVtxos(params *ExplorerServiceListVtxosParams, opts ...ClientOption) (*ExplorerServiceListVtxosOK, error)

	ExplorerServiceSubscribeForAddress(params *ExplorerServiceSubscribeForAddressParams, opts ...ClientOption) (*ExplorerServiceSubscribeForAddressOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
ExplorerServiceGetRound explorer service get round API
*/
func (a *Client) ExplorerServiceGetRound(params *ExplorerServiceGetRoundParams, opts ...ClientOption) (*ExplorerServiceGetRoundOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExplorerServiceGetRoundParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ExplorerService_GetRound",
		Method:             "GET",
		PathPattern:        "/v1/round/{txid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ExplorerServiceGetRoundReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ExplorerServiceGetRoundOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ExplorerServiceGetRoundDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ExplorerServiceGetRoundByID explorer service get round by Id API
*/
func (a *Client) ExplorerServiceGetRoundByID(params *ExplorerServiceGetRoundByIDParams, opts ...ClientOption) (*ExplorerServiceGetRoundByIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExplorerServiceGetRoundByIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ExplorerService_GetRoundById",
		Method:             "GET",
		PathPattern:        "/v1/round/id/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ExplorerServiceGetRoundByIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ExplorerServiceGetRoundByIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ExplorerServiceGetRoundByIDDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ExplorerServiceListVtxos explorer service list vtxos API
*/
func (a *Client) ExplorerServiceListVtxos(params *ExplorerServiceListVtxosParams, opts ...ClientOption) (*ExplorerServiceListVtxosOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExplorerServiceListVtxosParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ExplorerService_ListVtxos",
		Method:             "GET",
		PathPattern:        "/v1/vtxos/{address}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ExplorerServiceListVtxosReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ExplorerServiceListVtxosOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ExplorerServiceListVtxosDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ExplorerServiceSubscribeForAddress explorer service subscribe for address API
*/
func (a *Client) ExplorerServiceSubscribeForAddress(params *ExplorerServiceSubscribeForAddressParams, opts ...ClientOption) (*ExplorerServiceSubscribeForAddressOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExplorerServiceSubscribeForAddressParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ExplorerService_SubscribeForAddress",
		Method:             "GET",
		PathPattern:        "/v1/vtxos/{address}/subscribe",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ExplorerServiceSubscribeForAddressReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ExplorerServiceSubscribeForAddressOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ExplorerServiceSubscribeForAddressDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
