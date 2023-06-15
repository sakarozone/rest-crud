package handlers_test

import (
	"errors"
	handlers "learngo/restapiserver/pkg/movies/handlers/movie"
	model "learngo/restapiserver/pkg/movies/models"
	"learngo/restapiserver/pkg/movies/services"
	"learngo/restapiserver/pkg/movies/services/mocks"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestDeleteMovieHandler(t *testing.T) {
	testCases := []struct {
		name          string
		sS            services.Service
		expStatusCode int
		path          string
		params        gin.Params
		user          model.User
		preMW         gin.HandlerFunc
	}{
		{
			name: "success",
			sS: func() services.Service {
				sS := new(mocks.Service)
				sS.On("DeleteMovie", 10).Return(nil)
				return sS
			}(),
			expStatusCode: http.StatusOK,
			path:          "/movies/10",
			params: gin.Params{
				{Key: "id", Value: "10"},
			},
			preMW: gin.HandlerFunc(func(c *gin.Context) {
				c.Set("User", model.User{
					Email:    "test@test.com",
					Password: "test",
				})
			}),
		},
		{
			name:          "invalid ID",
			sS:            func() services.Service { return new(mocks.Service) }(),
			expStatusCode: http.StatusBadRequest,
			path:          "/movies/abc",
			params: gin.Params{
				{Key: "id", Value: "abc"},
			},
			preMW: gin.HandlerFunc(func(c *gin.Context) {
				c.Set("User", model.User{
					Email:    "test@test.com",
					Password: "test",
				})
			}),
		},
		{
			name: "delete error",
			sS: func() services.Service {
				sS := new(mocks.Service)
				sS.On("DeleteMovie", 20).Return(errors.New("delete error"))
				return sS
			}(),
			expStatusCode: http.StatusInternalServerError,
			path:          "/movies/20",
			params: gin.Params{
				{Key: "id", Value: "20"},
			},

			preMW: gin.HandlerFunc(func(c *gin.Context) {
				c.Set("User", model.User{
					Email:    "test@test.com",
					Password: "test",
				})
			}),
		},
	}

	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			engine := gin.New()
			sH := &handlers.Handler{Service: tC.sS}
			engine.DELETE("/movies/:id", tC.preMW, sH.DeleteMovie)

			req, err := http.NewRequest(http.MethodDelete, tC.path, nil)
			if err != nil {
				t.Fatalf("failed to create request: %v", err)
			}

			req.URL.Path = tC.path
			req.Header.Set("Content-Type", "application/json")

			recorder := httptest.NewRecorder()
			engine.ServeHTTP(recorder, req)

			if recorder.Code != tC.expStatusCode {
				t.Errorf("expected status code %d but got %d", tC.expStatusCode, recorder.Code)
			}
		})
	}
}
