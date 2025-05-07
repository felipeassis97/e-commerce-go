package model

// User represents one user
type User struct {
	ID       string `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Password string `json:"password,omitempty"`
}
