package controller

import (
	"fmt"

	"nhlpool.com/service/go/nhlpool/data"
	"nhlpool.com/service/go/nhlpool/store"
)

func getPlayer(playerID string) *data.Player {
	player := store.GetStore().Player().GetPlayer(playerID)
	if player == nil {
		player = store.GetStore().Player().GetPlayer(data.UserHash(playerID))
	}
	return player
}

// GetPlayer Process the get player request
func GetPlayer(playerID string) data.GetPlayerReply {
	var reply data.GetPlayerReply
	player := getPlayer(playerID)
	if player == nil {
		reply.Result.Code = data.NOTFOUND
		reply.Player = data.Player{}
		return reply
	}
	reply.Result.Code = data.SUCCESS
	reply.Player = *player
	reply.Player.Password = ""
	return reply
}

// DeletePlayer Process the delete player request
func DeletePlayer(playerID string, request data.DeletePlayerRequest) data.DeletePlayerReply {
	var reply data.DeletePlayerReply
	session := store.GetSessionManager().Get(request.SessionID)
	if session == nil {
		reply.Result.Code = data.ACCESSDENIED
		return reply
	}
	player := getPlayer(playerID)
	if player == nil {
		reply.Result.Code = data.NOTFOUND
		return reply
	}
	if !(session.Player.ID == player.ID || session.Player.Admin) {
		reply.Result.Code = data.ACCESSDENIED
		return reply
	}
	store.GetStore().Player().DeletePlayer(player)
	reply.Result.Code = data.SUCCESS
	return reply
}

// EditPlayer Process the edit player request
func EditPlayer(playerID string, request data.EditPlayerRequest) data.EditPlayerReply {
	var reply data.EditPlayerReply
	session := store.GetSessionManager().Get(request.SessionID)
	if session == nil {
		reply.Result.Code = data.ACCESSDENIED
		return reply
	}
	player := getPlayer(playerID)
	if player == nil {
		reply.Result.Code = data.NOTFOUND
		reply.Player = data.Player{}
		return reply
	}
	if !(session.Player.ID == player.ID || session.Player.Admin) {
		reply.Result.Code = data.ACCESSDENIED
		return reply
	}
	player.Name = request.Name
	player.Email = request.Email
	err := store.GetStore().Player().UpdatePlayer(player)
	if err != nil {
		reply.Result.Code = data.ERROR
		reply.Player = data.Player{}
		return reply
	}
	reply.Result.Code = data.SUCCESS
	reply.Player = *player
	reply.Player.Password = ""
	return reply
}

// Login Process the login request
func Login(playerID string, request data.LoginRequest) data.LoginReply {
	var reply data.LoginReply
	player := getPlayer(playerID)
	if player == nil {
		reply.Result.Code = data.NOTFOUND
		reply.SessionID = ""
		return reply
	}
	reply.SessionID = store.GetSessionManager().Login(player, request.Password)
	if reply.SessionID == "" {
		fmt.Printf("Bad psw no session\n")
		reply.Result.Code = data.ACCESSDENIED
		return reply
	}
	reply.Result.Code = data.SUCCESS
	return reply
}

// Logout Process the logout request
func Logout(playerID string, request data.LogoutRequest) data.LogoutReply {
	var reply data.LogoutReply
	player := getPlayer(playerID)
	if player == nil {
		reply.Result.Code = data.NOTFOUND
		return reply
	}
	err := store.GetSessionManager().Logout(request.SessionID)
	if err != nil {
		reply.Result.Code = data.NOTFOUND
		return reply
	}
	reply.Result.Code = data.SUCCESS
	return reply
}

// ChangePassword Process the change password request
func ChangePassword(playerID string, request data.ChangePasswordRequest) data.ChangePasswordReply {
	var reply data.ChangePasswordReply
	player := getPlayer(playerID)
	if player == nil {
		reply.Result.Code = data.NOTFOUND
		return reply
	}
	session := store.GetSessionManager().Get(request.SessionID)
	if session == nil {
		reply.Result.Code = data.ACCESSDENIED
		return reply
	}
	if !(session.Player.ID == player.ID || session.Player.Admin) {
		reply.Result.Code = data.ACCESSDENIED
		return reply
	}
	player.ChangePassword(request.NewPassword)
	err := store.GetStore().Player().UpdatePlayer(player)
	if err != nil {
		reply.Result.Code = data.ERROR
		return reply
	}
	reply.Result.Code = data.SUCCESS
	return reply
}
