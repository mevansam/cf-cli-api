package mock_test

import (
	"sync"

	"code.cloudfoundry.org/cli/cf/api/spaces"
	"code.cloudfoundry.org/cli/cf/models"
)

type FakeSpaceRepository struct {
	ListSpacesStub        func(func(models.Space) bool) error
	listSpacesMutex       sync.RWMutex
	listSpacesArgsForCall []struct {
		arg1 func(models.Space) bool
	}
	listSpacesReturns struct {
		result1 error
	}
	ListSpacesFromOrgStub        func(orgGUID string, spaceFunc func(models.Space) bool) error
	listSpacesFromOrgMutex       sync.RWMutex
	listSpacesFromOrgArgsForCall []struct {
		orgGUID   string
		spaceFunc func(models.Space) bool
	}
	listSpacesFromOrgReturns struct {
		result1 error
	}
	FindByNameStub        func(name string) (space models.Space, apiErr error)
	findByNameMutex       sync.RWMutex
	findByNameArgsForCall []struct {
		name string
	}
	findByNameReturns struct {
		result1 models.Space
		result2 error
	}
	FindByNameInOrgStub        func(name, orgGUID string) (space models.Space, apiErr error)
	findByNameInOrgMutex       sync.RWMutex
	findByNameInOrgArgsForCall []struct {
		name    string
		orgGUID string
	}
	findByNameInOrgReturns struct {
		result1 models.Space
		result2 error
	}
	CreateStub        func(name string, orgGUID string, spaceQuotaGUID string) (space models.Space, apiErr error)
	createMutex       sync.RWMutex
	createArgsForCall []struct {
		name           string
		orgGUID        string
		spaceQuotaGUID string
	}
	createReturns struct {
		result1 models.Space
		result2 error
	}
	RenameStub        func(spaceGUID, newName string) (apiErr error)
	renameMutex       sync.RWMutex
	renameArgsForCall []struct {
		spaceGUID string
		newName   string
	}
	renameReturns struct {
		result1 error
	}
	SetAllowSSHStub        func(spaceGUID string, allow bool) (apiErr error)
	setAllowSSHMutex       sync.RWMutex
	setAllowSSHArgsForCall []struct {
		spaceGUID string
		allow     bool
	}
	setAllowSSHReturns struct {
		result1 error
	}
	DeleteStub        func(spaceGUID string) (apiErr error)
	deleteMutex       sync.RWMutex
	deleteArgsForCall []struct {
		spaceGUID string
	}
	deleteReturns struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeSpaceRepository) ListSpaces(arg1 func(models.Space) bool) error {
	fake.listSpacesMutex.Lock()
	fake.listSpacesArgsForCall = append(fake.listSpacesArgsForCall, struct {
		arg1 func(models.Space) bool
	}{arg1})
	fake.recordInvocation("ListSpaces", []interface{}{arg1})
	fake.listSpacesMutex.Unlock()
	if fake.ListSpacesStub != nil {
		return fake.ListSpacesStub(arg1)
	} else {
		return fake.listSpacesReturns.result1
	}
}

func (fake *FakeSpaceRepository) ListSpacesCallCount() int {
	fake.listSpacesMutex.RLock()
	defer fake.listSpacesMutex.RUnlock()
	return len(fake.listSpacesArgsForCall)
}

func (fake *FakeSpaceRepository) ListSpacesArgsForCall(i int) func(models.Space) bool {
	fake.listSpacesMutex.RLock()
	defer fake.listSpacesMutex.RUnlock()
	return fake.listSpacesArgsForCall[i].arg1
}

func (fake *FakeSpaceRepository) ListSpacesReturns(result1 error) {
	fake.ListSpacesStub = nil
	fake.listSpacesReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeSpaceRepository) ListSpacesFromOrg(orgGUID string, spaceFunc func(models.Space) bool) error {
	fake.listSpacesFromOrgMutex.Lock()
	fake.listSpacesFromOrgArgsForCall = append(fake.listSpacesFromOrgArgsForCall, struct {
		orgGUID   string
		spaceFunc func(models.Space) bool
	}{orgGUID, spaceFunc})
	fake.recordInvocation("ListSpacesFromOrg", []interface{}{orgGUID, spaceFunc})
	fake.listSpacesFromOrgMutex.Unlock()
	if fake.ListSpacesFromOrgStub != nil {
		return fake.ListSpacesFromOrgStub(orgGUID, spaceFunc)
	} else {
		return fake.listSpacesFromOrgReturns.result1
	}
}

