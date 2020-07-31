package data

// Conference Define the information about a conference in the pool
type Conference struct {
	ID     string `json:"id"`
	League League `json:"league"`
	Name   string `json:"name"`
}
