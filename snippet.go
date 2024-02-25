package main

type Snippet struct {
	Content   string `json:"content"`
	Language  string `json:"language,omitempty"`
	CreatedAt int64  `json:"created_at,omitempty"`
	UpdatedAt int64  `json:"updated_at,omitempty"`
}
