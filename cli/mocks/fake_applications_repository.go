package mock_test

import (
	"sync"

	"code.cloudfoundry.org/cli/cf/api/applications"
	"code.cloudfoundry.org/cli/cf/models"
)

type FakeApplicationsRepository struct {
	CreateStub        func(params models.AppParams) (createdApp models.Application, apiErr error)
	createMutex       sync.RWMutex
	createArgsForCall []struct {
		params models.AppParams
	}
	createReturns struct {
		result1 models.Application
		result2 error
	}
	GetAppStub        func(appGUID string) (models.Application, error)
	getAppMutex       sync.RWMutex
	getAppArgsForCall []struct {
		appGUID string
	}
	getAppReturns struct {
		result1 models.Application
		result2 error
	}
	ReadStub        func(name string) (app models.Application, apiErr error)
	readMutex       sync.RWMutex
	readArgsForCall []struct {
		name string
	}
	readReturns struct {
		result1 models.Application
		result2 error
	}
	ReadFromSpaceStub        func(name string, spaceGUID string) (app models.Application, apiErr error)
	readFromSpaceMutex       sync.RWMutex
	readFromSpaceArgsForCall []struct {
		name      string
		spaceGUID string
	}
	readFromSpaceReturns struct {
		result1 models.Application
		result2 error
	}
	UpdateStub        func(appGUID string, params models.AppParams) (updatedApp models.Application, apiErr error)
	updateMutex       sync.RWMutex
	updateArgsForCall []struct {
		appGUID string
		params  models.AppParams
	}
	updateReturns struct {
		result1 models.Application
		result2 error
	}
	DeleteStub        func(appGUID string) (apiErr error)
	deleteMutex       sync.RWMutex
	deleteArgsForCall []struct {
		appGUID string
	}
	deleteReturns struct {
		result1 error
	}
	ReadEnvStub        func(guid string) (*models.Environment, error)
	readEnvMutex       sync.RWMutex
	readEnvArgsForCall []struct {
		guid string
	}
	readEnvReturns struct {
		result1 *models.Environment
		result2 error
	}
	CreateRestageRequestStub        func(guid string) (apiErr error)
	createRestageRequestMutex       sync.RWMutex
	createRestageRequestArgsForCall []struct {
		guid string
	}
	createRestageRequestReturns struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeApplicationsRepository) Create(params models.AppParams) (createdApp models.Application, apiErr error) {
	fake.createMutex.Lock()
	fake.createArgsForCall = append(fake.createArgsForCall, struct {
		params models.AppParams
	}{params})
	fake.recordInvocation("Create", []interface{}{params})
	fake.createMutex.Unlock()
	if fake.CreateStub != nil {
		return fake.CreateStub(params)
	} else {
		return fake.createReturns.result1, fake.createReturns.result2
	}
}

func (fake *FakeApplicationsRepository) CreateCallCount() int {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return len(fake.createArgsForCall)
}

func (fake *FakeApplicationsRepository) CreateArgsForCall(i int) models.AppParams {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return fake.createArgsForCall[i].params
}

func (fake *FakeApplicationsRepository) CreateReturns(result1 models.Application, result2 error) {
	fake.CreateStub = nil
	fake.createReturns = struct {
		result1 models.Application
		result2 error
	}{result1, result2}
}

func (fake *FakeApplicationsRepository) GetApp(appGUID string) (models.Application, error) {
	fake.getAppMutex.Lock()
	fake.getAppArgsForCall = append(fake.getAppArgsForCall, struct {
		appGUID string
	}{appGUID})
	fake.recordInvocation("GetApp", []interface{}{appGUID})
	fake.getAppMutex.Unlock()
	if fake.GetAppStub != nil {
		return fake.GetAppStub(appGUID)
	} else {
		return fake.getAppReturns.result1, fake.getAppReturns.result2
	}
}

func (fake *FakeApplicationsRepository) GetAppCallCount() int {
	fake.getAppMutex.RLock()
	defer fake.getAppMutex.RUnlock()
	return len(fake.getAppArgsForCall)
}

func (fake *FakeApplicationsRepository) GetAppArgsForCall(i int) string {
	fake.getAppMutex.RLock()
	defer fake.getAppMutex.RUnlock()
	return fake.getAppArgsForCall[i].appGUID
}

