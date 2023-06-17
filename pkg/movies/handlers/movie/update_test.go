package handlers_test

import (
	"bytes"
	"encoding/json"
	handlers "learngo/restapiserver/pkg/movies/handlers/movie"
	model "learngo/restapiserver/pkg/movies/models"
	"learngo/restapiserver/pkg/movies/services/mocks"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUpdateMovieHandler(t *testing.T) {
	testCases := []struct {
		name           string
		id             string
		payload        interface{}
		serviceError   error
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "success",
			id:             "1",
			payload:        model.CreateMovieRequest{ID: 1, Name: "Name", Year: 2022, Director: "Updated Director", Rating: 4},
			serviceError:   nil,
			expectedStatus: http.StatusOK,
			expectedBody:   `{"message":"Updated successfully"}`,
		},
		// {
		// 	name:           "invalid payload",
		// 	id:             "2",
		// 	payload:        model.CreateMovieRequest{},
		// 	serviceError:   errors.New("Invalid Payload"),
		// 	expectedStatus: http.StatusBadRequest,
		// 	expectedBody:   `{"error":"Invalid Payload"}`,
		// },
		// {
		// 	name:           "invalid ID",
		// 	id:             "invalid",
		// 	payload:        model.CreateMovieRequest{ID: 5, Name: "Updated Movie", Year: 2022, Director: "Updated Director", Rating: 4},
		// 	serviceError:   nil,
		// 	expectedStatus: http.StatusBadRequest,
		// 	expectedBody:   `{"error":"Invalid ID"}`,
		// },
		// {
		// 	name:           "movie not found",
		// 	id:             "1",
		// 	payload:        model.CreateMovieRequest{Name: "Updated Movie", Year: 2022, Director: "Updated Director", Rating: 4},
		// 	serviceError:   errors.New("Movie not found"),
		// 	expectedStatus: http.StatusNotFound,
		// 	expectedBody:   `{"movie":"Movie not found"}`,
		// },
		// {
		// 	name:           "update error",
		// 	id:             "1",
		// 	payload:        model.CreateMovieRequest{Name: "Updated Movie", Year: 2022, Director: "Updated Director", Rating: 4},
		// 	serviceError:   errors.New("failed to update movie"),
		// 	expectedStatus: http.StatusInternalServerError,
		// 	expectedBody:   `{"error":"Failed to update movie"}`,
		// },
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			engine := gin.New()
			handler := &handlers.Handler{
				Service: &mocks.Service{},
			}
			engine.PUT("/:id", handler.UpdateMovie)

			// Mock the service response
			mockService := handler.Service.(*mocks.Service)
			mockService.On("UpdateMovie", mock.Anything, mock.Anything).Return(tc.serviceError)

			// Prepare the request payload
			payloadBytes, _ := json.Marshal(tc.payload)

			// Perform the request
			w := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodPut, "/"+tc.id, bytes.NewReader(payloadBytes))
			req.Header.Set("Content-Type", "application/json")
			engine.ServeHTTP(w, req)

			// Check the response status code
			assert.Equal(t, tc.expectedStatus, w.Code)

			// Check the response body
			assert.Equal(t, tc.expectedBody, w.Body.String())

			// Assert expectations on the mock service
			mockService.AssertExpectations(t)
		})
	}
}
