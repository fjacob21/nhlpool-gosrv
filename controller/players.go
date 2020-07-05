package controller

import (
	"log"

	"nhlpool.com/service/go/nhlpool/store"

	"nhlpool.com/service/go/nhlpool/data"
)

// GetPlayers Process the get players request
func GetPlayers() data.GetPlayersReply {
	log.Println("Get Players")
	var reply data.GetPlayersReply
	reply.Result.Code = data.SUCCESS
	reply.Players = store.GetStore().GetPlayers()
	// Filter password
	for i := range reply.Players {
		reply.Players[i].Password = ""
	}
	return reply
}

// AddPlayer Process the add player request
func AddPlayer(request data.AddPlayerRequest) data.AddPlayerReply {
	var reply data.AddPlayerReply
	log.Println("Add Player", request)
	player := data.CreatePlayer(request.Name, request.Email, request.Admin, request.Password)
	err := store.GetStore().AddPlayer(player)
	if err != nil {
		reply.Result.Code = data.EXISTS
		return reply
	}

	reply.Result.Code = data.SUCCESS
	reply.Player = *player
	// Filter password
	reply.Player.Password = ""
	return reply
}

// ImportPlayer Process the import player request
func ImportPlayer(request data.ImportPlayerRequest) data.ImportPlayerReply {
	var reply data.ImportPlayerReply
	err := store.GetStore().AddPlayer(&request.Player)
	if err != nil {
		reply.Result.Code = data.EXISTS
		return reply
	}
	session := store.GetSessionManager().Get(request.SessionID)
	if session == nil || !session.Player.Admin {
		reply.Result.Code = data.ACCESSDENIED
		return reply
	}
	reply.Result.Code = data.SUCCESS
	reply.Player = request.Player
	return reply
}
