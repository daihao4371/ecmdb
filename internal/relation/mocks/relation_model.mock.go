// Code generated by MockGen. DO NOT EDIT.
// Source: ./relation_model.go
//
// Generated by this command:
//
//	mockgen -source=./relation_model.go -destination=../../mocks/relation_model.mock.go -package=relationmocks -typed RelationModelService
//

// Package relationmocks is a generated GoMock package.
package relationmocks

import (
	context "context"
	reflect "reflect"

	domain "github.com/Duke1616/ecmdb/internal/relation/internal/domain"
	gomock "go.uber.org/mock/gomock"
)

// MockRelationModelService is a mock of RelationModelService interface.
type MockRelationModelService struct {
	ctrl     *gomock.Controller
	recorder *MockRelationModelServiceMockRecorder
}

// MockRelationModelServiceMockRecorder is the mock recorder for MockRelationModelService.
type MockRelationModelServiceMockRecorder struct {
	mock *MockRelationModelService
}

// NewMockRelationModelService creates a new mock instance.
func NewMockRelationModelService(ctrl *gomock.Controller) *MockRelationModelService {
	mock := &MockRelationModelService{ctrl: ctrl}
	mock.recorder = &MockRelationModelServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRelationModelService) EXPECT() *MockRelationModelServiceMockRecorder {
	return m.recorder
}

// CreateModelRelation mocks base method.
func (m *MockRelationModelService) CreateModelRelation(ctx context.Context, req domain.ModelRelation) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateModelRelation", ctx, req)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateModelRelation indicates an expected call of CreateModelRelation.
func (mr *MockRelationModelServiceMockRecorder) CreateModelRelation(ctx, req any) *MockRelationModelServiceCreateModelRelationCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateModelRelation", reflect.TypeOf((*MockRelationModelService)(nil).CreateModelRelation), ctx, req)
	return &MockRelationModelServiceCreateModelRelationCall{Call: call}
}

// MockRelationModelServiceCreateModelRelationCall wrap *gomock.Call
type MockRelationModelServiceCreateModelRelationCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockRelationModelServiceCreateModelRelationCall) Return(arg0 int64, arg1 error) *MockRelationModelServiceCreateModelRelationCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockRelationModelServiceCreateModelRelationCall) Do(f func(context.Context, domain.ModelRelation) (int64, error)) *MockRelationModelServiceCreateModelRelationCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockRelationModelServiceCreateModelRelationCall) DoAndReturn(f func(context.Context, domain.ModelRelation) (int64, error)) *MockRelationModelServiceCreateModelRelationCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// FindModelDiagramBySrcUids mocks base method.
func (m *MockRelationModelService) FindModelDiagramBySrcUids(ctx context.Context, srcUids []string) ([]domain.ModelDiagram, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindModelDiagramBySrcUids", ctx, srcUids)
	ret0, _ := ret[0].([]domain.ModelDiagram)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindModelDiagramBySrcUids indicates an expected call of FindModelDiagramBySrcUids.
func (mr *MockRelationModelServiceMockRecorder) FindModelDiagramBySrcUids(ctx, srcUids any) *MockRelationModelServiceFindModelDiagramBySrcUidsCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindModelDiagramBySrcUids", reflect.TypeOf((*MockRelationModelService)(nil).FindModelDiagramBySrcUids), ctx, srcUids)
	return &MockRelationModelServiceFindModelDiagramBySrcUidsCall{Call: call}
}

// MockRelationModelServiceFindModelDiagramBySrcUidsCall wrap *gomock.Call
type MockRelationModelServiceFindModelDiagramBySrcUidsCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockRelationModelServiceFindModelDiagramBySrcUidsCall) Return(arg0 []domain.ModelDiagram, arg1 error) *MockRelationModelServiceFindModelDiagramBySrcUidsCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockRelationModelServiceFindModelDiagramBySrcUidsCall) Do(f func(context.Context, []string) ([]domain.ModelDiagram, error)) *MockRelationModelServiceFindModelDiagramBySrcUidsCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockRelationModelServiceFindModelDiagramBySrcUidsCall) DoAndReturn(f func(context.Context, []string) ([]domain.ModelDiagram, error)) *MockRelationModelServiceFindModelDiagramBySrcUidsCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ListModelUidRelation mocks base method.
func (m *MockRelationModelService) ListModelUidRelation(ctx context.Context, offset, limit int64, modelUid string) ([]domain.ModelRelation, int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListModelUidRelation", ctx, offset, limit, modelUid)
	ret0, _ := ret[0].([]domain.ModelRelation)
	ret1, _ := ret[1].(int64)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListModelUidRelation indicates an expected call of ListModelUidRelation.
func (mr *MockRelationModelServiceMockRecorder) ListModelUidRelation(ctx, offset, limit, modelUid any) *MockRelationModelServiceListModelUidRelationCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListModelUidRelation", reflect.TypeOf((*MockRelationModelService)(nil).ListModelUidRelation), ctx, offset, limit, modelUid)
	return &MockRelationModelServiceListModelUidRelationCall{Call: call}
}

// MockRelationModelServiceListModelUidRelationCall wrap *gomock.Call
type MockRelationModelServiceListModelUidRelationCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockRelationModelServiceListModelUidRelationCall) Return(arg0 []domain.ModelRelation, arg1 int64, arg2 error) *MockRelationModelServiceListModelUidRelationCall {
	c.Call = c.Call.Return(arg0, arg1, arg2)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockRelationModelServiceListModelUidRelationCall) Do(f func(context.Context, int64, int64, string) ([]domain.ModelRelation, int64, error)) *MockRelationModelServiceListModelUidRelationCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockRelationModelServiceListModelUidRelationCall) DoAndReturn(f func(context.Context, int64, int64, string) ([]domain.ModelRelation, int64, error)) *MockRelationModelServiceListModelUidRelationCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