func (fake *FakeApplicationsRepository) GetAppReturns(result1 models.Application, result2 error) {
	fake.GetAppStub = nil
	fake.getAppReturns = struct {
		result1 models.Application
		result2 error
	}{result1, result2}
}

func (fake *FakeApplicationsRepository) Read(name string) (app models.Application, apiErr error) {
	fake.readMutex.Lock()
	fake.readArgsForCall = append(fake.readArgsForCall, struct {
		name string
	}{name})
	fake.recordInvocation("Read", []interface{}{name})
	fake.readMutex.Unlock()
	if fake.ReadStub != nil {
		return fake.ReadStub(name)
	} else {
		return fake.readReturns.result1, fake.readReturns.result2
	}
}

func (fake *FakeApplicationsRepository) ReadCallCount() int {
	fake.readMutex.RLock()
	defer fake.readMutex.RUnlock()
	return len(fake.readArgsForCall)
}

func (fake *FakeApplicationsRepository) ReadArgsForCall(i int) string {
	fake.readMutex.RLock()
	defer fake.readMutex.RUnlock()
	return fake.readArgsForCall[i].name
}

func (fake *FakeApplicationsRepository) ReadReturns(result1 models.Application, result2 error) {
	fake.ReadStub = nil
	fake.readReturns = struct {
		result1 models.Application
		result2 error
	}{result1, result2}
}

func (fake *FakeApplicationsRepository) ReadFromSpace(name string, spaceGUID string) (app models.Application, apiErr error) {
	fake.readFromSpaceMutex.Lock()
	fake.readFromSpaceArgsForCall = append(fake.readFromSpaceArgsForCall, struct {
		name      string
		spaceGUID string
	}{name, spaceGUID})
	fake.recordInvocation("ReadFromSpace", []interface{}{name, spaceGUID})
	fake.readFromSpaceMutex.Unlock()
	if fake.ReadFromSpaceStub != nil {
		return fake.ReadFromSpaceStub(name, spaceGUID)
	} else {
		return fake.readFromSpaceReturns.result1, fake.readFromSpaceReturns.result2
	}
}

func (fake *FakeApplicationsRepository) ReadFromSpaceCallCount() int {
	fake.readFromSpaceMutex.RLock()
	defer fake.readFromSpaceMutex.RUnlock()
	return len(fake.readFromSpaceArgsForCall)
}

func (fake *FakeApplicationsRepository) ReadFromSpaceArgsForCall(i int) (string, string) {
	fake.readFromSpaceMutex.RLock()
	defer fake.readFromSpaceMutex.RUnlock()
	return fake.readFromSpaceArgsForCall[i].name, fake.readFromSpaceArgsForCall[i].spaceGUID
}

func (fake *FakeApplicationsRepository) ReadFromSpaceReturns(result1 models.Application, result2 error) {
	fake.ReadFromSpaceStub = nil
	fake.readFromSpaceReturns = struct {
		result1 models.Application
		result2 error
	}{result1, result2}
}

func (fake *FakeApplicationsRepository) Update(appGUID string, params models.AppParams) (updatedApp models.Application, apiErr error) {
	fake.updateMutex.Lock()
	fake.updateArgsForCall = append(fake.updateArgsForCall, struct {
		appGUID string
		params  models.AppParams
	}{appGUID, params})
	fake.recordInvocation("Update", []interface{}{appGUID, params})
	fake.updateMutex.Unlock()
	if fake.UpdateStub != nil {
		return fake.UpdateStub(appGUID, params)
	} else {
		return fake.updateReturns.result1, fake.updateReturns.result2
	}
}

func (fake *FakeApplicationsRepository) UpdateCallCount() int {
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	return len(fake.updateArgsForCall)
}

func (fake *FakeApplicationsRepository) UpdateArgsForCall(i int) (string, models.AppParams) {
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	return fake.updateArgsForCall[i].appGUID, fake.updateArgsForCall[i].params
}

func (fake *FakeApplicationsRepository) UpdateReturns(result1 models.Application, result2 error) {
	fake.UpdateStub = nil
	fake.updateReturns = struct {
		result1 models.Application
		result2 error
	}{result1, result2}
}

