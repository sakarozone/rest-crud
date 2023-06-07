package services

import (
	model "learngo/restapiserver/pkg/movies/models"
	"learngo/restapiserver/pkg/movies/store"
)

func ListOneMovie(id int) (error, model.MovieTable) {
	err, movie := store.GetStore().ListOneMovie(id)
	if err != nil {
		return err, model.MovieTable{}
	}
	return nil, movie
}
