package application

import (
	"testing"

	"github.com/rogeriobispo/go-hexagonal/application"
	"github.com/stretchr/testify/require"
)

func TestProduct_Enable(t *testing.T) {
	product := application.Product{}
	product.Name = "Hello"
	product.Status = application.DISABLE
	product.price = 10

	err := product.Enabled()

	require.Nil(t, err)
}
