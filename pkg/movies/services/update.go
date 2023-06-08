package services

import (
	model "learngo/restapiserver/pkg/movies/models"
	"learngo/restapiserver/pkg/movies/store"
)

func (s *service) UpdateMovie(id int, updatedMovie model.MovieTable) error {
	err, movie := s.store.ListOneMovie(id)

	if err != nil {
		return err
	}

	err = store.GetStore().UpdateMovie(movie, updatedMovie)
	if err != nil {
		return err
	}
	return nil

}
