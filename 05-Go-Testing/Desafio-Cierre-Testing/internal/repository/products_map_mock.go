package repository

import (
	"app/internal"

	"github.com/stretchr/testify/mock"
)

func NewProductsMapMock() *Mock {
	return &Mock{}
}

type Mock struct {
	mock.Mock
	FunSearchProducts func(query internal.ProductQuery) (p map[int]internal.Product, err error)
}

func (m *Mock) SearchProducts(query internal.ProductQuery) (p map[int]internal.Product, err error) {

	output := m.Called(query)

	if m.FunSearchProducts != nil {
		return m.FunSearchProducts(query)
	}

	return output.Get(0).(map[int]internal.Product), output.Error(1)
}
