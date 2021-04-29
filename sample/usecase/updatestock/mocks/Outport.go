// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import context "context"
import entity "accounting/domain/entity"
import mock "github.com/stretchr/testify/mock"

// Outport is an autogenerated mock type for the Outport type
type Outport struct {
	mock.Mock
}

// FindLastStockPrice provides a mock function with given fields: ctx, stockPriceID
func (_m *Outport) FindLastStockPrice(ctx context.Context, stockPriceID string) ([]*entity.StockPrice, error) {
	ret := _m.Called(ctx, stockPriceID)

	var r0 []*entity.StockPrice
	if rf, ok := ret.Get(0).(func(context.Context, string) []*entity.StockPrice); ok {
		r0 = rf(ctx, stockPriceID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*entity.StockPrice)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, stockPriceID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindOneInventory provides a mock function with given fields: ctx, inventoryID
func (_m *Outport) FindOneInventory(ctx context.Context, inventoryID string) (*entity.Inventory, error) {
	ret := _m.Called(ctx, inventoryID)

	var r0 *entity.Inventory
	if rf, ok := ret.Get(0).(func(context.Context, string) *entity.Inventory); ok {
		r0 = rf(ctx, inventoryID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Inventory)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, inventoryID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SaveInventoryStock provides a mock function with given fields: ctx, obj
func (_m *Outport) SaveInventoryStock(ctx context.Context, obj *entity.InventoryBalance) error {
	ret := _m.Called(ctx, obj)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *entity.InventoryBalance) error); ok {
		r0 = rf(ctx, obj)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
