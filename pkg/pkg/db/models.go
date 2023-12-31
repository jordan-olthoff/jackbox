// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0

package db

import ()

type Action struct {
	ID   int64
	Name string
}

type Permission struct {
	ID         int64
	Name       string
	ResourceID int64
	ActionID   int64
}

type Resource struct {
	ID   int64
	Name string
}

type Role struct {
	ID   int64
	Name string
}

type RolePermission struct {
	RoleID       int64
	PermissionID int64
}

type User struct {
	ID        int64
	Username  string
	Firstname string
	Lastname  string
	Email     string
	Password  string
}

type UserRole struct {
	UserID  int64
	RolesID int64
}
