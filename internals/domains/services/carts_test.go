package services

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mrdhira/210622-kuncie-tchnl-test/internals/domains/repository/mocks"
	"github.com/mrdhira/210622-kuncie-tchnl-test/tests"
	"github.com/stretchr/testify/assert"
)

func TestGetCarts(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockRepository := mocks.NewMockIRepository(mockCtrl)
	cartServices := InitCartServices(mockRepository)

	t.Run("Success", func(t *testing.T) {
		mockRepository.EXPECT().FindOrCreateCarts(gomock.Any(), gomock.Any()).Return(tests.CARTS, nil)
		mockRepository.EXPECT().FindAllCartProducts(gomock.Any(), gomock.Any()).Return(tests.CART_PRODUCTS, nil)

		carts, err := cartServices.GetCarts(context.TODO(), 1)
		assert.Equal(t, carts.UserID, uint(1))
		assert.Equal(t, err, nil)
	})

	t.Run("ErrorFindOrCreateCarts", func(t *testing.T) {
		mockRepository.EXPECT().FindOrCreateCarts(gomock.Any(), gomock.Any()).Return(tests.CARTS, errors.New("error"))

		_, err := cartServices.GetCarts(context.TODO(), 1)
		assert.Equal(t, err, errors.New("error"))
	})

	t.Run("ErrorFindAllCartProducts", func(t *testing.T) {
		mockRepository.EXPECT().FindOrCreateCarts(gomock.Any(), gomock.Any()).Return(tests.CARTS, nil)
		mockRepository.EXPECT().FindAllCartProducts(gomock.Any(), gomock.Any()).Return(tests.CART_PRODUCTS, errors.New("error"))

		_, err := cartServices.GetCarts(context.TODO(), 1)
		assert.Equal(t, err, errors.New("error"))
	})
}

func TestCreateCartProducts(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockRepository := mocks.NewMockIRepository(mockCtrl)
	cartServices := InitCartServices(mockRepository)

	t.Run("Success", func(t *testing.T) {
		mockRepository.EXPECT().FindOrCreateCarts(gomock.Any(), gomock.Any()).Return(tests.CARTS, nil)
		mockRepository.EXPECT().InsertOrUpdateCartProducts(gomock.Any(), gomock.Any()).Return(tests.CART_PRODUCT, nil)
		mockRepository.EXPECT().FindAllCartProducts(gomock.Any(), gomock.Any()).Return(tests.CART_PRODUCTS, nil)

		carts, err := cartServices.CreateCartProducts(context.TODO(), 1, tests.CREATE_CART_PRODUCTS)
		assert.Equal(t, carts.UserID, uint(1))
		assert.Equal(t, err, nil)
	})

	t.Run("ErrorFindOrCreateCarts", func(t *testing.T) {
		mockRepository.EXPECT().FindOrCreateCarts(gomock.Any(), gomock.Any()).Return(tests.CARTS, errors.New("error"))

		_, err := cartServices.CreateCartProducts(context.TODO(), 1, tests.CREATE_CART_PRODUCTS)
		assert.Equal(t, err, errors.New("error"))
	})

	t.Run("ErrorInsertOrUpdateCartProducts", func(t *testing.T) {
		mockRepository.EXPECT().FindOrCreateCarts(gomock.Any(), gomock.Any()).Return(tests.CARTS, nil)
		mockRepository.EXPECT().InsertOrUpdateCartProducts(gomock.Any(), gomock.Any()).Return(tests.CART_PRODUCT, errors.New("error"))

		_, err := cartServices.CreateCartProducts(context.TODO(), 1, tests.CREATE_CART_PRODUCTS)
		assert.Equal(t, err, errors.New("error"))
	})

	t.Run("ErrorFindAllCartProducts", func(t *testing.T) {
		mockRepository.EXPECT().FindOrCreateCarts(gomock.Any(), gomock.Any()).Return(tests.CARTS, nil)
		mockRepository.EXPECT().InsertOrUpdateCartProducts(gomock.Any(), gomock.Any()).Return(tests.CART_PRODUCT, nil)
		mockRepository.EXPECT().FindAllCartProducts(gomock.Any(), gomock.Any()).Return(tests.CART_PRODUCTS, errors.New("error"))

		_, err := cartServices.CreateCartProducts(context.TODO(), 1, tests.CREATE_CART_PRODUCTS)
		assert.Equal(t, err, errors.New("error"))
	})
}

