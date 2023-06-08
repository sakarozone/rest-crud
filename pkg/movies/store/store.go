package store

import (
	"learngo/restapiserver/pkg/common/db"
	model "learngo/restapiserver/pkg/movies/models"

	"gorm.io/gorm"
)

type Store interface {
	CreateMovie(movie model.MovieTable) error
	UpdateMovie(initialMovie, updatedMovie model.MovieTable) error
	ListMovies(page, pagesize, offset int) (error, []model.MovieTable)
	ListOneMovie(id int) (error, model.MovieTable)
	DeleteMovie(id int) error
}
type store struct {
	DB *gorm.DB
}

var movieStore Store

func NewStore() Store {
	db, err := db.ReturnDB()
	if err != nil {
		panic("Failed to connect to database: " + err.Error())
	}
	return &store{
		DB: db,
	}
}

// func SetStore(s Store) {
// 	movieStore = s
// }

func GetStore() Store {
	if movieStore == nil {
		movieStore = NewStore()
	}
	return movieStore
}

func (s *store) CreateMovie(movie model.MovieTable) error {
	return s.DB.Create(movie).Error
}
func (s *store) UpdateMovie(initialMovie, updatedMovie model.MovieTable) error {
	return s.DB.Model(&initialMovie).Updates(updatedMovie).Error
}

func (s *store) DeleteMovie(id int) error {
	return s.DB.Delete(&model.MovieTable{}, id).Error
}

func (s *store) ListMovies(page, pagesize, offset int) (error, []model.MovieTable) {
	var movies []model.MovieTable
	return s.DB.Offset(offset).Limit(pagesize).Find(&movies).Error, movies
}

func (s *store) ListOneMovie(id int) (error, model.MovieTable) {
	var movie model.MovieTable

	return s.DB.First(&movie, id).Error, movie

}
