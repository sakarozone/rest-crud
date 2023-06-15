package handlers_test

// import (
// 	handlers "learngo/restapiserver/pkg/movies/handlers/movie"
// 	model "learngo/restapiserver/pkg/movies/models"
// 	"learngo/restapiserver/pkg/movies/services/mocks"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"

// 	"github.com/gin-gonic/gin"
// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/mock"
// )

// func TestGetMovieByIdHandler(t *testing.T) {
// 	testCases := []struct {
// 		name            string
// 		id              string
// 		serviceResponse interface{}
// 		expectedStatus  int
// 		expectedBody    string
// 	}{
// 		{
// 			name:            "success",
// 			id:              "1",
// 			serviceResponse: model.MovieTable{ID: 1, Name: "Test Movie", Year: 2021, Director: "Test Director", Rating: 5},
// 			expectedStatus:  http.StatusOK,
// 			expectedBody:    `{"movie":{"ID":1,"Name":"Test Movie","Year":2021,"Director":"Test Director","Rating":5}}`,
// 		},

// 		{
// 			name:            "invalid ID format",
// 			id:              "invalid",
// 			serviceResponse: nil,
// 			expectedStatus:  http.StatusBadRequest,
// 			expectedBody:    `{"error":"Invalid ID format"}`,
// 		},
// 		{
// 			name:            "error retrieving movie",
// 			id:              "1",
// 			serviceResponse: errors.New("failed to retrieve movie"),
// 			expectedStatus:  http.StatusInternalServerError,
// 			expectedBody:    `{"error":"Failed to retrieve the movie"}`,
// 		},
// 	}

// 	for _, tc := range testCases {
// 		t.Run(tc.name, func(t *testing.T) {
// 			engine := gin.New()
// 			handler := &handlers.Handler{
// 				Service: &mocks.Service{},
// 			}
// 			engine.GET("/:id", handler.GetMovieById)

// 			// Mock the service response
// 			mockService := handler.Service.(*mocks.Service)
// 			mockService.On("ListOneMovie", mock.Anything).Return(tc.serviceResponse, nil)

// 			// Perform the request
// 			w := httptest.NewRecorder()
// 			req, _ := http.NewRequest(http.MethodGet, "/"+tc.id, nil)
// 			engine.ServeHTTP(w, req)

// 			// Check the response status code
// 			assert.Equal(t, tc.expectedStatus, w.Code)

// 			// Check the response body
// 			assert.JSONEq(t, tc.expectedBody, w.Body.String())

// 			// Assert expectations on the mock service
// 			mockService.AssertExpectations(t)
// 		})
// 	}
// }
