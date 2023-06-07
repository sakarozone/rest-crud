package services

import "learngo/restapiserver/pkg/movies/store"

// import "learngo/restapiserver/pkg/movies/store"

func DeleteMovie(id int) error {
	err := store.GetStore().DeleteMovie(id)
	if err != nil {
		return err
	}
	return nil
}
