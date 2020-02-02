// Code generated by go-swagger; DO NOT EDIT.

package common

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new common API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for common API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	GetSelfServiceBrowserLoginRequest(params *GetSelfServiceBrowserLoginRequestParams) (*GetSelfServiceBrowserLoginRequestOK, error)

	GetSelfServiceBrowserProfileManagementRequest(params *GetSelfServiceBrowserProfileManagementRequestParams) (*GetSelfServiceBrowserProfileManagementRequestOK, error)

	GetSelfServiceBrowserRegistrationRequest(params *GetSelfServiceBrowserRegistrationRequestParams) (*GetSelfServiceBrowserRegistrationRequestOK, error)

	GetSelfServiceError(params *GetSelfServiceErrorParams) (*GetSelfServiceErrorOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetSelfServiceBrowserLoginRequest gets the request context of browser based login user flows

  This endpoint returns a login request's context with, for example, error details and
other information.

When accessing this endpoint through ORY Kratos' Public API, ensure that cookies are set as they are required for CSRF to work. To prevent
token scanning attacks, the public endpoint does not return 404 status codes to prevent scanning attacks.

More information can be found at [ORY Kratos User Login and User Registration Documentation](https://www.ory.sh/docs/next/kratos/self-service/flows/user-login-user-registration).
*/
func (a *Client) GetSelfServiceBrowserLoginRequest(params *GetSelfServiceBrowserLoginRequestParams) (*GetSelfServiceBrowserLoginRequestOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSelfServiceBrowserLoginRequestParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getSelfServiceBrowserLoginRequest",
		Method:             "GET",
		PathPattern:        "/self-service/browser/flows/requests/login",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-www-form-urlencoded"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetSelfServiceBrowserLoginRequestReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetSelfServiceBrowserLoginRequestOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getSelfServiceBrowserLoginRequest: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetSelfServiceBrowserProfileManagementRequest gets the request context of browser based profile management flows

  When accessing this endpoint through ORY Kratos' Public API, ensure that cookies are set as they are required
for checking the auth session. To prevent scanning attacks, the public endpoint does not return 404 status codes
but instead 403 or 500.

More information can be found at [ORY Kratos Profile Management Documentation](https://www.ory.sh/docs/next/kratos/self-service/flows/user-profile-management).
*/
func (a *Client) GetSelfServiceBrowserProfileManagementRequest(params *GetSelfServiceBrowserProfileManagementRequestParams) (*GetSelfServiceBrowserProfileManagementRequestOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSelfServiceBrowserProfileManagementRequestParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getSelfServiceBrowserProfileManagementRequest",
		Method:             "GET",
		PathPattern:        "/self-service/browser/flows/requests/profile",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-www-form-urlencoded"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetSelfServiceBrowserProfileManagementRequestReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetSelfServiceBrowserProfileManagementRequestOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getSelfServiceBrowserProfileManagementRequest: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetSelfServiceBrowserRegistrationRequest gets the request context of browser based registration user flows

  This endpoint returns a registration request's context with, for example, error details and
other information.

When accessing this endpoint through ORY Kratos' Public API, ensure that cookies are set as they are required for CSRF to work. To prevent
token scanning attacks, the public endpoint does not return 404 status codes to prevent scanning attacks.

More information can be found at [ORY Kratos User Login and User Registration Documentation](https://www.ory.sh/docs/next/kratos/self-service/flows/user-login-user-registration).
*/
func (a *Client) GetSelfServiceBrowserRegistrationRequest(params *GetSelfServiceBrowserRegistrationRequestParams) (*GetSelfServiceBrowserRegistrationRequestOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSelfServiceBrowserRegistrationRequestParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getSelfServiceBrowserRegistrationRequest",
		Method:             "GET",
		PathPattern:        "/self-service/browser/flows/requests/registration",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-www-form-urlencoded"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetSelfServiceBrowserRegistrationRequestReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetSelfServiceBrowserRegistrationRequestOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getSelfServiceBrowserRegistrationRequest: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetSelfServiceError gets user facing self service errors

  This endpoint returns the error associated with a user-facing self service errors.

When accessing this endpoint through ORY Kratos' Public API, ensure that cookies are set as they are required for CSRF to work. To prevent
token scanning attacks, the public endpoint does not return 404 status codes to prevent scanning attacks.

More information can be found at [ORY Kratos User User Facing Error Documentation](https://www.ory.sh/docs/kratos/self-service/flows/user-facing-errors).
*/
func (a *Client) GetSelfServiceError(params *GetSelfServiceErrorParams) (*GetSelfServiceErrorOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSelfServiceErrorParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getSelfServiceError",
		Method:             "GET",
		PathPattern:        "/self-service/errors",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-www-form-urlencoded"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetSelfServiceErrorReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetSelfServiceErrorOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getSelfServiceError: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
