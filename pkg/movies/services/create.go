package services

import (
	model "learngo/restapiserver/pkg/movies/models"
	"learngo/restapiserver/pkg/movies/store"
)

func CreateMovie(movie model.MovieTable) error {
	err := store.GetStore().CreateMovie(movie)

	if err != nil {
		return err
	}
	return nil

}
