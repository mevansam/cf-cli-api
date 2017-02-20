package mock_test

import (
	"sync"

	"code.cloudfoundry.org/cli/cf/api"
	"code.cloudfoundry.org/cli/cf/models"
)

type FakeAppSummaryRepository struct {
	GetSummariesInCurrentSpaceStub        func() (apps []models.Application, apiErr error)
	getSummariesInCurrentSpaceMutex       sync.RWMutex
	getSummariesInCurrentSpaceArgsForCall []struct{}
	getSummariesInCurrentSpaceReturns     struct {
		result1 []models.Application
		result2 error
	}
	GetSummaryStub        func(appGUID string) (summary models.Application, apiErr error)
	getSummaryMutex       sync.RWMutex
	getSummaryArgsForCall []struct {
		appGUID string
	}
	getSummaryReturns struct {
		result1 models.Application
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeAppSummaryRepository) GetSummariesInCurrentSpace() (apps []models.Application, apiErr error) {
	fake.getSummariesInCurrentSpaceMutex.Lock()
	fake.getSummariesInCurrentSpaceArgsForCall = append(fake.getSummariesInCurrentSpaceArgsForCall, struct{}{})
	fake.recordInvocation("GetSummariesInCurrentSpace", []interface{}{})
	fake.getSummariesInCurrentSpaceMutex.Unlock()
	if fake.GetSummariesInCurrentSpaceStub != nil {
		return fake.GetSummariesInCurrentSpaceStub()
	} else {
		return fake.getSummariesInCurrentSpaceReturns.result1, fake.getSummariesInCurrentSpaceReturns.result2
	}
}

func (fake *FakeAppSummaryRepository) GetSummariesInCurrentSpaceCallCount() int {
	fake.getSummariesInCurrentSpaceMutex.RLock()
	defer fake.getSummariesInCurrentSpaceMutex.RUnlock()
	return len(fake.getSummariesInCurrentSpaceArgsForCall)
}

func (fake *FakeAppSummaryRepository) GetSummariesInCurrentSpaceReturns(result1 []models.Application, result2 error) {
	fake.GetSummariesInCurrentSpaceStub = nil
	fake.getSummariesInCurrentSpaceReturns = struct {
		result1 []models.Application
		result2 error
	}{result1, result2}
}

func (fake *FakeAppSummaryRepository) GetSummary(appGUID string) (summary models.Application, apiErr error) {
	fake.getSummaryMutex.Lock()
	fake.getSummaryArgsForCall = append(fake.getSummaryArgsForCall, struct {
		appGUID string
	}{appGUID})
	fake.recordInvocation("GetSummary", []interface{}{appGUID})
	fake.getSummaryMutex.Unlock()
	if fake.GetSummaryStub != nil {
		return fake.GetSummaryStub(appGUID)
	} else {
		return fake.getSummaryReturns.result1, fake.getSummaryReturns.result2
	}
}

func (fake *FakeAppSummaryRepository) GetSummaryCallCount() int {
	fake.getSummaryMutex.RLock()
	defer fake.getSummaryMutex.RUnlock()
	return len(fake.getSummaryArgsForCall)
}

func (fake *FakeAppSummaryRepository) GetSummaryArgsForCall(i int) string {
	fake.getSummaryMutex.RLock()
	defer fake.getSummaryMutex.RUnlock()
	return fake.getSummaryArgsForCall[i].appGUID
}

func (fake *FakeAppSummaryRepository) GetSummaryReturns(result1 models.Application, result2 error) {
	fake.GetSummaryStub = nil
	fake.getSummaryReturns = struct {
		result1 models.Application
		result2 error
	}{result1, result2}
}

func (fake *FakeAppSummaryRepository) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getSummariesInCurrentSpaceMutex.RLock()
	defer fake.getSummariesInCurrentSpaceMutex.RUnlock()
	fake.getSummaryMutex.RLock()
	defer fake.getSummaryMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeAppSummaryRepository) recordInvocation(key string, args []interface{}) {
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

var _ api.AppSummaryRepository = new(FakeAppSummaryRepository)
