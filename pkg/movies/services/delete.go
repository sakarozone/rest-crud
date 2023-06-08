package services

func (s *service) DeleteMovie(id int) error {
	err := s.store.DeleteMovie(id)
	if err != nil {
		return err
	}
	return nil
}
