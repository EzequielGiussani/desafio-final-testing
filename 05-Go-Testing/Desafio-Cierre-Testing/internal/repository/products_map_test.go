package repository_test

import (
	"app/internal"
	"app/internal/repository"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSearchProducts(t *testing.T) {

	t.Run("should return the solely product found by the product query", func(t *testing.T) {

		// arrange
		var db = make(map[int]internal.Product)

		expectedProduct := internal.Product{
			Id: 1,
			ProductAttributes: internal.ProductAttributes{
				Description: "product 1",
				Price:       100,
				SellerId:    1,
			},
		}

		func(db map[int]internal.Product) {

			db[1] = expectedProduct

		}(db)

		query := internal.ProductQuery{
			Id: 1,
		}

		pm := repository.NewProductsMap(db)

		// act

		res, err := pm.SearchProducts(query)

		// assert

		require.NoError(t, err)
		require.Len(t, res, 1)
		require.Equal(t, expectedProduct, res[1])
	})

	t.Run("should return empty due to no product present", func(t *testing.T) {

		// arrange
		var db = make(map[int]internal.Product)

		query := internal.ProductQuery{
			Id: 1,
		}

		pm := repository.NewProductsMap(db)

		// act

		res, err := pm.SearchProducts(query)

		// assert

		require.NoError(t, err)
		require.Len(t, res, 0)
	})

	t.Run("should return empty due to no element with the provided id found", func(t *testing.T) {

		// arrange
		var db = make(map[int]internal.Product)

		expectedProduct := internal.Product{
			Id: 2,
			ProductAttributes: internal.ProductAttributes{
				Description: "product 1",
				Price:       100,
				SellerId:    1,
			},
		}

		func(db map[int]internal.Product) {

			db[2] = expectedProduct

		}(db)

		query := internal.ProductQuery{
			Id: 1,
		}

		pm := repository.NewProductsMap(db)

		// act

		res, err := pm.SearchProducts(query)

		// assert

		require.NoError(t, err)
		require.Len(t, res, 0)
	})

	t.Run("this test should not pass", func(t *testing.T) {

		// arrange
		var db = make(map[int]internal.Product)

		expectedProduct := internal.Product{
			Id: 1,
			ProductAttributes: internal.ProductAttributes{
				Description: "product 1",
				Price:       100,
				SellerId:    1,
			},
		}

		func(db map[int]internal.Product) {

			db[1] = expectedProduct

		}(db)

		query := internal.ProductQuery{
			Id: -1,
		}

		pm := repository.NewProductsMap(db)

		// act

		res, err := pm.SearchProducts(query)

		// assert

		require.NoError(t, err)
		require.Len(t, res, 1)
		require.Equal(t, expectedProduct, res[1])
	})

}