func (fake *FakeSpaceRepository) ListSpacesFromOrgCallCount() int {
	fake.listSpacesFromOrgMutex.RLock()
	defer fake.listSpacesFromOrgMutex.RUnlock()
	return len(fake.listSpacesFromOrgArgsForCall)
}

func (fake *FakeSpaceRepository) ListSpacesFromOrgArgsForCall(i int) (string, func(models.Space) bool) {
	fake.listSpacesFromOrgMutex.RLock()
	defer fake.listSpacesFromOrgMutex.RUnlock()
	return fake.listSpacesFromOrgArgsForCall[i].orgGUID, fake.listSpacesFromOrgArgsForCall[i].spaceFunc
}

func (fake *FakeSpaceRepository) ListSpacesFromOrgReturns(result1 error) {
	fake.ListSpacesFromOrgStub = nil
	fake.listSpacesFromOrgReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeSpaceRepository) FindByName(name string) (space models.Space, apiErr error) {
	fake.findByNameMutex.Lock()
	fake.findByNameArgsForCall = append(fake.findByNameArgsForCall, struct {
		name string
	}{name})
	fake.recordInvocation("FindByName", []interface{}{name})
	fake.findByNameMutex.Unlock()
	if fake.FindByNameStub != nil {
		return fake.FindByNameStub(name)
	} else {
		return fake.findByNameReturns.result1, fake.findByNameReturns.result2
	}
}

func (fake *FakeSpaceRepository) FindByNameCallCount() int {
	fake.findByNameMutex.RLock()
	defer fake.findByNameMutex.RUnlock()
	return len(fake.findByNameArgsForCall)
}

func (fake *FakeSpaceRepository) FindByNameArgsForCall(i int) string {
	fake.findByNameMutex.RLock()
	defer fake.findByNameMutex.RUnlock()
	return fake.findByNameArgsForCall[i].name
}

func (fake *FakeSpaceRepository) FindByNameReturns(result1 models.Space, result2 error) {
	fake.FindByNameStub = nil
	fake.findByNameReturns = struct {
		result1 models.Space
		result2 error
	}{result1, result2}
}

func (fake *FakeSpaceRepository) FindByNameInOrg(name string, orgGUID string) (space models.Space, apiErr error) {
	fake.findByNameInOrgMutex.Lock()
	fake.findByNameInOrgArgsForCall = append(fake.findByNameInOrgArgsForCall, struct {
		name    string
		orgGUID string
	}{name, orgGUID})
	fake.recordInvocation("FindByNameInOrg", []interface{}{name, orgGUID})
	fake.findByNameInOrgMutex.Unlock()
	if fake.FindByNameInOrgStub != nil {
		return fake.FindByNameInOrgStub(name, orgGUID)
	} else {
		return fake.findByNameInOrgReturns.result1, fake.findByNameInOrgReturns.result2
	}
}

func (fake *FakeSpaceRepository) FindByNameInOrgCallCount() int {
	fake.findByNameInOrgMutex.RLock()
	defer fake.findByNameInOrgMutex.RUnlock()
	return len(fake.findByNameInOrgArgsForCall)
}

func (fake *FakeSpaceRepository) FindByNameInOrgArgsForCall(i int) (string, string) {
	fake.findByNameInOrgMutex.RLock()
	defer fake.findByNameInOrgMutex.RUnlock()
	return fake.findByNameInOrgArgsForCall[i].name, fake.findByNameInOrgArgsForCall[i].orgGUID
}

func (fake *FakeSpaceRepository) FindByNameInOrgReturns(result1 models.Space, result2 error) {
	fake.FindByNameInOrgStub = nil
	fake.findByNameInOrgReturns = struct {
		result1 models.Space
		result2 error
	}{result1, result2}
}

func (fake *FakeSpaceRepository) Create(name string, orgGUID string, spaceQuotaGUID string) (space models.Space, apiErr error) {
	fake.createMutex.Lock()
	fake.createArgsForCall = append(fake.createArgsForCall, struct {
		name           string
		orgGUID        string
		spaceQuotaGUID string
	}{name, orgGUID, spaceQuotaGUID})
	fake.recordInvocation("Create", []interface{}{name, orgGUID, spaceQuotaGUID})
	fake.createMutex.Unlock()
	if fake.CreateStub != nil {
		return fake.CreateStub(name, orgGUID, spaceQuotaGUID)
	} else {
		return fake.createReturns.result1, fake.createReturns.result2
	}
}

func (fake *FakeSpaceRepository) CreateCallCount() int {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return len(fake.createArgsForCall)
}

func (fake *FakeSpaceRepository) CreateArgsForCall(i int) (string, string, string) {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return fake.createArgsForCall[i].name, fake.createArgsForCall[i].orgGUID, fake.createArgsForCall[i].spaceQuotaGUID
}

