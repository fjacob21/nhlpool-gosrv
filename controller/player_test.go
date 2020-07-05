package controller

import (
	"testing"

	"nhlpool.com/service/go/nhlpool/store"

	"github.com/stretchr/testify/assert"
	"nhlpool.com/service/go/nhlpool/data"
)

func TestGetPlayer(t *testing.T) {
	store.Clean()
	request := data.AddPlayerRequest{Name: "name", Email: "email", Admin: true, Password: "password"}
	reply := AddPlayer(request)
	playerReply := GetPlayer("name")
	assert.NotNil(t, playerReply, "Should not be nil")
	assert.Equal(t, playerReply.Player.ID, reply.Player.ID, "Invalid ID")
	assert.Equal(t, playerReply.Player.Name, "name", "Invalid name")
	assert.Equal(t, playerReply.Player.Email, "email", "Invalid email")
	assert.Equal(t, playerReply.Player.Admin, true, "Invalid admin")
	assert.Equal(t, playerReply.Player.Password, "", "Invalid password")
}

func TestLogin(t *testing.T) {
	store.Clean()
	request := data.AddPlayerRequest{Name: "name", Email: "email", Admin: false, Password: "password"}
	reply := AddPlayer(request)
	loginReq := data.LoginRequest{Password: "password"}
	loginReply := Login(reply.Player.ID, loginReq)
	assert.Equal(t, loginReply.Result.Code, data.SUCCESS, "Should be a success")
	assert.NotEqual(t, loginReply.SessionID, "", "Should not be empty")
}

func TestLoginBadPassword(t *testing.T) {
	store.Clean()
	request := data.AddPlayerRequest{Name: "name", Email: "email", Admin: false, Password: "password"}
	reply := AddPlayer(request)
	loginReq := data.LoginRequest{Password: "badpassword"}
	loginReply := Login(reply.Player.ID, loginReq)
	assert.Equal(t, loginReply.Result.Code, data.ACCESSDENIED, "Should not be a success")
	assert.Equal(t, loginReply.SessionID, "", "Should be empty")
}

func TestLoginEmptyPassword(t *testing.T) {
	store.Clean()
	request := data.AddPlayerRequest{Name: "name", Email: "email", Admin: false, Password: "password"}
	reply := AddPlayer(request)
	loginReq := data.LoginRequest{Password: ""}
	loginReply := Login(reply.Player.ID, loginReq)
	assert.Equal(t, loginReply.Result.Code, data.ACCESSDENIED, "Should not be a success")
	assert.Equal(t, loginReply.SessionID, "", "Should be empty")
}

func TestLoginBadPlayer(t *testing.T) {
	store.Clean()
	request := data.AddPlayerRequest{Name: "name", Email: "email", Admin: false, Password: "password"}
	AddPlayer(request)
	loginReq := data.LoginRequest{Password: "password"}
	loginReply := Login("invaliduser", loginReq)
	assert.Equal(t, loginReply.Result.Code, data.NOTFOUND, "Should not be a success")
	assert.Equal(t, loginReply.SessionID, "", "Should be empty")
}

func TestLoginEmptyPlayer(t *testing.T) {
	store.Clean()
	request := data.AddPlayerRequest{Name: "name", Email: "email", Admin: false, Password: "password"}
	AddPlayer(request)
	loginReq := data.LoginRequest{Password: "password"}
	loginReply := Login("", loginReq)
	assert.Equal(t, loginReply.Result.Code, data.NOTFOUND, "Should not be a success")
	assert.Equal(t, loginReply.SessionID, "", "Should be empty")
}

func TestLogout(t *testing.T) {
	store.Clean()
	request := data.AddPlayerRequest{Name: "name", Email: "email", Admin: false, Password: "password"}
	reply := AddPlayer(request)
	loginReq := data.LoginRequest{Password: "password"}
	loginReply := Login(reply.Player.ID, loginReq)
	assert.Equal(t, loginReply.Result.Code, data.SUCCESS, "Should be a success")
	assert.NotEqual(t, loginReply.SessionID, "", "Should not be empty")
	logoutReq := data.LogoutRequest{SessionID: loginReply.SessionID}
	logoutReply := Logout(reply.Player.ID, logoutReq)
	assert.Equal(t, logoutReply.Result.Code, data.SUCCESS, "Should be a success")
}

func TestLogoutInvalidSession(t *testing.T) {
	store.Clean()
	request := data.AddPlayerRequest{Name: "name", Email: "email", Admin: false, Password: "password"}
	reply := AddPlayer(request)
	loginReq := data.LoginRequest{Password: "password"}
	loginReply := Login(reply.Player.ID, loginReq)
	assert.Equal(t, loginReply.Result.Code, data.SUCCESS, "Should be a success")
	assert.NotEqual(t, loginReply.SessionID, "", "Should not be empty")
	logoutReq := data.LogoutRequest{SessionID: "invalidsession"}
	logoutReply := Logout(reply.Player.ID, logoutReq)
	assert.Equal(t, logoutReply.Result.Code, data.NOTFOUND, "Should be an error")
}

func TestLogoutEmptySession(t *testing.T) {
	store.Clean()
	request := data.AddPlayerRequest{Name: "name", Email: "email", Admin: false, Password: "password"}
	reply := AddPlayer(request)
	loginReq := data.LoginRequest{Password: "password"}
	loginReply := Login(reply.Player.ID, loginReq)
	assert.Equal(t, loginReply.Result.Code, data.SUCCESS, "Should be a success")
	assert.NotEqual(t, loginReply.SessionID, "", "Should not be empty")
	logoutReq := data.LogoutRequest{SessionID: ""}
	logoutReply := Logout(reply.Player.ID, logoutReq)
	assert.Equal(t, logoutReply.Result.Code, data.NOTFOUND, "Should be an error")
}

func TestLogoutBadPlayer(t *testing.T) {
	store.Clean()
	request := data.AddPlayerRequest{Name: "name", Email: "email", Admin: false, Password: "password"}
	reply := AddPlayer(request)
	loginReq := data.LoginRequest{Password: "password"}
	loginReply := Login(reply.Player.ID, loginReq)
	assert.Equal(t, loginReply.Result.Code, data.SUCCESS, "Should be a success")
	assert.NotEqual(t, loginReply.SessionID, "", "Should not be empty")
	logoutReq := data.LogoutRequest{SessionID: loginReply.SessionID}
	logoutReply := Logout("badplayer", logoutReq)
	assert.Equal(t, logoutReply.Result.Code, data.NOTFOUND, "Should be an error")
}

func TestLogoutEmptyPlayer(t *testing.T) {
	store.Clean()
	request := data.AddPlayerRequest{Name: "name", Email: "email", Admin: false, Password: "password"}
	reply := AddPlayer(request)
	loginReq := data.LoginRequest{Password: "password"}
	loginReply := Login(reply.Player.ID, loginReq)
	assert.Equal(t, loginReply.Result.Code, data.SUCCESS, "Should be a success")
	assert.NotEqual(t, loginReply.SessionID, "", "Should not be empty")
	logoutReq := data.LogoutRequest{SessionID: loginReply.SessionID}
	logoutReply := Logout("", logoutReq)
	assert.Equal(t, logoutReply.Result.Code, data.NOTFOUND, "Should be an error")
}
