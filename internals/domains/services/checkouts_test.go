package services

import (
	"context"
	"errors"
	"sync"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mrdhira/210622-kuncie-tchnl-test/internals/domains/repository/mocks"
	"github.com/mrdhira/210622-kuncie-tchnl-test/tests"
	"github.com/stretchr/testify/assert"
)

func TestCheckouts(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockRepository := mocks.NewMockIRepository(mockCtrl)
	checkoutServices := InitCheckoutServices(mockRepository)

	t.Run("SuccessNoPromo", func(t *testing.T) {
		mockRepository.EXPECT().GetPromos(gomock.Any()).Return(tests.PROMOS, nil)
		mockRepository.EXPECT().BeginTrx(gomock.Any()).Return(context.TODO())
		mockRepository.EXPECT().GetProducts(gomock.Any(), gomock.Any()).Return(tests.PRODUCT_RASPBERRY_PI, nil)
		mockRepository.EXPECT().UpdateProductsQty(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
		mockRepository.EXPECT().CreateOrders(gomock.Any(), gomock.Any()).Return(tests.ORDERS_NO_PROMO, nil)
		mockRepository.EXPECT().CommitTrx(gomock.Any()).Return(nil)

		wg := sync.WaitGroup{}
		wg.Add(1)
		mockRepository.EXPECT().
			DeleteCartProducts(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(nil).
			Do(func(ctx, cartID, productIDs interface{}) {
				defer wg.Done()
			})

		orders, err := checkoutServices.Checkouts(context.TODO(), tests.CHECKOUTS_NO_PROMO)
		wg.Wait()
		assert.Equal(t, orders.ID, uint(1))
		assert.NoError(t, err)
	})

	t.Run("SuccessMacbookProGetRaspberryPi", func(t *testing.T) {
		mockRepository.EXPECT().GetPromos(gomock.Any()).Return(tests.PROMOS, nil)
		mockRepository.EXPECT().BeginTrx(gomock.Any()).Return(context.TODO())
		mockRepository.EXPECT().GetProducts(gomock.Any(), tests.PRODUCT_MACBOOK.ID).Return(tests.PRODUCT_MACBOOK, nil)
		mockRepository.EXPECT().GetProducts(gomock.Any(), tests.PRODUCT_RASPBERRY_PI.ID).Return(tests.PRODUCT_RASPBERRY_PI, nil)
		mockRepository.EXPECT().UpdateProductsQty(gomock.Any(), tests.PRODUCT_MACBOOK.ID, gomock.Any()).Return(nil)
		mockRepository.EXPECT().UpdateProductsQty(gomock.Any(), tests.PRODUCT_RASPBERRY_PI.ID, gomock.Any()).Return(nil)
		mockRepository.EXPECT().CreateOrders(gomock.Any(), gomock.Any()).Return(tests.ORDERS_MACBOOK_PRO_GET_RASPBERRY_PI, nil)
		mockRepository.EXPECT().CommitTrx(gomock.Any()).Return(nil)

		wg := sync.WaitGroup{}
		wg.Add(1)
		mockRepository.EXPECT().
			DeleteCartProducts(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(nil).
			Do(func(ctx, cartID, productIDs interface{}) {
				defer wg.Done()
			})

		orders, err := checkoutServices.Checkouts(context.TODO(), tests.CHECKOUTS_MACBOOK_PRO_GET_RASPBERRY_PI)
		wg.Wait()
		assert.Equal(t, orders.ID, uint(1))
		assert.NoError(t, err)
	})

	t.Run("SuccessGoogleHomeBuy2Get3", func(t *testing.T) {
		mockRepository.EXPECT().GetPromos(gomock.Any()).Return(tests.PROMOS, nil)
		mockRepository.EXPECT().BeginTrx(gomock.Any()).Return(context.TODO())
		mockRepository.EXPECT().GetProducts(gomock.Any(), gomock.Any()).Return(tests.PRODUCT_GOOGLE_HOME, nil)
		mockRepository.EXPECT().UpdateProductsQty(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
		mockRepository.EXPECT().CreateOrders(gomock.Any(), gomock.Any()).Return(tests.ORDERS_GOOGLE_HOME_BY_2_GET_3, nil)
		mockRepository.EXPECT().CommitTrx(gomock.Any()).Return(nil)

		wg := sync.WaitGroup{}
		wg.Add(1)
		mockRepository.EXPECT().
			DeleteCartProducts(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(nil).
			Do(func(ctx, cartID, productIDs interface{}) {
				defer wg.Done()
			})

		orders, err := checkoutServices.Checkouts(context.TODO(), tests.CHECKOUTS_GOOGLE_HOME_BY_2_GET_3)
		wg.Wait()
		assert.Equal(t, orders.ID, uint(1))
		assert.NoError(t, err)
	})

	t.Run("SuccessAlexaSpeakerMinBuy3GetDiscount", func(t *testing.T) {
		mockRepository.EXPECT().GetPromos(gomock.Any()).Return(tests.PROMOS, nil)
		mockRepository.EXPECT().BeginTrx(gomock.Any()).Return(context.TODO())
		mockRepository.EXPECT().GetProducts(gomock.Any(), gomock.Any()).Return(tests.PRODUCT_ALEXA_SPEAKER, nil)
		mockRepository.EXPECT().UpdateProductsQty(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
		mockRepository.EXPECT().CreateOrders(gomock.Any(), gomock.Any()).Return(tests.ORDERS_ALEXA_SPEAKER_MIN_BUY_3_GET_DISCOUNT, nil)
		mockRepository.EXPECT().CommitTrx(gomock.Any()).Return(nil)

		wg := sync.WaitGroup{}
		wg.Add(1)
		mockRepository.EXPECT().
			DeleteCartProducts(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(nil).
			Do(func(ctx, cartID, productIDs interface{}) {
				defer wg.Done()
			})

		orders, err := checkoutServices.Checkouts(context.TODO(), tests.CHECKOUTS_ALEXA_SPEAKER_MIN_BUY_3_GET_DISCOUNT)
		wg.Wait()
		assert.Equal(t, orders.ID, uint(1))
		assert.NoError(t, err)
	})

	t.Run("ErrorGetPromos", func(t *testing.T) {
		mockRepository.EXPECT().GetPromos(gomock.Any()).Return(tests.PROMOS, errors.New("error"))

		_, err := checkoutServices.Checkouts(context.TODO(), tests.CHECKOUTS_ALEXA_SPEAKER_MIN_BUY_3_GET_DISCOUNT)
		assert.Equal(t, err, errors.New("error"))
	})

	t.Run("ErrorGetProducts", func(t *testing.T) {
		mockRepository.EXPECT().GetPromos(gomock.Any()).Return(tests.PROMOS, nil)
		mockRepository.EXPECT().BeginTrx(gomock.Any()).Return(context.TODO())
		mockRepository.EXPECT().GetProducts(gomock.Any(), gomock.Any()).Return(tests.PRODUCT_RASPBERRY_PI, errors.New("error"))
		mockRepository.EXPECT().RollbackTrx(gomock.Any()).Return(nil)

		_, err := checkoutServices.Checkouts(context.TODO(), tests.CHECKOUTS_NO_PROMO)
		assert.Equal(t, err, errors.New("error"))
	})

	t.Run("ErrorUpdateProductsQty", func(t *testing.T) {
		mockRepository.EXPECT().GetPromos(gomock.Any()).Return(tests.PROMOS, nil)
		mockRepository.EXPECT().BeginTrx(gomock.Any()).Return(context.TODO())
		mockRepository.EXPECT().GetProducts(gomock.Any(), gomock.Any()).Return(tests.PRODUCT_RASPBERRY_PI, nil)
		mockRepository.EXPECT().UpdateProductsQty(gomock.Any(), gomock.Any(), gomock.Any()).Return(errors.New("error"))
		mockRepository.EXPECT().RollbackTrx(gomock.Any()).Return(nil)

		_, err := checkoutServices.Checkouts(context.TODO(), tests.CHECKOUTS_NO_PROMO)
		assert.Equal(t, err, errors.New("error"))
	})

	t.Run("ErrorCreateOrders", func(t *testing.T) {
		mockRepository.EXPECT().GetPromos(gomock.Any()).Return(tests.PROMOS, nil)
		mockRepository.EXPECT().BeginTrx(gomock.Any()).Return(context.TODO())
		mockRepository.EXPECT().GetProducts(gomock.Any(), gomock.Any()).Return(tests.PRODUCT_RASPBERRY_PI, nil)
		mockRepository.EXPECT().UpdateProductsQty(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
		mockRepository.EXPECT().CreateOrders(gomock.Any(), gomock.Any()).Return(tests.ORDERS_NO_PROMO, errors.New("error"))
		mockRepository.EXPECT().RollbackTrx(gomock.Any()).Return(nil)

		_, err := checkoutServices.Checkouts(context.TODO(), tests.CHECKOUTS_NO_PROMO)
		assert.Equal(t, err, errors.New("error"))
	})

	t.Run("ErrorCommitTrx", func(t *testing.T) {
		mockRepository.EXPECT().GetPromos(gomock.Any()).Return(tests.PROMOS, nil)
		mockRepository.EXPECT().BeginTrx(gomock.Any()).Return(context.TODO())
		mockRepository.EXPECT().GetProducts(gomock.Any(), gomock.Any()).Return(tests.PRODUCT_RASPBERRY_PI, nil)
		mockRepository.EXPECT().UpdateProductsQty(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
		mockRepository.EXPECT().CreateOrders(gomock.Any(), gomock.Any()).Return(tests.ORDERS_NO_PROMO, nil)
		mockRepository.EXPECT().CommitTrx(gomock.Any()).Return(errors.New("error"))
		mockRepository.EXPECT().RollbackTrx(gomock.Any()).Return(nil)

		_, err := checkoutServices.Checkouts(context.TODO(), tests.CHECKOUTS_NO_PROMO)
		assert.Equal(t, err, errors.New("error"))
	})

	t.Run("SuccessButDeleteCartProductsError", func(t *testing.T) {
		mockRepository.EXPECT().GetPromos(gomock.Any()).Return(tests.PROMOS, nil)
		mockRepository.EXPECT().BeginTrx(gomock.Any()).Return(context.TODO())
		mockRepository.EXPECT().GetProducts(gomock.Any(), gomock.Any()).Return(tests.PRODUCT_RASPBERRY_PI, nil)
		mockRepository.EXPECT().UpdateProductsQty(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
		mockRepository.EXPECT().CreateOrders(gomock.Any(), gomock.Any()).Return(tests.ORDERS_NO_PROMO, nil)
		mockRepository.EXPECT().CommitTrx(gomock.Any()).Return(nil)

		wg := sync.WaitGroup{}
		wg.Add(1)
		mockRepository.EXPECT().
			DeleteCartProducts(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(nil).
			Do(func(ctx, cartID, productIDs interface{}) {
				defer wg.Done()
			})

		orders, err := checkoutServices.Checkouts(context.TODO(), tests.CHECKOUTS_NO_PROMO)
		wg.Wait()
		assert.Equal(t, orders.ID, uint(1))
		assert.NoError(t, err)
	})
}
