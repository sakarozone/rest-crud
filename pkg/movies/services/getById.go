package services

import (
	model "learngo/restapiserver/pkg/movies/models"
)

func (s *service) ListOneMovie(id int) (error, model.MovieTable) {
	err, movie := s.store.ListOneMovie(id)
	if err != nil {
		return err, model.MovieTable{}
	}
	return nil, movie
}
