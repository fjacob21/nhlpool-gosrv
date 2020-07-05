package controller

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"nhlpool.com/service/go/nhlpool/data"
	"nhlpool.com/service/go/nhlpool/store"
)

func TestChangePassword(t *testing.T) {
	store.Clean()
	request := data.AddPlayerRequest{Name: "name", Email: "email", Admin: false, Password: "password"}
	reply := AddPlayer(request)
	loginReq := data.LoginRequest{Password: "password"}
	loginReply := Login(reply.Player.ID, loginReq)
	changePasswordReq := data.ChangePasswordRequest{SessionID: loginReply.SessionID, NewPassword: "password2"}
	changePasswordReply := ChangePassword(reply.Player.ID, changePasswordReq)
	assert.Equal(t, changePasswordReply.Result.Code, data.SUCCESS, "Should be a success")
	logoutReq := data.LogoutRequest{SessionID: loginReply.SessionID}
	logoutReply := Logout(reply.Player.ID, logoutReq)
	assert.Equal(t, logoutReply.Result.Code, data.SUCCESS, "Should be a success")
	loginReq = data.LoginRequest{Password: "password"}
	loginReply = Login(reply.Player.ID, loginReq)
	assert.Equal(t, loginReply.Result.Code, data.ACCESSDENIED, "Should be acccess denied")
	loginReq = data.LoginRequest{Password: "password2"}
	loginReply = Login(reply.Player.ID, loginReq)
	assert.Equal(t, loginReply.Result.Code, data.SUCCESS, "Should be a success")
}

func TestChangePasswordAdmin(t *testing.T) {
	store.Clean()
	requestAdmin := data.AddPlayerRequest{Name: "admin", Email: "email", Admin: true, Password: "password"}
	replyAdmin := AddPlayer(requestAdmin)
	request := data.AddPlayerRequest{Name: "name", Email: "email", Admin: false, Password: "password"}
	reply := AddPlayer(request)
	loginReq := data.LoginRequest{Password: "password"}
	loginReply := Login(replyAdmin.Player.ID, loginReq)
	changePasswordReq := data.ChangePasswordRequest{SessionID: loginReply.SessionID, NewPassword: "password2"}
	changePasswordReply := ChangePassword(reply.Player.ID, changePasswordReq)
	assert.Equal(t, changePasswordReply.Result.Code, data.SUCCESS, "Should be a success")
	loginReq = data.LoginRequest{Password: "password"}
	loginReply = Login(reply.Player.ID, loginReq)
	assert.Equal(t, loginReply.Result.Code, data.ACCESSDENIED, "Should be acccess denied")
	loginReq = data.LoginRequest{Password: "password2"}
	loginReply = Login(reply.Player.ID, loginReq)
	assert.Equal(t, loginReply.Result.Code, data.SUCCESS, "Should be a success")
}

func TestChangePasswordNoAccess(t *testing.T) {
	store.Clean()
	requestAdmin := data.AddPlayerRequest{Name: "admin", Email: "email", Admin: false, Password: "password"}
	replyAdmin := AddPlayer(requestAdmin)
	request := data.AddPlayerRequest{Name: "name", Email: "email", Admin: false, Password: "password"}
	reply := AddPlayer(request)
	loginReq := data.LoginRequest{Password: "password"}
	loginReply := Login(replyAdmin.Player.ID, loginReq)
	changePasswordReq := data.ChangePasswordRequest{SessionID: loginReply.SessionID, NewPassword: "password2"}
	changePasswordReply := ChangePassword(reply.Player.ID, changePasswordReq)
	assert.Equal(t, changePasswordReply.Result.Code, data.ACCESSDENIED, "Should be access denied")
	loginReq = data.LoginRequest{Password: "password2"}
	loginReply = Login(reply.Player.ID, loginReq)
	assert.Equal(t, loginReply.Result.Code, data.ACCESSDENIED, "Should be acccess denied")
	loginReq = data.LoginRequest{Password: "password"}
	loginReply = Login(reply.Player.ID, loginReq)
	assert.Equal(t, loginReply.Result.Code, data.SUCCESS, "Should be a success")
}

