package services

import (
	model "learngo/restapiserver/pkg/models"
	"learngo/restapiserver/pkg/movies/store"
)

func UpdateMovie(id int, updatedMovie model.MovieTable) error {
	err, movie := store.GetStore().ListOneMovie(id)

	if err != nil {
		return err
	}

	err = store.GetStore().UpdateMovie(movie, updatedMovie)
	if err != nil {
		return err
	}
	return nil

}
