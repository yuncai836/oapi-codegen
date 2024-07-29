// Package xorder provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/yuncai836/oapi-codegen/v2 version v2.0.0-00010101000000-000000000000 DO NOT EDIT.
package xorder

// Client defines model for Client.
type Client struct {
	AName *string  `json:"a_name,omitempty"`
	Id    *float32 `json:"id,omitempty"`
}

// ClientWithExtension defines model for ClientWithExtension.
type ClientWithExtension struct {
	Id    *float32 `json:"id,omitempty"`
	AName *string  `json:"a_name,omitempty"`
}
