package user

import (
	"kholabazar/config"
	"kholabazar/repo"
)

type Handler struct {
	userRepo repo.UserRepo
	conf *config.Config
}

func NewHandler(userRepo repo.UserRepo,conf *config.Config) *Handler {
	return &Handler{
		userRepo: userRepo,
		conf: conf,
	}
}