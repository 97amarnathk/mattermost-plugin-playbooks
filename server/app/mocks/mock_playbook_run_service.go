// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mattermost/mattermost-plugin-playbooks/server/app (interfaces: PlaybookRunService)

// Package mock_app is a generated GoMock package.
package mock_app

import (
	gomock "github.com/golang/mock/gomock"
	app "github.com/mattermost/mattermost-plugin-playbooks/server/app"
	model "github.com/mattermost/mattermost-server/v6/model"
	reflect "reflect"
	time "time"
)

// MockPlaybookRunService is a mock of PlaybookRunService interface
type MockPlaybookRunService struct {
	ctrl     *gomock.Controller
	recorder *MockPlaybookRunServiceMockRecorder
}

// MockPlaybookRunServiceMockRecorder is the mock recorder for MockPlaybookRunService
type MockPlaybookRunServiceMockRecorder struct {
	mock *MockPlaybookRunService
}

// NewMockPlaybookRunService creates a new mock instance
func NewMockPlaybookRunService(ctrl *gomock.Controller) *MockPlaybookRunService {
	mock := &MockPlaybookRunService{ctrl: ctrl}
	mock.recorder = &MockPlaybookRunServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPlaybookRunService) EXPECT() *MockPlaybookRunServiceMockRecorder {
	return m.recorder
}

// AddChecklist mocks base method
func (m *MockPlaybookRunService) AddChecklist(arg0, arg1 string, arg2 app.Checklist) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddChecklist", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddChecklist indicates an expected call of AddChecklist
func (mr *MockPlaybookRunServiceMockRecorder) AddChecklist(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddChecklist", reflect.TypeOf((*MockPlaybookRunService)(nil).AddChecklist), arg0, arg1, arg2)
}

// AddChecklistItem mocks base method
func (m *MockPlaybookRunService) AddChecklistItem(arg0, arg1 string, arg2 int, arg3 app.ChecklistItem) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddChecklistItem", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddChecklistItem indicates an expected call of AddChecklistItem
func (mr *MockPlaybookRunServiceMockRecorder) AddChecklistItem(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddChecklistItem", reflect.TypeOf((*MockPlaybookRunService)(nil).AddChecklistItem), arg0, arg1, arg2, arg3)
}