func TestChangePasswordInvalidPlayer(t *testing.T) {
	store.Clean()
	requestAdmin := data.AddPlayerRequest{Name: "admin", Email: "email", Admin: false, Password: "password"}
	replyAdmin := AddPlayer(requestAdmin)
	request := data.AddPlayerRequest{Name: "name", Email: "email", Admin: false, Password: "password"}
	reply := AddPlayer(request)
	loginReq := data.LoginRequest{Password: "password"}
	loginReply := Login(replyAdmin.Player.ID, loginReq)
	changePasswordReq := data.ChangePasswordRequest{SessionID: loginReply.SessionID, NewPassword: "password2"}
	changePasswordReply := ChangePassword("invalidplayer", changePasswordReq)
	assert.Equal(t, changePasswordReply.Result.Code, data.NOTFOUND, "Should be not found")
	loginReq = data.LoginRequest{Password: "password2"}
	loginReply = Login(reply.Player.ID, loginReq)
	assert.Equal(t, loginReply.Result.Code, data.ACCESSDENIED, "Should be acccess denied")
	loginReq = data.LoginRequest{Password: "password"}
	loginReply = Login(reply.Player.ID, loginReq)
	assert.Equal(t, loginReply.Result.Code, data.SUCCESS, "Should be a success")
}

func TestChangePasswordEmptyPlayer(t *testing.T) {
	store.Clean()
	requestAdmin := data.AddPlayerRequest{Name: "admin", Email: "email", Admin: false, Password: "password"}
	replyAdmin := AddPlayer(requestAdmin)
	request := data.AddPlayerRequest{Name: "name", Email: "email", Admin: false, Password: "password"}
	reply := AddPlayer(request)
	loginReq := data.LoginRequest{Password: "password"}
	loginReply := Login(replyAdmin.Player.ID, loginReq)
	changePasswordReq := data.ChangePasswordRequest{SessionID: loginReply.SessionID, NewPassword: "password2"}
	changePasswordReply := ChangePassword("", changePasswordReq)
	assert.Equal(t, changePasswordReply.Result.Code, data.NOTFOUND, "Should be not found")
	loginReq = data.LoginRequest{Password: "password2"}
	loginReply = Login(reply.Player.ID, loginReq)
	assert.Equal(t, loginReply.Result.Code, data.ACCESSDENIED, "Should be acccess denied")
	loginReq = data.LoginRequest{Password: "password"}
	loginReply = Login(reply.Player.ID, loginReq)
	assert.Equal(t, loginReply.Result.Code, data.SUCCESS, "Should be a success")
}

func TestChangePasswordInvalidSession(t *testing.T) {
	store.Clean()
	requestAdmin := data.AddPlayerRequest{Name: "admin", Email: "email", Admin: false, Password: "password"}
	replyAdmin := AddPlayer(requestAdmin)
	request := data.AddPlayerRequest{Name: "name", Email: "email", Admin: false, Password: "password"}
	reply := AddPlayer(request)
	loginReq := data.LoginRequest{Password: "password"}
	loginReply := Login(replyAdmin.Player.ID, loginReq)
	changePasswordReq := data.ChangePasswordRequest{SessionID: "invalidsession", NewPassword: "password2"}
	changePasswordReply := ChangePassword(reply.Player.ID, changePasswordReq)
	assert.Equal(t, changePasswordReply.Result.Code, data.ACCESSDENIED, "Should be not found")
	loginReq = data.LoginRequest{Password: "password2"}
	loginReply = Login(reply.Player.ID, loginReq)
	assert.Equal(t, loginReply.Result.Code, data.ACCESSDENIED, "Should be acccess denied")
	loginReq = data.LoginRequest{Password: "password"}
	loginReply = Login(reply.Player.ID, loginReq)
	assert.Equal(t, loginReply.Result.Code, data.SUCCESS, "Should be a success")
}

func TestChangePasswordEmptySession(t *testing.T) {
	store.Clean()
	requestAdmin := data.AddPlayerRequest{Name: "admin", Email: "email", Admin: false, Password: "password"}
	replyAdmin := AddPlayer(requestAdmin)
	request := data.AddPlayerRequest{Name: "name", Email: "email", Admin: false, Password: "password"}
	reply := AddPlayer(request)
	loginReq := data.LoginRequest{Password: "password"}
	loginReply := Login(replyAdmin.Player.ID, loginReq)
	changePasswordReq := data.ChangePasswordRequest{SessionID: "", NewPassword: "password2"}
	changePasswordReply := ChangePassword(reply.Player.ID, changePasswordReq)
	assert.Equal(t, changePasswordReply.Result.Code, data.ACCESSDENIED, "Should be not found")
	loginReq = data.LoginRequest{Password: "password2"}
	loginReply = Login(reply.Player.ID, loginReq)
	assert.Equal(t, loginReply.Result.Code, data.ACCESSDENIED, "Should be acccess denied")
	loginReq = data.LoginRequest{Password: "password"}
	loginReply = Login(reply.Player.ID, loginReq)
	assert.Equal(t, loginReply.Result.Code, data.SUCCESS, "Should be a success")
}