func (fake *FakeApplicationsRepository) Delete(appGUID string) (apiErr error) {
	fake.deleteMutex.Lock()
	fake.deleteArgsForCall = append(fake.deleteArgsForCall, struct {
		appGUID string
	}{appGUID})
	fake.recordInvocation("Delete", []interface{}{appGUID})
	fake.deleteMutex.Unlock()
	if fake.DeleteStub != nil {
		return fake.DeleteStub(appGUID)
	} else {
		return fake.deleteReturns.result1
	}
}

func (fake *FakeApplicationsRepository) DeleteCallCount() int {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return len(fake.deleteArgsForCall)
}

func (fake *FakeApplicationsRepository) DeleteArgsForCall(i int) string {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return fake.deleteArgsForCall[i].appGUID
}

func (fake *FakeApplicationsRepository) DeleteReturns(result1 error) {
	fake.DeleteStub = nil
	fake.deleteReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeApplicationsRepository) ReadEnv(guid string) (*models.Environment, error) {
	fake.readEnvMutex.Lock()
	fake.readEnvArgsForCall = append(fake.readEnvArgsForCall, struct {
		guid string
	}{guid})
	fake.recordInvocation("ReadEnv", []interface{}{guid})
	fake.readEnvMutex.Unlock()
	if fake.ReadEnvStub != nil {
		return fake.ReadEnvStub(guid)
	} else {
		return fake.readEnvReturns.result1, fake.readEnvReturns.result2
	}
}

func (fake *FakeApplicationsRepository) ReadEnvCallCount() int {
	fake.readEnvMutex.RLock()
	defer fake.readEnvMutex.RUnlock()
	return len(fake.readEnvArgsForCall)
}

func (fake *FakeApplicationsRepository) ReadEnvArgsForCall(i int) string {
	fake.readEnvMutex.RLock()
	defer fake.readEnvMutex.RUnlock()
	return fake.readEnvArgsForCall[i].guid
}

func (fake *FakeApplicationsRepository) ReadEnvReturns(result1 *models.Environment, result2 error) {
	fake.ReadEnvStub = nil
	fake.readEnvReturns = struct {
		result1 *models.Environment
		result2 error
	}{result1, result2}
}

func (fake *FakeApplicationsRepository) CreateRestageRequest(guid string) (apiErr error) {
	fake.createRestageRequestMutex.Lock()
	fake.createRestageRequestArgsForCall = append(fake.createRestageRequestArgsForCall, struct {
		guid string
	}{guid})
	fake.recordInvocation("CreateRestageRequest", []interface{}{guid})
	fake.createRestageRequestMutex.Unlock()
	if fake.CreateRestageRequestStub != nil {
		return fake.CreateRestageRequestStub(guid)
	} else {
		return fake.createRestageRequestReturns.result1
	}
}

func (fake *FakeApplicationsRepository) CreateRestageRequestCallCount() int {
	fake.createRestageRequestMutex.RLock()
	defer fake.createRestageRequestMutex.RUnlock()
	return len(fake.createRestageRequestArgsForCall)
}

func (fake *FakeApplicationsRepository) CreateRestageRequestArgsForCall(i int) string {
	fake.createRestageRequestMutex.RLock()
	defer fake.createRestageRequestMutex.RUnlock()
	return fake.createRestageRequestArgsForCall[i].guid
}

func (fake *FakeApplicationsRepository) CreateRestageRequestReturns(result1 error) {
	fake.CreateRestageRequestStub = nil
	fake.createRestageRequestReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeApplicationsRepository) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	fake.getAppMutex.RLock()
	defer fake.getAppMutex.RUnlock()
	fake.readMutex.RLock()
	defer fake.readMutex.RUnlock()
	fake.readFromSpaceMutex.RLock()
	defer fake.readFromSpaceMutex.RUnlock()
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	fake.readEnvMutex.RLock()
	defer fake.readEnvMutex.RUnlock()
	fake.createRestageRequestMutex.RLock()
	defer fake.createRestageRequestMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeApplicationsRepository) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ applications.Repository = new(FakeApplicationsRepository)