func TestUpdateCartProduct(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockRepository := mocks.NewMockIRepository(mockCtrl)
	cartServices := InitCartServices(mockRepository)

	t.Run("Success", func(t *testing.T) {
		mockRepository.EXPECT().FindOrCreateCarts(gomock.Any(), gomock.Any()).Return(tests.CARTS, nil)
		mockRepository.EXPECT().InsertOrUpdateCartProducts(gomock.Any(), gomock.Any()).Return(tests.CART_PRODUCT, nil)
		mockRepository.EXPECT().FindAllCartProducts(gomock.Any(), gomock.Any()).Return(tests.CART_PRODUCTS, nil)

		carts, err := cartServices.UpdateCartProducts(context.TODO(), 1, tests.UPDATE_CART_PRODUCTS)
		assert.Equal(t, carts.UserID, uint(1))
		assert.Equal(t, err, nil)
	})

	t.Run("ErrorFindOrCreateCarts", func(t *testing.T) {
		mockRepository.EXPECT().FindOrCreateCarts(gomock.Any(), gomock.Any()).Return(tests.CARTS, errors.New("error"))

		_, err := cartServices.UpdateCartProducts(context.TODO(), 1, tests.UPDATE_CART_PRODUCTS)
		assert.Equal(t, err, errors.New("error"))
	})

	t.Run("ErrorInsertOrUpdateCartProducts", func(t *testing.T) {
		mockRepository.EXPECT().FindOrCreateCarts(gomock.Any(), gomock.Any()).Return(tests.CARTS, nil)
		mockRepository.EXPECT().InsertOrUpdateCartProducts(gomock.Any(), gomock.Any()).Return(tests.CART_PRODUCT, errors.New("error"))

		_, err := cartServices.UpdateCartProducts(context.TODO(), 1, tests.UPDATE_CART_PRODUCTS)
		assert.Equal(t, err, errors.New("error"))
	})

	t.Run("ErrorFindAllCartProducts", func(t *testing.T) {
		mockRepository.EXPECT().FindOrCreateCarts(gomock.Any(), gomock.Any()).Return(tests.CARTS, nil)
		mockRepository.EXPECT().InsertOrUpdateCartProducts(gomock.Any(), gomock.Any()).Return(tests.CART_PRODUCT, nil)
		mockRepository.EXPECT().FindAllCartProducts(gomock.Any(), gomock.Any()).Return(tests.CART_PRODUCTS, errors.New("error"))

		_, err := cartServices.UpdateCartProducts(context.TODO(), 1, tests.UPDATE_CART_PRODUCTS)
		assert.Equal(t, err, errors.New("error"))
	})
}

func TestDeleteCartProduct(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockRepository := mocks.NewMockIRepository(mockCtrl)
	cartServices := InitCartServices(mockRepository)

	t.Run("Success", func(t *testing.T) {
		mockRepository.EXPECT().FindOrCreateCarts(gomock.Any(), gomock.Any()).Return(tests.CARTS, nil)
		mockRepository.EXPECT().DeleteCartProducts(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)

		err := cartServices.DeleteCartProducts(context.TODO(), 1, tests.DELETE_CART_PRODUCT)
		assert.Equal(t, err, nil)
	})

	t.Run("ErrorFindOrCreateCarts", func(t *testing.T) {
		mockRepository.EXPECT().FindOrCreateCarts(gomock.Any(), gomock.Any()).Return(tests.CARTS, errors.New("error"))

		err := cartServices.DeleteCartProducts(context.TODO(), 1, tests.DELETE_CART_PRODUCT)
		assert.Equal(t, err, errors.New("error"))
	})

	t.Run("ErrorDeleteCartProducts", func(t *testing.T) {
		mockRepository.EXPECT().FindOrCreateCarts(gomock.Any(), gomock.Any()).Return(tests.CARTS, nil)
		mockRepository.EXPECT().DeleteCartProducts(gomock.Any(), gomock.Any(), gomock.Any()).Return(errors.New("error"))

		err := cartServices.DeleteCartProducts(context.TODO(), 1, tests.DELETE_CART_PRODUCT)
		assert.Equal(t, err, errors.New("error"))
	})
}
