package controllers

import (
	"log"

	"github.com/ankur-toko/muzz/internal/repositories/match"
	"github.com/ankur-toko/muzz/internal/repositories/sql"
	"github.com/ankur-toko/muzz/internal/repositories/user"
)

var mRepo match.MatcherRepository
var uRepo user.UserRepository

var userController UserController
var matchController MatchController

func Initialize() {
	conn, err := sql.Connection()
	if err != nil {
		log.Fatal("error: unable to connect to DB", err)
	}
	mRepo = &match.MatchPostgresRepo{Conn: conn}
	uRepo = &user.UserPostgresRepo{Conn: conn}

	userController = UserController{uRepo, mRepo, MyDummyHash{}}
	matchController = MatchController{mRepo, uRepo, AgeAndLocationScorer{}}
}

func GetUserController() UserController {
	return userController
}

func GetMatchController() MatchController {
	return matchController
}
