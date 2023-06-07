package store

import (
	model "learngo/restapiserver/models"

	"gorm.io/gorm"
)

type Store interface {
	CreateMovie(movie *model.MovieTable) error
	UpdateMovie(initialMovie, updatedMovie *model.MovieTable) error
	ListMovies() (error, []model.MovieTable)
	ListOneMovie() (error, model.MovieTable)
	DeleteMovie(id int) error
}
type store struct {
	DB gorm.DB
}

var movieStore Store

// func NewStore() Store {
// 	return &store{
// 		DB: config.GetMongoDB().Collection("filters"),
// 	}
// }

func SetStore(s Store) {
	movieStore = s
}

func (s *store) CreateMovie(movie *model.MovieTable) error {
	return s.DB.Create(movie).Error
}
func (s *store) UpdateMovie(initialMovie, updatedMovie *model.MovieTable) error {
	return s.DB.Model(&initialMovie).Updates(updatedMovie).Error
}

func (s *store) DeleteMovie(id int) error {
	return s.DB.Delete(&model.MovieTable{}, id).Error
}

func (s *store) ListMovies() (error, []model.MovieTable) {
	var movies []model.MovieTable
	return s.DB.Find(&movies).Error, movies
}

func (s *store) ListOneMovie() (error, model.MovieTable) {
	var movie model.MovieTable
	return s.DB.First(&movie).Error, movie
}
