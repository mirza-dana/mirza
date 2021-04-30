// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import context "context"

import entity "accounting/domain/entity"
import mock "github.com/stretchr/testify/mock"
import vo "accounting/domain/vo"

// Outport is an autogenerated mock type for the Outport type
type Outport struct {
	mock.Mock
}

// BeginTransaction provides a mock function with given fields: ctx
func (_m *Outport) BeginTransaction(ctx context.Context) context.Context {
	ret := _m.Called(ctx)

	var r0 context.Context
	if rf, ok := ret.Get(0).(func(context.Context) context.Context); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(context.Context)
		}
	}

	return r0
}

// CommitTransaction provides a mock function with given fields: ctx
func (_m *Outport) CommitTransaction(ctx context.Context) {
	_m.Called(ctx)
}

// FindAllAccountSideByCodes provides a mock function with given fields: ctx, businessID, accountCode
func (_m *Outport) FindAllAccountSideByCodes(ctx context.Context, businessID string, accountCode []string) (map[string]vo.AccountSide, error) {
	ret := _m.Called(ctx, businessID, accountCode)

	var r0 map[string]vo.AccountSide
	if rf, ok := ret.Get(0).(func(context.Context, string, []string) map[string]vo.AccountSide); ok {
		r0 = rf(ctx, businessID, accountCode)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]vo.AccountSide)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, []string) error); ok {
		r1 = rf(ctx, businessID, accountCode)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindAllLastJournalBalance provides a mock function with given fields: ctx, businessID, accountCodes
func (_m *Outport) FindAllLastJournalBalance(ctx context.Context, businessID string, accountCodes []string) (map[string]float64, error) {
	ret := _m.Called(ctx, businessID, accountCodes)

	var r0 map[string]float64
	if rf, ok := ret.Get(0).(func(context.Context, string, []string) map[string]float64); ok {
		r0 = rf(ctx, businessID, accountCodes)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]float64)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, []string) error); ok {
		r1 = rf(ctx, businessID, accountCodes)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GenerateUUID provides a mock function with given fields: ctx
func (_m *Outport) GenerateUUID(ctx context.Context) string {
	ret := _m.Called(ctx)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context) string); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// RollbackTransaction provides a mock function with given fields: ctx
func (_m *Outport) RollbackTransaction(ctx context.Context) {
	_m.Called(ctx)
}

// SaveJournal provides a mock function with given fields: ctx, obj
func (_m *Outport) SaveJournal(ctx context.Context, obj *entity.Journal) error {
	ret := _m.Called(ctx, obj)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *entity.Journal) error); ok {
		r0 = rf(ctx, obj)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}