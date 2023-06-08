package services

import (
	model "learngo/restapiserver/pkg/movies/models"
)

func (s *service) ListMovies() (error, []model.MovieTable) {
	err, movie := s.store.ListMovies()
	if err != nil {
		return err, nil
	}
	return nil, movie
}
