// Code generated by MockGen. DO NOT EDIT.
// Source: ./internals/domains/repository/repository.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	dao "github.com/mrdhira/210622-kuncie-tchnl-test/internals/domains/dao"
	gorm "gorm.io/gorm"
)

// MockIRepository is a mock of IRepository interface.
type MockIRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIRepositoryMockRecorder
}

// MockIRepositoryMockRecorder is the mock recorder for MockIRepository.
type MockIRepositoryMockRecorder struct {
	mock *MockIRepository
}

// NewMockIRepository creates a new mock instance.
func NewMockIRepository(ctrl *gomock.Controller) *MockIRepository {
	mock := &MockIRepository{ctrl: ctrl}
	mock.recorder = &MockIRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIRepository) EXPECT() *MockIRepositoryMockRecorder {
	return m.recorder
}

// Begin mocks base method.
func (m *MockIRepository) Begin() *gorm.DB {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Begin")
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

// Begin indicates an expected call of Begin.
func (mr *MockIRepositoryMockRecorder) Begin() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Begin", reflect.TypeOf((*MockIRepository)(nil).Begin))
}

// Commit mocks base method.
func (m *MockIRepository) Commit() *gorm.DB {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Commit")
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

// Commit indicates an expected call of Commit.
func (mr *MockIRepositoryMockRecorder) Commit() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Commit", reflect.TypeOf((*MockIRepository)(nil).Commit))
}

// CreateOrders mocks base method.
func (m *MockIRepository) CreateOrders(ctx context.Context, orders dao.Orders) (dao.Orders, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrders", ctx, orders)
	ret0, _ := ret[0].(dao.Orders)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateOrders indicates an expected call of CreateOrders.
func (mr *MockIRepositoryMockRecorder) CreateOrders(ctx, orders interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrders", reflect.TypeOf((*MockIRepository)(nil).CreateOrders), ctx, orders)
}

// DeleteCartProducts mocks base method.
func (m *MockIRepository) DeleteCartProducts(ctx context.Context, cartID uint, productIDs []uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCartProducts", ctx, cartID, productIDs)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCartProducts indicates an expected call of DeleteCartProducts.
func (mr *MockIRepositoryMockRecorder) DeleteCartProducts(ctx, cartID, productIDs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCartProducts", reflect.TypeOf((*MockIRepository)(nil).DeleteCartProducts), ctx, cartID, productIDs)
}

// FindAllCartProducts mocks base method.
func (m *MockIRepository) FindAllCartProducts(ctx context.Context, cartID uint) ([]dao.CartProducts, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAllCartProducts", ctx, cartID)
	ret0, _ := ret[0].([]dao.CartProducts)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAllCartProducts indicates an expected call of FindAllCartProducts.
func (mr *MockIRepositoryMockRecorder) FindAllCartProducts(ctx, cartID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAllCartProducts", reflect.TypeOf((*MockIRepository)(nil).FindAllCartProducts), ctx, cartID)
}

// FindOrCreateCarts mocks base method.
func (m *MockIRepository) FindOrCreateCarts(ctx context.Context, carts dao.Carts) (dao.Carts, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOrCreateCarts", ctx, carts)
	ret0, _ := ret[0].(dao.Carts)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOrCreateCarts indicates an expected call of FindOrCreateCarts.
func (mr *MockIRepositoryMockRecorder) FindOrCreateCarts(ctx, carts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOrCreateCarts", reflect.TypeOf((*MockIRepository)(nil).FindOrCreateCarts), ctx, carts)
}

// GetProducts mocks base method.
func (m *MockIRepository) GetProducts(ctx context.Context, productID uint) (dao.Products, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProducts", ctx, productID)
	ret0, _ := ret[0].(dao.Products)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProducts indicates an expected call of GetProducts.
func (mr *MockIRepositoryMockRecorder) GetProducts(ctx, productID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProducts", reflect.TypeOf((*MockIRepository)(nil).GetProducts), ctx, productID)
}

// GetPromos mocks base method.
func (m *MockIRepository) GetPromos(ctx context.Context) ([]dao.Promos, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPromos", ctx)
	ret0, _ := ret[0].([]dao.Promos)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPromos indicates an expected call of GetPromos.
func (mr *MockIRepositoryMockRecorder) GetPromos(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPromos", reflect.TypeOf((*MockIRepository)(nil).GetPromos), ctx)
}

// InsertOrUpdateCartProducts mocks base method.
func (m *MockIRepository) InsertOrUpdateCartProducts(ctx context.Context, cartProducts dao.CartProducts) (dao.CartProducts, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertOrUpdateCartProducts", ctx, cartProducts)
	ret0, _ := ret[0].(dao.CartProducts)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertOrUpdateCartProducts indicates an expected call of InsertOrUpdateCartProducts.
func (mr *MockIRepositoryMockRecorder) InsertOrUpdateCartProducts(ctx, cartProducts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertOrUpdateCartProducts", reflect.TypeOf((*MockIRepository)(nil).InsertOrUpdateCartProducts), ctx, cartProducts)
}

// Rollback mocks base method.
func (m *MockIRepository) Rollback() *gorm.DB {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Rollback")
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

// Rollback indicates an expected call of Rollback.
func (mr *MockIRepositoryMockRecorder) Rollback() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Rollback", reflect.TypeOf((*MockIRepository)(nil).Rollback))
}

// UpdateProductsQty mocks base method.
func (m *MockIRepository) UpdateProductsQty(ctx context.Context, productID uint, qty int32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateProductsQty", ctx, productID, qty)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateProductsQty indicates an expected call of UpdateProductsQty.
func (mr *MockIRepositoryMockRecorder) UpdateProductsQty(ctx, productID, qty interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateProductsQty", reflect.TypeOf((*MockIRepository)(nil).UpdateProductsQty), ctx, productID, qty)
}
