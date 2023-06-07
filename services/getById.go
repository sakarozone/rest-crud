package services

import (
	model "learngo/restapiserver/models"
	"learngo/restapiserver/store"
)

func ListOneMovie(id int) (error, model.MovieTable) {
	err, movie := store.GetStore().ListOneMovie(id)
	if err != nil {
		return err, model.MovieTable{}
	}
	return nil, movie
}
