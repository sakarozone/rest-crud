package services

import (
	model "learngo/restapiserver/pkg/movies/models"
	"learngo/restapiserver/pkg/movies/store"
)

type Service interface {
	CreateMovie(movie model.MovieTable) error
	UpdateMovie(id int, updatedMovie model.MovieTable) error
	ListMovies(page, pagesize int) (error, []model.MovieTable)
	ListOneMovie(id int) (error, model.MovieTable)
	DeleteMovie(id int) error
}

type service struct {
	store store.Store
}

func New() Service {
	return &service{
		store: store.GetStore(),
	}
}

// acc, sErr := h.Service.CreateFilter(c.Request.Context(), req, user.ID)
