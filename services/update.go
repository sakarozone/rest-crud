package services

import (
	model "learngo/restapiserver/models"
	"learngo/restapiserver/store"
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
