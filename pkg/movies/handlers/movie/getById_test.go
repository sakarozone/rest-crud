package handlers_test

import (
	handlers "learngo/restapiserver/pkg/movies/handlers/movie"
	model "learngo/restapiserver/pkg/movies/models"
	"learngo/restapiserver/pkg/movies/services"
	"learngo/restapiserver/pkg/movies/services/mocks"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetByIDMovieHandler(t *testing.T) {
	testCases := []struct {
		name          string
		sS            services.Service
		expStatusCode int
		path          string
		headers       map[string]string
		preMW         gin.HandlerFunc
	}{
		{
			name: "success",
			sS: func() services.Service {
				sS := new(mocks.Service)
				sS.On("ListOneMovie", mock.Anything).Return(nil, model.MovieTable{
					ID:       10,
					Name:     "test",
					Year:     12,
					Director: "test",
					Rating:   5,
				})
				return sS
			}(),
			expStatusCode: http.StatusOK,
			path:          "/movies/1",
			headers: map[string]string{
				"Content-Type": "application/json",
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
			sH := &handlers.Handler{}
			sH.Service = tC.sS
			engine.GET("/movies/:id", sH.GetMovieById)
			testServer := httptest.NewServer(engine)
			// defer testServer.Close()

			req, err := http.NewRequest(http.MethodGet, testServer.URL+tC.path, nil)
			if err != nil {
				t.Fatalf("failed to create request: %v", err)
			}

			for k, v := range tC.headers {
				req.Header.Set(k, v)
			}

			res, err := http.DefaultClient.Do(req)
			if err != nil {
				t.Fatalf("failed to send request: %v", err)
			}

			defer res.Body.Close()

			assert.Equal(t, tC.expStatusCode, res.StatusCode, "got unexpected statusCode")
			tC.sS.(*mocks.Service).AssertExpectations(t)
		})
	}
}
