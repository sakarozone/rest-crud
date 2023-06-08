package services

import (
	model "learngo/restapiserver/pkg/movies/models"
)

func (s *service) CreateMovie(movie model.MovieTable) error {
	err := s.store.CreateMovie(movie)
	if err != nil {
		return err
	}
	return nil
}
