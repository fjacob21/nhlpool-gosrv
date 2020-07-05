package data

// EditPlayerRequest Is the info for an edit player request
type EditPlayerRequest struct {
	SessionID string `json:"session_id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
}

// EditPlayerReply Is the reply to an edit player request
type EditPlayerReply struct {
	Result Result `json:"result"`
	Player Player `json:"player"`
}
