package services

import (
	model "learngo/restapiserver/pkg/movies/models"
)

func (s *service) ListMovies(page, pagesize int) (error, []model.MovieTable) {
	offset := (page - 1) * pagesize
	err, movie := s.store.ListMovies(page, pagesize, offset)
	if err != nil {
		return err, nil
	}
	return nil, movie
}
