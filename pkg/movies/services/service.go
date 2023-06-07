package services

import (
	model "learngo/restapiserver/pkg/movies/models"
	"learngo/restapiserver/pkg/movies/store"
)

type Service interface {
	// C(ctx context.Context, req models.CreateFilterRequest, spUserID primitive.ObjectID) (models.Filter, *serror.ServiceError)
	CreateMovie(movie *model.MovieTable) error
	UpdateMovie(initialMovie, updatedMovie *model.MovieTable) error
	ListMovies() (error, []model.MovieTable)
	ListOneMovie() (error, model.MovieTable)
	DeleteMovie(id int) error
}

type service struct {
	store store.Store
}

// func New() Service {
// 	return &service{
// 		store: store.GetStore(),
// 	}
// }

// acc, sErr := h.Service.CreateFilter(c.Request.Context(), req, user.ID)
