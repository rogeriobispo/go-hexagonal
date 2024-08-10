package application_test

import (
	"testing"

	"github.com/rogeriobispo/go-hexagonal/application"
	"github.com/stretchr/testify/require"
)

func TestProduct_Enable(t *testing.T) {
	product := application.Product{}
	product.Name = "Hello"
	product.Status = application.DISABLE
	product.Price = 11
	err := product.Enabled()

	require.Nil(t, err)

	product.Price = 0
	err = product.Enabled()

	require.Equal(t, "o preço deve ser maior que zero", err.Error())
}

func TestProduct_Disable(t *testing.T) {
	product := application.Product{}
	product.Name = "Hello"
	product.Status = application.DISABLE
	product.Price = 0
	err := product.Disable()

	require.Nil(t, err)

	product.Price = 10
	err = product.Disable()

	require.Equal(t, "o produto não pode ser desabilitado", err.Error())

}
