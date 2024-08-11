package application_test

import (
	"testing"

	uuid "github.com/satori/go.uuid"

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

func TestProcut_IsValid(t *testing.T) {
	product := application.Product{}
	product.Id = uuid.NewV4().String()
	product.Name = "Hello"
	product.Status = application.DISABLE
	product.Price = 10
	_, err := product.IsValid()

	require.Nil(t, err)

	product.Status = "Bug do milenio"

	_, err = product.IsValid()

	require.Equal(t, "produto invalido", err.Error())

	product.Price = 0

	_, err = product.IsValid()

	require.Equal(t, "produto invalido", err.Error())
}
