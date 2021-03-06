// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	model "awesomeProject/Practice28-Mockery-Rest-Mongo/pkg/model"
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// Close provides a mock function with given fields:
func (_m *Repository) Close() {
	_m.Called()
}

// CreateStudent provides a mock function with given fields: ctx, Student
func (_m *Repository) CreateStudent(ctx context.Context, Student *model.StudentDetails) (*model.StudentDetails, error) {
	ret := _m.Called(ctx, Student)

	var r0 *model.StudentDetails
	if rf, ok := ret.Get(0).(func(context.Context, *model.StudentDetails) *model.StudentDetails); ok {
		r0 = rf(ctx, Student)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.StudentDetails)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *model.StudentDetails) error); ok {
		r1 = rf(ctx, Student)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteStudent provides a mock function with given fields: ctx, id
func (_m *Repository) DeleteStudent(ctx context.Context, id string) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetStudent provides a mock function with given fields: ctx, id
func (_m *Repository) GetStudent(ctx context.Context, id string) (*model.StudentDetails, error) {
	ret := _m.Called(ctx, id)

	var r0 *model.StudentDetails
	if rf, ok := ret.Get(0).(func(context.Context, string) *model.StudentDetails); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.StudentDetails)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListStudent provides a mock function with given fields: ctx
func (_m *Repository) ListStudent(ctx context.Context) ([]*model.StudentDetails, error) {
	ret := _m.Called(ctx)

	var r0 []*model.StudentDetails
	if rf, ok := ret.Get(0).(func(context.Context) []*model.StudentDetails); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.StudentDetails)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateStudent provides a mock function with given fields: ctx, Student
func (_m *Repository) UpdateStudent(ctx context.Context, Student *model.StudentDetails) (*model.StudentDetails, error) {
	ret := _m.Called(ctx, Student)

	var r0 *model.StudentDetails
	if rf, ok := ret.Get(0).(func(context.Context, *model.StudentDetails) *model.StudentDetails); ok {
		r0 = rf(ctx, Student)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.StudentDetails)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *model.StudentDetails) error); ok {
		r1 = rf(ctx, Student)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
