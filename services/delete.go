package services

import "learngo/restapiserver/store"

func DeleteMovie(id int) error {
	err := store.GetStore().DeleteMovie(id)
	if err != nil {
		return err
	}
	return nil
}
