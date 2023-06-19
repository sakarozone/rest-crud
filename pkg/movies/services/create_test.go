package services

import (
	"errors"
	model "learngo/restapiserver/pkg/movies/models"
	"learngo/restapiserver/pkg/movies/store"
	"learngo/restapiserver/pkg/movies/store/mocks"
	"testing"

	"github.com/go-playground/assert/v2"
	"github.com/stretchr/testify/mock"
)

func TestInsertItem(t *testing.T) {
	testCases := []struct {
		name    string
		args    model.MovieTable
		store   store.Store
		wantErr error
	}{
		{
			name: "success",
			args: model.MovieTable{
				ID:       18989090,
				Name:     "name",
				Year:     2000,
				Director: "Test",
				Rating:   7,
			},
			store: func() store.Store {
				mockStore := new(mocks.Store)
				mockStore.On("CreateMovie",
					mock.Anything,
				).Return(
					nil,
				)
				return mockStore
			}(),
			wantErr: nil,
		},
		{
			name: "insertion failure",
			args: model.MovieTable{
				ID:       18989090,
				Name:     "name",
				Year:     2000,
				Director: "Test",
				Rating:   7,
			},
			store: func() store.Store {
				mockStore := new(mocks.Store)
				mockStore.On("CreateMovie",
					mock.Anything,
				).Return(
					errors.New("Error adding the item"),
				)
				return mockStore
			}(),
			wantErr: errors.New("Error adding the item"),
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.store.CreateMovie(tt.args)
			assert.Equal(t, err, tt.wantErr)
		})
	}
}
