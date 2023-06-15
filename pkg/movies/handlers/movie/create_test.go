package handlers_test

import (
	"io"
	handlers "learngo/restapiserver/pkg/movies/handlers/movie"
	model "learngo/restapiserver/pkg/movies/models"
	"learngo/restapiserver/pkg/movies/services"
	"learngo/restapiserver/pkg/movies/services/mocks"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateMovieHandler(t *testing.T) {
	testCases := []struct {
		name          string
		sS            services.Service
		expStatusCode int
		path          string
		body          string
		headers       map[string]string
		preMW         gin.HandlerFunc
	}{
		{
			name: "success",
			sS: func() services.Service {
				sS := new(mocks.Service)
				sS.On("CreateMovie", mock.AnythingOfType("model.MovieTable")).Return(nil)
				return sS
			}(),
			expStatusCode: http.StatusOK,
			path:          "/",
			body: ` {
				"ID": 10,
				"Name": "Test",
				"Year": 1990,
				"Director": "TestDir",
				"Rating": 5
			}`,
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
		{
			name:          "fail invalid payload",
			sS:            func() services.Service { return new(mocks.Service) }(),
			expStatusCode: http.StatusBadRequest,
			path:          "/",
			body: `{
				"ID": 10,
				"Name": "Test",
				"Year": 1990,
				"Director": "TestDir",
			}`,
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
			engine.POST("/", tC.preMW, sH.CreateMovie)
			testServer := httptest.NewServer(engine)
			defer testServer.Close()
			var body io.Reader
			if tC.body != "" {
				body = strings.NewReader(tC.body)
			}
			req, err := http.NewRequest(http.MethodPost, testServer.URL+tC.path, body)
			if !assert.NoError(t, err, "unexpected error") {
				return
			}
			for k, v := range tC.headers {
				req.Header.Set(k, v)
			}
			res, err := http.DefaultClient.Do(req)
			if !assert.NoError(t, err, "unexpected error") {
				return
			}
			assert.Equal(t, tC.expStatusCode, res.StatusCode, "got unexpected statusCode")
			tC.sS.(*mocks.Service).AssertExpectations(t)
		})
	}
}
