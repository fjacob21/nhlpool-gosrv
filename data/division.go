package data

// Division Define the information about a division in the pool
type Division struct {
	ID     string `json:"id"`
	League League `json:"league"`
	Name   string `json:"name"`
}
