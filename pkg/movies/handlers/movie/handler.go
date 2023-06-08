package controllers

import "learngo/restapiserver/pkg/movies/services"

type Handler struct {
	Service services.Service
}

func New() *Handler {
	return &Handler{
		Service: services.New(),
	}
}
