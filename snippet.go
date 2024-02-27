package main

type Snippet struct {
	Content   string `json:"content"`
	Extension string `json:"extension,omitempty"`
	Favorite  bool   `json:"favorite,omitempty"`
	CreatedAt int64  `json:"created_at,omitempty"`
	UpdatedAt int64  `json:"updated_at,omitempty"`
}
