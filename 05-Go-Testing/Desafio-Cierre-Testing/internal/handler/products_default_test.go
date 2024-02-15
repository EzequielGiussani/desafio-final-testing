package handler_test

import (
	"app/internal"
	"app/internal/handler"
	"app/internal/repository"
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/require"
)

func TestGet(t *testing.T) {

	t.Run("should return 200 ok and the product that matched the id", func(t *testing.T) {

		// arrange
		mockRepo := repository.NewProductsMapMock()
		mockRepo.On("SearchProducts", internal.ProductQuery{Id: 1}).Return(map[int]internal.Product{
			1: {
				Id: 1,
				ProductAttributes: internal.ProductAttributes{
					Description: "product 1",
					Price:       100,
					SellerId:    1,
				},
			},
		}, nil)

		hd := handler.NewProductsDefault(mockRepo)

		req := httptest.NewRequest("GET", "/products?id=1", nil)
		chiCtx := chi.NewRouteContext()
		chiCtx.URLParams.Add("id", "1")
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, chiCtx))
		res := httptest.NewRecorder()

		expectedBody := `{"data":{"1":{"id":1,"description":"product 1","price":100,"seller_id":1}},"message":"success"}`

		// act
		hd.Get()(res, req)

		//assert
		require.Equal(t, http.StatusOK, res.Code)
		require.JSONEq(t, expectedBody, res.Body.String())

	})

	t.Run("should return 200 ok and no product", func(t *testing.T) {

		// arrange
		mockRepo := repository.NewProductsMapMock()
		mockRepo.On("SearchProducts", internal.ProductQuery{Id: 1}).Return(map[int]internal.Product{}, nil)

		hd := handler.NewProductsDefault(mockRepo)

		req := httptest.NewRequest("GET", "/products?id=1", nil)
		chiCtx := chi.NewRouteContext()
		chiCtx.URLParams.Add("id", "1")
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, chiCtx))
		res := httptest.NewRecorder()

		expectedBody := `{"data":{},"message":"success"}`

		// act
		hd.Get()(res, req)

		//assert
		require.Equal(t, http.StatusOK, res.Code)
		require.JSONEq(t, expectedBody, res.Body.String())

	})

	t.Run("should return 500 internal server error", func(t *testing.T) {

		// arrange
		mockRepo := repository.NewProductsMapMock()
		mockRepo.On("SearchProducts", internal.ProductQuery{Id: 1}).Return(map[int]internal.Product{}, errors.New("internal server error"))

		hd := handler.NewProductsDefault(mockRepo)

		req := httptest.NewRequest("GET", "/products?id=1", nil)
		chiCtx := chi.NewRouteContext()
		chiCtx.URLParams.Add("id", "1")
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, chiCtx))
		res := httptest.NewRecorder()

		expectedBody := `{"message":"internal error", "status":"Internal Server Error"}`

		// act
		hd.Get()(res, req)

		//assert
		require.Equal(t, http.StatusInternalServerError, res.Code)
		require.JSONEq(t, expectedBody, res.Body.String())

	})

	t.Run("should return 400 bad request due to invalid id", func(t *testing.T) {

		// arrange
		mockRepo := repository.NewProductsMapMock()

		hd := handler.NewProductsDefault(mockRepo)

		req := httptest.NewRequest("GET", "/products?id=testInvalidId", nil)
		chiCtx := chi.NewRouteContext()
		chiCtx.URLParams.Add("id", "testInvalidId")
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, chiCtx))
		res := httptest.NewRecorder()

		expectedBody := `{"message":"invalid id", "status":"Bad Request"}`

		// act
		hd.Get()(res, req)

		//assert
		require.Equal(t, http.StatusBadRequest, res.Code)
		require.JSONEq(t, expectedBody, res.Body.String())
	})

}
