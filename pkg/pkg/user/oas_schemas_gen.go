// Code generated by ogen, DO NOT EDIT.

package api

import (
	"fmt"
)

func (s *ErrorStatusCode) Error() string {
	return fmt.Sprintf("code %d: %+v", s.StatusCode, s.Response)
}

type BearerAuth struct {
	Token string
}

// GetToken returns the value of Token.
func (s *BearerAuth) GetToken() string {
	return s.Token
}

// SetToken sets the value of Token.
func (s *BearerAuth) SetToken(val string) {
	s.Token = val
}

// Represents error object.
// Ref: #/components/schemas/Error
type Error struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

// GetCode returns the value of Code.
func (s *Error) GetCode() int64 {
	return s.Code
}

// GetMessage returns the value of Message.
func (s *Error) GetMessage() string {
	return s.Message
}

// SetCode sets the value of Code.
func (s *Error) SetCode(val int64) {
	s.Code = val
}

// SetMessage sets the value of Message.
func (s *Error) SetMessage(val string) {
	s.Message = val
}

// ErrorStatusCode wraps Error with StatusCode.
type ErrorStatusCode struct {
	StatusCode int
	Response   Error
}

// GetStatusCode returns the value of StatusCode.
func (s *ErrorStatusCode) GetStatusCode() int {
	return s.StatusCode
}

// GetResponse returns the value of Response.
func (s *ErrorStatusCode) GetResponse() Error {
	return s.Response
}

// SetStatusCode sets the value of StatusCode.
func (s *ErrorStatusCode) SetStatusCode(val int) {
	s.StatusCode = val
}

// SetResponse sets the value of Response.
func (s *ErrorStatusCode) SetResponse(val Error) {
	s.Response = val
}

// GetUserByUsernameNotFound is response for GetUserByUsername operation.
type GetUserByUsernameNotFound struct{}

func (*GetUserByUsernameNotFound) getUserByUsernameRes() {}

// GetUsersNotFound is response for GetUsers operation.
type GetUsersNotFound struct{}

func (*GetUsersNotFound) getUsersRes() {}

// Ref: #/components/schemas/LoginResponse
type LoginResponse struct {
	Token Token `json:"token"`
	User  User  `json:"user"`
}

// GetToken returns the value of Token.
func (s *LoginResponse) GetToken() Token {
	return s.Token
}

// GetUser returns the value of User.
func (s *LoginResponse) GetUser() User {
	return s.User
}

// SetToken sets the value of Token.
func (s *LoginResponse) SetToken(val Token) {
	s.Token = val
}

// SetUser sets the value of User.
func (s *LoginResponse) SetUser(val User) {
	s.User = val
}

// NewOptInt returns new OptInt with value set to v.
func NewOptInt(v int) OptInt {
	return OptInt{
		Value: v,
		Set:   true,
	}
}

// OptInt is optional int.
type OptInt struct {
	Value int
	Set   bool
}

// IsSet returns true if OptInt was set.
func (o OptInt) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptInt) Reset() {
	var v int
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptInt) SetTo(v int) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptInt) Get() (v int, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptInt) Or(d int) int {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// Ref: #/components/schemas/Token
type Token struct {
	Token string `json:"token"`
}

// GetToken returns the value of Token.
func (s *Token) GetToken() string {
	return s.Token
}

// SetToken sets the value of Token.
func (s *Token) SetToken(val string) {
	s.Token = val
}

// Ref: #/components/schemas/User
type User struct {
	Username  string `json:"username"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

// GetUsername returns the value of Username.
func (s *User) GetUsername() string {
	return s.Username
}

// GetFirstname returns the value of Firstname.
func (s *User) GetFirstname() string {
	return s.Firstname
}

// GetLastname returns the value of Lastname.
func (s *User) GetLastname() string {
	return s.Lastname
}

// GetEmail returns the value of Email.
func (s *User) GetEmail() string {
	return s.Email
}

// GetPassword returns the value of Password.
func (s *User) GetPassword() string {
	return s.Password
}

// SetUsername sets the value of Username.
func (s *User) SetUsername(val string) {
	s.Username = val
}

// SetFirstname sets the value of Firstname.
func (s *User) SetFirstname(val string) {
	s.Firstname = val
}

// SetLastname sets the value of Lastname.
func (s *User) SetLastname(val string) {
	s.Lastname = val
}

// SetEmail sets the value of Email.
func (s *User) SetEmail(val string) {
	s.Email = val
}

// SetPassword sets the value of Password.
func (s *User) SetPassword(val string) {
	s.Password = val
}

func (*User) getUserByUsernameRes() {}

type UserList []User

func (*UserList) getUsersRes() {}

// Ref: #/components/schemas/Username
type Username struct {
	Username string `json:"username"`
}

// GetUsername returns the value of Username.
func (s *Username) GetUsername() string {
	return s.Username
}

// SetUsername sets the value of Username.
func (s *Username) SetUsername(val string) {
	s.Username = val
}

// Ref: #/components/schemas/UsernamePassword
type UsernamePassword struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// GetUsername returns the value of Username.
func (s *UsernamePassword) GetUsername() string {
	return s.Username
}

// GetPassword returns the value of Password.
func (s *UsernamePassword) GetPassword() string {
	return s.Password
}

// SetUsername sets the value of Username.
func (s *UsernamePassword) SetUsername(val string) {
	s.Username = val
}

// SetPassword sets the value of Password.
func (s *UsernamePassword) SetPassword(val string) {
	s.Password = val
}
