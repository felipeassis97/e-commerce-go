package model

// Response representation of the API response
type Response struct {
	Type    Type   `json:"type"`
	Message string `json:"message"`
}

type Type string

const (
	MissingParams Type = "MissingParams"
	InternalError Type = "InternalError"
	EmptyResponse Type = "EmptyResponse"
	InvalidBody   Type = "InvalidBody"
)
