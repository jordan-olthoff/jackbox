// Code generated by ogen, DO NOT EDIT.

package api

import (
	"context"
)

// Handler handles operations described by OpenAPI v3 specification.
type Handler interface {
	// AddUser implements addUser operation.
	//
	// Add a new user.
	//
	// POST /user
	AddUser(ctx context.Context, req *User) (*User, error)
	// GetToken implements getToken operation.
	//
	// Login as user.
	//
	// POST /auth/login
	GetToken(ctx context.Context, req *UsernamePassword) (*LoginResponse, error)
	// GetUserByUsername implements getUserByUsername operation.
	//
	// Returns a single user.
	//
	// GET /user/{username}
	GetUserByUsername(ctx context.Context, params GetUserByUsernameParams) (GetUserByUsernameRes, error)
	// GetUsers implements getUsers operation.
	//
	// Returns a list of users.
	//
	// GET /users
	GetUsers(ctx context.Context, params GetUsersParams) (GetUsersRes, error)
	// RefreshToken implements refreshToken operation.
	//
	// Retrieve new token for user.
	//
	// POST /auth/token
	RefreshToken(ctx context.Context, req *Username) (*Token, error)
	// NewError creates *ErrorStatusCode from error returned by handler.
	//
	// Used for common default response.
	NewError(ctx context.Context, err error) *ErrorStatusCode
}

// Server implements http server based on OpenAPI v3 specification and
// calls Handler to handle requests.
type Server struct {
	h   Handler
	sec SecurityHandler
	baseServer
}

// NewServer creates new Server.
func NewServer(h Handler, sec SecurityHandler, opts ...ServerOption) (*Server, error) {
	s, err := newServerConfig(opts...).baseServer()
	if err != nil {
		return nil, err
	}
	return &Server{
		h:          h,
		sec:        sec,
		baseServer: s,
	}, nil
}
