package services

import (
	model "learngo/restapiserver/pkg/models"
	"learngo/restapiserver/pkg/movies/store"
)

func ListMovies() (error, []model.MovieTable) {
	err, movie := store.GetStore().ListMovies()
	if err != nil {
		return err, nil
	}
	return nil, movie
}