// AddPostToTimeline mocks base method
func (m *MockPlaybookRunService) AddPostToTimeline(arg0, arg1, arg2, arg3 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddPostToTimeline", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddPostToTimeline indicates an expected call of AddPostToTimeline
func (mr *MockPlaybookRunServiceMockRecorder) AddPostToTimeline(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddPostToTimeline", reflect.TypeOf((*MockPlaybookRunService)(nil).AddPostToTimeline), arg0, arg1, arg2, arg3)
}

// CancelRetrospective mocks base method
func (m *MockPlaybookRunService) CancelRetrospective(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CancelRetrospective", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CancelRetrospective indicates an expected call of CancelRetrospective
func (mr *MockPlaybookRunServiceMockRecorder) CancelRetrospective(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CancelRetrospective", reflect.TypeOf((*MockPlaybookRunService)(nil).CancelRetrospective), arg0, arg1)
}

// ChangeCreationDate mocks base method
func (m *MockPlaybookRunService) ChangeCreationDate(arg0 string, arg1 time.Time) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangeCreationDate", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ChangeCreationDate indicates an expected call of ChangeCreationDate
func (mr *MockPlaybookRunServiceMockRecorder) ChangeCreationDate(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeCreationDate", reflect.TypeOf((*MockPlaybookRunService)(nil).ChangeCreationDate), arg0, arg1)
}

// ChangeOwner mocks base method
func (m *MockPlaybookRunService) ChangeOwner(arg0, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangeOwner", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// ChangeOwner indicates an expected call of ChangeOwner
func (mr *MockPlaybookRunServiceMockRecorder) ChangeOwner(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeOwner", reflect.TypeOf((*MockPlaybookRunService)(nil).ChangeOwner), arg0, arg1, arg2)
}

// CheckAndSendMessageOnJoin mocks base method
func (m *MockPlaybookRunService) CheckAndSendMessageOnJoin(arg0, arg1, arg2 string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckAndSendMessageOnJoin", arg0, arg1, arg2)
	ret0, _ := ret[0].(bool)
	return ret0
}

// CheckAndSendMessageOnJoin indicates an expected call of CheckAndSendMessageOnJoin
func (mr *MockPlaybookRunServiceMockRecorder) CheckAndSendMessageOnJoin(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckAndSendMessageOnJoin", reflect.TypeOf((*MockPlaybookRunService)(nil).CheckAndSendMessageOnJoin), arg0, arg1, arg2)
}

// CreatePlaybookRun mocks base method
func (m *MockPlaybookRunService) CreatePlaybookRun(arg0 *app.PlaybookRun, arg1 *app.Playbook, arg2 string, arg3 bool) (*app.PlaybookRun, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePlaybookRun", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*app.PlaybookRun)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreatePlaybookRun indicates an expected call of CreatePlaybookRun
func (mr *MockPlaybookRunServiceMockRecorder) CreatePlaybookRun(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePlaybookRun", reflect.TypeOf((*MockPlaybookRunService)(nil).CreatePlaybookRun), arg0, arg1, arg2, arg3)
}

// DMTodoDigestToUser mocks base method
func (m *MockPlaybookRunService) DMTodoDigestToUser(arg0 string, arg1 bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DMTodoDigestToUser", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DMTodoDigestToUser indicates an expected call of DMTodoDigestToUser
func (mr *MockPlaybookRunServiceMockRecorder) DMTodoDigestToUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DMTodoDigestToUser", reflect.TypeOf((*MockPlaybookRunService)(nil).DMTodoDigestToUser), arg0, arg1)
}

// EditChecklistItem mocks base method
func (m *MockPlaybookRunService) EditChecklistItem(arg0, arg1 string, arg2, arg3 int, arg4, arg5, arg6 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EditChecklistItem", arg0, arg1, arg2, arg3, arg4, arg5, arg6)
	ret0, _ := ret[0].(error)
	return ret0
}

// EditChecklistItem indicates an expected call of EditChecklistItem
func (mr *MockPlaybookRunServiceMockRecorder) EditChecklistItem(arg0, arg1, arg2, arg3, arg4, arg5, arg6 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EditChecklistItem", reflect.TypeOf((*MockPlaybookRunService)(nil).EditChecklistItem), arg0, arg1, arg2, arg3, arg4, arg5, arg6)
}

// FinishPlaybookRun mocks base method
func (m *MockPlaybookRunService) FinishPlaybookRun(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FinishPlaybookRun", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// FinishPlaybookRun indicates an expected call of FinishPlaybookRun
func (mr *MockPlaybookRunServiceMockRecorder) FinishPlaybookRun(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FinishPlaybookRun", reflect.TypeOf((*MockPlaybookRunService)(nil).FinishPlaybookRun), arg0, arg1)
}

// Follow mocks base method
func (m *MockPlaybookRunService) Follow(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Follow", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Follow indicates an expected call of Follow
func (mr *MockPlaybookRunServiceMockRecorder) Follow(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Follow", reflect.TypeOf((*MockPlaybookRunService)(nil).Follow), arg0, arg1)
}

// GetChecklistAutocomplete mocks base method
func (m *MockPlaybookRunService) GetChecklistAutocomplete(arg0 string) ([]model.AutocompleteListItem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetChecklistAutocomplete", arg0)
	ret0, _ := ret[0].([]model.AutocompleteListItem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetChecklistAutocomplete indicates an expected call of GetChecklistAutocomplete
func (mr *MockPlaybookRunServiceMockRecorder) GetChecklistAutocomplete(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetChecklistAutocomplete", reflect.TypeOf((*MockPlaybookRunService)(nil).GetChecklistAutocomplete), arg0)
}

// GetChecklistItemAutocomplete mocks base method
func (m *MockPlaybookRunService) GetChecklistItemAutocomplete(arg0 string) ([]model.AutocompleteListItem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetChecklistItemAutocomplete", arg0)
	ret0, _ := ret[0].([]model.AutocompleteListItem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetChecklistItemAutocomplete indicates an expected call of GetChecklistItemAutocomplete
func (mr *MockPlaybookRunServiceMockRecorder) GetChecklistItemAutocomplete(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetChecklistItemAutocomplete", reflect.TypeOf((*MockPlaybookRunService)(nil).GetChecklistItemAutocomplete), arg0)
}

// GetFollowers mocks base method
func (m *MockPlaybookRunService) GetFollowers(arg0 string) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFollowers", arg0)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFollowers indicates an expected call of GetFollowers
func (mr *MockPlaybookRunServiceMockRecorder) GetFollowers(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFollowers", reflect.TypeOf((*MockPlaybookRunService)(nil).GetFollowers), arg0)
}

// GetOverdueUpdateRuns mocks base method
func (m *MockPlaybookRunService) GetOverdueUpdateRuns(arg0 string) ([]app.RunLink, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOverdueUpdateRuns", arg0)
	ret0, _ := ret[0].([]app.RunLink)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOverdueUpdateRuns indicates an expected call of GetOverdueUpdateRuns
func (mr *MockPlaybookRunServiceMockRecorder) GetOverdueUpdateRuns(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOverdueUpdateRuns", reflect.TypeOf((*MockPlaybookRunService)(nil).GetOverdueUpdateRuns), arg0)
}

// GetOwners mocks base method
func (m *MockPlaybookRunService) GetOwners(arg0 app.RequesterInfo, arg1 app.PlaybookRunFilterOptions) ([]app.OwnerInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOwners", arg0, arg1)
	ret0, _ := ret[0].([]app.OwnerInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOwners indicates an expected call of GetOwners
func (mr *MockPlaybookRunServiceMockRecorder) GetOwners(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOwners", reflect.TypeOf((*MockPlaybookRunService)(nil).GetOwners), arg0, arg1)
}

// GetParticipatingRuns mocks base method
func (m *MockPlaybookRunService) GetParticipatingRuns(arg0 string) ([]app.RunLink, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetParticipatingRuns", arg0)
	ret0, _ := ret[0].([]app.RunLink)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetParticipatingRuns indicates an expected call of GetParticipatingRuns
func (mr *MockPlaybookRunServiceMockRecorder) GetParticipatingRuns(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetParticipatingRuns", reflect.TypeOf((*MockPlaybookRunService)(nil).GetParticipatingRuns), arg0)
}

// GetPlaybookRun mocks base method
func (m *MockPlaybookRunService) GetPlaybookRun(arg0 string) (*app.PlaybookRun, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPlaybookRun", arg0)
	ret0, _ := ret[0].(*app.PlaybookRun)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPlaybookRun indicates an expected call of GetPlaybookRun
func (mr *MockPlaybookRunServiceMockRecorder) GetPlaybookRun(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPlaybookRun", reflect.TypeOf((*MockPlaybookRunService)(nil).GetPlaybookRun), arg0)
}

// GetPlaybookRunIDForChannel mocks base method
func (m *MockPlaybookRunService) GetPlaybookRunIDForChannel(arg0 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPlaybookRunIDForChannel", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPlaybookRunIDForChannel indicates an expected call of GetPlaybookRunIDForChannel
func (mr *MockPlaybookRunServiceMockRecorder) GetPlaybookRunIDForChannel(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPlaybookRunIDForChannel", reflect.TypeOf((*MockPlaybookRunService)(nil).GetPlaybookRunIDForChannel), arg0)
}

// GetPlaybookRunMetadata mocks base method
func (m *MockPlaybookRunService) GetPlaybookRunMetadata(arg0 string) (*app.Metadata, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPlaybookRunMetadata", arg0)
	ret0, _ := ret[0].(*app.Metadata)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPlaybookRunMetadata indicates an expected call of GetPlaybookRunMetadata
func (mr *MockPlaybookRunServiceMockRecorder) GetPlaybookRunMetadata(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPlaybookRunMetadata", reflect.TypeOf((*MockPlaybookRunService)(nil).GetPlaybookRunMetadata), arg0)
}

// GetPlaybookRuns mocks base method
func (m *MockPlaybookRunService) GetPlaybookRuns(arg0 app.RequesterInfo, arg1 app.PlaybookRunFilterOptions) (*app.GetPlaybookRunsResults, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPlaybookRuns", arg0, arg1)
	ret0, _ := ret[0].(*app.GetPlaybookRunsResults)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPlaybookRuns indicates an expected call of GetPlaybookRuns
func (mr *MockPlaybookRunServiceMockRecorder) GetPlaybookRuns(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPlaybookRuns", reflect.TypeOf((*MockPlaybookRunService)(nil).GetPlaybookRuns), arg0, arg1)
}

// GetRunsWithAssignedTasks mocks base method
func (m *MockPlaybookRunService) GetRunsWithAssignedTasks(arg0 string) ([]app.AssignedRun, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRunsWithAssignedTasks", arg0)
	ret0, _ := ret[0].([]app.AssignedRun)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRunsWithAssignedTasks indicates an expected call of GetRunsWithAssignedTasks
func (mr *MockPlaybookRunServiceMockRecorder) GetRunsWithAssignedTasks(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRunsWithAssignedTasks", reflect.TypeOf((*MockPlaybookRunService)(nil).GetRunsWithAssignedTasks), arg0)
}

// HandleReminder mocks base method
func (m *MockPlaybookRunService) HandleReminder(arg0 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "HandleReminder", arg0)
}

// HandleReminder indicates an expected call of HandleReminder
func (mr *MockPlaybookRunServiceMockRecorder) HandleReminder(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleReminder", reflect.TypeOf((*MockPlaybookRunService)(nil).HandleReminder), arg0)
}

// IsOwner mocks base method
func (m *MockPlaybookRunService) IsOwner(arg0, arg1 string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsOwner", arg0, arg1)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsOwner indicates an expected call of IsOwner
func (mr *MockPlaybookRunServiceMockRecorder) IsOwner(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsOwner", reflect.TypeOf((*MockPlaybookRunService)(nil).IsOwner), arg0, arg1)
}

// ModifyCheckedState mocks base method
func (m *MockPlaybookRunService) ModifyCheckedState(arg0, arg1, arg2 string, arg3, arg4 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModifyCheckedState", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(error)
	return ret0
}

// ModifyCheckedState indicates an expected call of ModifyCheckedState
func (mr *MockPlaybookRunServiceMockRecorder) ModifyCheckedState(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModifyCheckedState", reflect.TypeOf((*MockPlaybookRunService)(nil).ModifyCheckedState), arg0, arg1, arg2, arg3, arg4)
}

// MoveChecklist mocks base method
func (m *MockPlaybookRunService) MoveChecklist(arg0, arg1 string, arg2, arg3 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MoveChecklist", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// MoveChecklist indicates an expected call of MoveChecklist
func (mr *MockPlaybookRunServiceMockRecorder) MoveChecklist(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MoveChecklist", reflect.TypeOf((*MockPlaybookRunService)(nil).MoveChecklist), arg0, arg1, arg2, arg3)
}

// MoveChecklistItem mocks base method
func (m *MockPlaybookRunService) MoveChecklistItem(arg0, arg1 string, arg2, arg3, arg4, arg5 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MoveChecklistItem", arg0, arg1, arg2, arg3, arg4, arg5)
	ret0, _ := ret[0].(error)
	return ret0
}

// MoveChecklistItem indicates an expected call of MoveChecklistItem
func (mr *MockPlaybookRunServiceMockRecorder) MoveChecklistItem(arg0, arg1, arg2, arg3, arg4, arg5 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MoveChecklistItem", reflect.TypeOf((*MockPlaybookRunService)(nil).MoveChecklistItem), arg0, arg1, arg2, arg3, arg4, arg5)
}

// NukeDB mocks base method
func (m *MockPlaybookRunService) NukeDB() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NukeDB")
	ret0, _ := ret[0].(error)
	return ret0
}

// NukeDB indicates an expected call of NukeDB
func (mr *MockPlaybookRunServiceMockRecorder) NukeDB() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NukeDB", reflect.TypeOf((*MockPlaybookRunService)(nil).NukeDB))
}

// OpenAddChecklistItemDialog mocks base method
func (m *MockPlaybookRunService) OpenAddChecklistItemDialog(arg0, arg1 string, arg2 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OpenAddChecklistItemDialog", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// OpenAddChecklistItemDialog indicates an expected call of OpenAddChecklistItemDialog
func (mr *MockPlaybookRunServiceMockRecorder) OpenAddChecklistItemDialog(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OpenAddChecklistItemDialog", reflect.TypeOf((*MockPlaybookRunService)(nil).OpenAddChecklistItemDialog), arg0, arg1, arg2)
}

// OpenAddToTimelineDialog mocks base method
func (m *MockPlaybookRunService) OpenAddToTimelineDialog(arg0 app.RequesterInfo, arg1, arg2, arg3 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OpenAddToTimelineDialog", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// OpenAddToTimelineDialog indicates an expected call of OpenAddToTimelineDialog
func (mr *MockPlaybookRunServiceMockRecorder) OpenAddToTimelineDialog(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OpenAddToTimelineDialog", reflect.TypeOf((*MockPlaybookRunService)(nil).OpenAddToTimelineDialog), arg0, arg1, arg2, arg3)
}

// OpenCreatePlaybookRunDialog mocks base method
func (m *MockPlaybookRunService) OpenCreatePlaybookRunDialog(arg0, arg1, arg2, arg3, arg4 string, arg5 []app.Playbook, arg6 bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OpenCreatePlaybookRunDialog", arg0, arg1, arg2, arg3, arg4, arg5, arg6)
	ret0, _ := ret[0].(error)
	return ret0
}

// OpenCreatePlaybookRunDialog indicates an expected call of OpenCreatePlaybookRunDialog
func (mr *MockPlaybookRunServiceMockRecorder) OpenCreatePlaybookRunDialog(arg0, arg1, arg2, arg3, arg4, arg5, arg6 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OpenCreatePlaybookRunDialog", reflect.TypeOf((*MockPlaybookRunService)(nil).OpenCreatePlaybookRunDialog), arg0, arg1, arg2, arg3, arg4, arg5, arg6)
}

// OpenFinishPlaybookRunDialog mocks base method
func (m *MockPlaybookRunService) OpenFinishPlaybookRunDialog(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OpenFinishPlaybookRunDialog", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// OpenFinishPlaybookRunDialog indicates an expected call of OpenFinishPlaybookRunDialog
func (mr *MockPlaybookRunServiceMockRecorder) OpenFinishPlaybookRunDialog(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OpenFinishPlaybookRunDialog", reflect.TypeOf((*MockPlaybookRunService)(nil).OpenFinishPlaybookRunDialog), arg0, arg1)
}

// OpenUpdateStatusDialog mocks base method
func (m *MockPlaybookRunService) OpenUpdateStatusDialog(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OpenUpdateStatusDialog", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// OpenUpdateStatusDialog indicates an expected call of OpenUpdateStatusDialog
func (mr *MockPlaybookRunServiceMockRecorder) OpenUpdateStatusDialog(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OpenUpdateStatusDialog", reflect.TypeOf((*MockPlaybookRunService)(nil).OpenUpdateStatusDialog), arg0, arg1)
}

// PublishRetrospective mocks base method
func (m *MockPlaybookRunService) PublishRetrospective(arg0, arg1 string, arg2 app.RetrospectiveUpdate) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PublishRetrospective", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// PublishRetrospective indicates an expected call of PublishRetrospective
func (mr *MockPlaybookRunServiceMockRecorder) PublishRetrospective(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PublishRetrospective", reflect.TypeOf((*MockPlaybookRunService)(nil).PublishRetrospective), arg0, arg1, arg2)
}

// RemoveChecklist mocks base method
func (m *MockPlaybookRunService) RemoveChecklist(arg0, arg1 string, arg2 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveChecklist", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveChecklist indicates an expected call of RemoveChecklist
func (mr *MockPlaybookRunServiceMockRecorder) RemoveChecklist(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveChecklist", reflect.TypeOf((*MockPlaybookRunService)(nil).RemoveChecklist), arg0, arg1, arg2)
}

// RemoveChecklistItem mocks base method
func (m *MockPlaybookRunService) RemoveChecklistItem(arg0, arg1 string, arg2, arg3 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveChecklistItem", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveChecklistItem indicates an expected call of RemoveChecklistItem
func (mr *MockPlaybookRunServiceMockRecorder) RemoveChecklistItem(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveChecklistItem", reflect.TypeOf((*MockPlaybookRunService)(nil).RemoveChecklistItem), arg0, arg1, arg2, arg3)
}

// RemoveReminder mocks base method
func (m *MockPlaybookRunService) RemoveReminder(arg0 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RemoveReminder", arg0)
}

// RemoveReminder indicates an expected call of RemoveReminder
func (mr *MockPlaybookRunServiceMockRecorder) RemoveReminder(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveReminder", reflect.TypeOf((*MockPlaybookRunService)(nil).RemoveReminder), arg0)
}

// RemoveTimelineEvent mocks base method
func (m *MockPlaybookRunService) RemoveTimelineEvent(arg0, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveTimelineEvent", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveTimelineEvent indicates an expected call of RemoveTimelineEvent
func (mr *MockPlaybookRunServiceMockRecorder) RemoveTimelineEvent(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveTimelineEvent", reflect.TypeOf((*MockPlaybookRunService)(nil).RemoveTimelineEvent), arg0, arg1, arg2)
}

// RenameChecklist mocks base method
func (m *MockPlaybookRunService) RenameChecklist(arg0, arg1 string, arg2 int, arg3 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RenameChecklist", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// RenameChecklist indicates an expected call of RenameChecklist
func (mr *MockPlaybookRunServiceMockRecorder) RenameChecklist(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RenameChecklist", reflect.TypeOf((*MockPlaybookRunService)(nil).RenameChecklist), arg0, arg1, arg2, arg3)
}

// RestoreChecklistItem mocks base method
func (m *MockPlaybookRunService) RestoreChecklistItem(arg0, arg1 string, arg2, arg3 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RestoreChecklistItem", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// RestoreChecklistItem indicates an expected call of RestoreChecklistItem
func (mr *MockPlaybookRunServiceMockRecorder) RestoreChecklistItem(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RestoreChecklistItem", reflect.TypeOf((*MockPlaybookRunService)(nil).RestoreChecklistItem), arg0, arg1, arg2, arg3)
}

// RestorePlaybookRun mocks base method
func (m *MockPlaybookRunService) RestorePlaybookRun(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RestorePlaybookRun", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// RestorePlaybookRun indicates an expected call of RestorePlaybookRun
func (mr *MockPlaybookRunServiceMockRecorder) RestorePlaybookRun(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RestorePlaybookRun", reflect.TypeOf((*MockPlaybookRunService)(nil).RestorePlaybookRun), arg0, arg1)
}

// RunChecklistItemSlashCommand mocks base method
func (m *MockPlaybookRunService) RunChecklistItemSlashCommand(arg0, arg1 string, arg2, arg3 int) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RunChecklistItemSlashCommand", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RunChecklistItemSlashCommand indicates an expected call of RunChecklistItemSlashCommand
func (mr *MockPlaybookRunServiceMockRecorder) RunChecklistItemSlashCommand(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunChecklistItemSlashCommand", reflect.TypeOf((*MockPlaybookRunService)(nil).RunChecklistItemSlashCommand), arg0, arg1, arg2, arg3)
}

// SetAssignee mocks base method
func (m *MockPlaybookRunService) SetAssignee(arg0, arg1, arg2 string, arg3, arg4 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetAssignee", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetAssignee indicates an expected call of SetAssignee
func (mr *MockPlaybookRunServiceMockRecorder) SetAssignee(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetAssignee", reflect.TypeOf((*MockPlaybookRunService)(nil).SetAssignee), arg0, arg1, arg2, arg3, arg4)
}

// SetNewReminder mocks base method
func (m *MockPlaybookRunService) SetNewReminder(arg0 string, arg1 time.Duration) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetNewReminder", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetNewReminder indicates an expected call of SetNewReminder
func (mr *MockPlaybookRunServiceMockRecorder) SetNewReminder(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetNewReminder", reflect.TypeOf((*MockPlaybookRunService)(nil).SetNewReminder), arg0, arg1)
}

// SetReminder mocks base method
func (m *MockPlaybookRunService) SetReminder(arg0 string, arg1 time.Duration) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetReminder", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetReminder indicates an expected call of SetReminder
func (mr *MockPlaybookRunServiceMockRecorder) SetReminder(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetReminder", reflect.TypeOf((*MockPlaybookRunService)(nil).SetReminder), arg0, arg1)
}

// SkipChecklistItem mocks base method
func (m *MockPlaybookRunService) SkipChecklistItem(arg0, arg1 string, arg2, arg3 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SkipChecklistItem", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// SkipChecklistItem indicates an expected call of SkipChecklistItem
func (mr *MockPlaybookRunServiceMockRecorder) SkipChecklistItem(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SkipChecklistItem", reflect.TypeOf((*MockPlaybookRunService)(nil).SkipChecklistItem), arg0, arg1, arg2, arg3)
}

// ToggleCheckedState mocks base method
func (m *MockPlaybookRunService) ToggleCheckedState(arg0, arg1 string, arg2, arg3 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ToggleCheckedState", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// ToggleCheckedState indicates an expected call of ToggleCheckedState
func (mr *MockPlaybookRunServiceMockRecorder) ToggleCheckedState(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ToggleCheckedState", reflect.TypeOf((*MockPlaybookRunService)(nil).ToggleCheckedState), arg0, arg1, arg2, arg3)
}

// Unfollow mocks base method
func (m *MockPlaybookRunService) Unfollow(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Unfollow", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Unfollow indicates an expected call of Unfollow
func (mr *MockPlaybookRunServiceMockRecorder) Unfollow(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unfollow", reflect.TypeOf((*MockPlaybookRunService)(nil).Unfollow), arg0, arg1)
}

// UpdateDescription mocks base method
func (m *MockPlaybookRunService) UpdateDescription(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateDescription", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateDescription indicates an expected call of UpdateDescription
func (mr *MockPlaybookRunServiceMockRecorder) UpdateDescription(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateDescription", reflect.TypeOf((*MockPlaybookRunService)(nil).UpdateDescription), arg0, arg1)
}

// UpdateRetrospective mocks base method
func (m *MockPlaybookRunService) UpdateRetrospective(arg0, arg1 string, arg2 app.RetrospectiveUpdate) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateRetrospective", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateRetrospective indicates an expected call of UpdateRetrospective
func (mr *MockPlaybookRunServiceMockRecorder) UpdateRetrospective(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateRetrospective", reflect.TypeOf((*MockPlaybookRunService)(nil).UpdateRetrospective), arg0, arg1, arg2)
}

// UpdateStatus mocks base method
func (m *MockPlaybookRunService) UpdateStatus(arg0, arg1 string, arg2 app.StatusUpdateOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateStatus", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateStatus indicates an expected call of UpdateStatus
func (mr *MockPlaybookRunServiceMockRecorder) UpdateStatus(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateStatus", reflect.TypeOf((*MockPlaybookRunService)(nil).UpdateStatus), arg0, arg1, arg2)
}

// UserHasJoinedChannel mocks base method
func (m *MockPlaybookRunService) UserHasJoinedChannel(arg0, arg1, arg2 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UserHasJoinedChannel", arg0, arg1, arg2)
}

// UserHasJoinedChannel indicates an expected call of UserHasJoinedChannel
func (mr *MockPlaybookRunServiceMockRecorder) UserHasJoinedChannel(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserHasJoinedChannel", reflect.TypeOf((*MockPlaybookRunService)(nil).UserHasJoinedChannel), arg0, arg1, arg2)
}

// UserHasLeftChannel mocks base method
func (m *MockPlaybookRunService) UserHasLeftChannel(arg0, arg1, arg2 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UserHasLeftChannel", arg0, arg1, arg2)
}

// UserHasLeftChannel indicates an expected call of UserHasLeftChannel
func (mr *MockPlaybookRunServiceMockRecorder) UserHasLeftChannel(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserHasLeftChannel", reflect.TypeOf((*MockPlaybookRunService)(nil).UserHasLeftChannel), arg0, arg1, arg2)
}
