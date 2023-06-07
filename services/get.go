package services

import (
	model "learngo/restapiserver/models"
	"learngo/restapiserver/store"
)

func ListMovies() (error, []model.MovieTable) {
	err, movie := store.GetStore().ListMovies()
	if err != nil {
		return err, nil
	}
	return nil, movie
}
