package services

import (
	model "learngo/restapiserver/models"
	"learngo/restapiserver/store"
)

func CreateMovie(movie model.MovieTable) error {
	err := store.GetStore().CreateMovie(movie)

	if err != nil {
		return err
	}
	return nil

}