func (fake *FakeSpaceRepository) CreateReturns(result1 models.Space, result2 error) {
	fake.CreateStub = nil
	fake.createReturns = struct {
		result1 models.Space
		result2 error
	}{result1, result2}
}

func (fake *FakeSpaceRepository) Rename(spaceGUID string, newName string) (apiErr error) {
	fake.renameMutex.Lock()
	fake.renameArgsForCall = append(fake.renameArgsForCall, struct {
		spaceGUID string
		newName   string
	}{spaceGUID, newName})
	fake.recordInvocation("Rename", []interface{}{spaceGUID, newName})
	fake.renameMutex.Unlock()
	if fake.RenameStub != nil {
		return fake.RenameStub(spaceGUID, newName)
	} else {
		return fake.renameReturns.result1
	}
}

func (fake *FakeSpaceRepository) RenameCallCount() int {
	fake.renameMutex.RLock()
	defer fake.renameMutex.RUnlock()
	return len(fake.renameArgsForCall)
}

func (fake *FakeSpaceRepository) RenameArgsForCall(i int) (string, string) {
	fake.renameMutex.RLock()
	defer fake.renameMutex.RUnlock()
	return fake.renameArgsForCall[i].spaceGUID, fake.renameArgsForCall[i].newName
}

func (fake *FakeSpaceRepository) RenameReturns(result1 error) {
	fake.RenameStub = nil
	fake.renameReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeSpaceRepository) SetAllowSSH(spaceGUID string, allow bool) (apiErr error) {
	fake.setAllowSSHMutex.Lock()
	fake.setAllowSSHArgsForCall = append(fake.setAllowSSHArgsForCall, struct {
		spaceGUID string
		allow     bool
	}{spaceGUID, allow})
	fake.recordInvocation("SetAllowSSH", []interface{}{spaceGUID, allow})
	fake.setAllowSSHMutex.Unlock()
	if fake.SetAllowSSHStub != nil {
		return fake.SetAllowSSHStub(spaceGUID, allow)
	} else {
		return fake.setAllowSSHReturns.result1
	}
}

func (fake *FakeSpaceRepository) SetAllowSSHCallCount() int {
	fake.setAllowSSHMutex.RLock()
	defer fake.setAllowSSHMutex.RUnlock()
	return len(fake.setAllowSSHArgsForCall)
}

func (fake *FakeSpaceRepository) SetAllowSSHArgsForCall(i int) (string, bool) {
	fake.setAllowSSHMutex.RLock()
	defer fake.setAllowSSHMutex.RUnlock()
	return fake.setAllowSSHArgsForCall[i].spaceGUID, fake.setAllowSSHArgsForCall[i].allow
}

func (fake *FakeSpaceRepository) SetAllowSSHReturns(result1 error) {
	fake.SetAllowSSHStub = nil
	fake.setAllowSSHReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeSpaceRepository) Delete(spaceGUID string) (apiErr error) {
	fake.deleteMutex.Lock()
	fake.deleteArgsForCall = append(fake.deleteArgsForCall, struct {
		spaceGUID string
	}{spaceGUID})
	fake.recordInvocation("Delete", []interface{}{spaceGUID})
	fake.deleteMutex.Unlock()
	if fake.DeleteStub != nil {
		return fake.DeleteStub(spaceGUID)
	} else {
		return fake.deleteReturns.result1
	}
}

func (fake *FakeSpaceRepository) DeleteCallCount() int {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return len(fake.deleteArgsForCall)
}

func (fake *FakeSpaceRepository) DeleteArgsForCall(i int) string {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return fake.deleteArgsForCall[i].spaceGUID
}

func (fake *FakeSpaceRepository) DeleteReturns(result1 error) {
	fake.DeleteStub = nil
	fake.deleteReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeSpaceRepository) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.listSpacesMutex.RLock()
	defer fake.listSpacesMutex.RUnlock()
	fake.listSpacesFromOrgMutex.RLock()
	defer fake.listSpacesFromOrgMutex.RUnlock()
	fake.findByNameMutex.RLock()
	defer fake.findByNameMutex.RUnlock()
	fake.findByNameInOrgMutex.RLock()
	defer fake.findByNameInOrgMutex.RUnlock()
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	fake.renameMutex.RLock()
	defer fake.renameMutex.RUnlock()
	fake.setAllowSSHMutex.RLock()
	defer fake.setAllowSSHMutex.RUnlock()
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeSpaceRepository) recordInvocation(key string, args []interface{}) {
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

var _ spaces.SpaceRepository = new(FakeSpaceRepository)
