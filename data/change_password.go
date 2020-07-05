package data

// ChangePasswordRequest Is the info for a change password request
type ChangePasswordRequest struct {
	SessionID   string `json:"session_id"`
	NewPassword string `json:"new_password"`
}

// ChangePasswordReply Is the reply to a change password request
type ChangePasswordReply struct {
	Result Result `json:"result"`
}
