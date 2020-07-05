package data

// AddPlayerRequest Is the info for an add player request
type AddPlayerRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Admin    bool   `json:"admin"`
	Password string `json:"password"`
}

// AddPlayerReply Is the reply to an add player request
type AddPlayerReply struct {
	Result Result `json:"result"`
	Player Player `json:"player"`
}
