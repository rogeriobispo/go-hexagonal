package application_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/rogeriobispo/go-hexagonal/application"
	mock_application "github.com/rogeriobispo/go-hexagonal/application/mocks"
	"github.com/stretchr/testify/require"
)

func TestProductService_Get(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	product := mock_application.NewMockProductInterface(ctrl)
	persistance := mock_application.NewMockProductPersistenseInterface(ctrl)

	persistance.EXPECT().Get(gomock.Any()).Return(product, nil).AnyTimes()

	service := application.ProductService{
		Persistance: persistance,
	}

	result, err := service.Get("10")
	require.Nil(t, err)
	require.Equal(t, product, result)
}

func TestProductService_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	product := mock_application.NewMockProductInterface(ctrl)
	persistance := mock_application.NewMockProductPersistenseInterface(ctrl)

	persistance.EXPECT().Save(gomock.Any()).Return(product, nil).AnyTimes()

	service := application.ProductService{
		Persistance: persistance,
	}

	result, err := service.Create("Produt one", 10.00)
	require.Nil(t, err)
	require.Equal(t, product, result)
}

func TestProductService_EnabledDisable(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	product := mock_application.NewMockProductInterface(ctrl)

	product.EXPECT().Enabled().Return(nil)
	product.EXPECT().Disable().Return(nil)
	persistance := mock_application.NewMockProductPersistenseInterface(ctrl)

	persistance.EXPECT().Save(gomock.Any()).Return(product, nil).AnyTimes()

	service := application.ProductService{
		Persistance: persistance,
	}

	result, err := service.Enable(product)
	require.Nil(t, err)
	require.Equal(t, product, result)

	result, err = service.Disable(product)
	require.Nil(t, err)
	require.Equal(t, product, result)
}
