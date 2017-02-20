package mock_test

import (
	"os"
	"sync"

	"code.cloudfoundry.org/cli/cf/api/applicationbits"
	"code.cloudfoundry.org/cli/cf/api/resources"
)

type FakeApplicationBitsRepository struct {
	GetApplicationFilesStub        func(appFilesRequest []resources.AppFileResource) ([]resources.AppFileResource, error)
	getApplicationFilesMutex       sync.RWMutex
	getApplicationFilesArgsForCall []struct {
		appFilesRequest []resources.AppFileResource
	}
	getApplicationFilesReturns struct {
		result1 []resources.AppFileResource
		result2 error
	}
	UploadBitsStub        func(appGUID string, zipFile *os.File, presentFiles []resources.AppFileResource) (apiErr error)
	uploadBitsMutex       sync.RWMutex
	uploadBitsArgsForCall []struct {
		appGUID      string
		zipFile      *os.File
		presentFiles []resources.AppFileResource
	}
	uploadBitsReturns struct {
		result1 error
	}
}

func (fake *FakeApplicationBitsRepository) GetApplicationFiles(appFilesRequest []resources.AppFileResource) ([]resources.AppFileResource, error) {
	var appFilesRequestCopy []resources.AppFileResource
	if appFilesRequest != nil {
		appFilesRequestCopy = make([]resources.AppFileResource, len(appFilesRequest))
		copy(appFilesRequestCopy, appFilesRequest)
	}
	fake.getApplicationFilesMutex.Lock()
	fake.getApplicationFilesArgsForCall = append(fake.getApplicationFilesArgsForCall, struct {
		appFilesRequest []resources.AppFileResource
	}{appFilesRequestCopy})
	fake.getApplicationFilesMutex.Unlock()
	if fake.GetApplicationFilesStub != nil {
		return fake.GetApplicationFilesStub(appFilesRequest)
	} else {
		return fake.getApplicationFilesReturns.result1, fake.getApplicationFilesReturns.result2
	}
}

func (fake *FakeApplicationBitsRepository) GetApplicationFilesCallCount() int {
	fake.getApplicationFilesMutex.RLock()
	defer fake.getApplicationFilesMutex.RUnlock()
	return len(fake.getApplicationFilesArgsForCall)
}

func (fake *FakeApplicationBitsRepository) GetApplicationFilesArgsForCall(i int) []resources.AppFileResource {
	fake.getApplicationFilesMutex.RLock()
	defer fake.getApplicationFilesMutex.RUnlock()
	return fake.getApplicationFilesArgsForCall[i].appFilesRequest
}

func (fake *FakeApplicationBitsRepository) GetApplicationFilesReturns(result1 []resources.AppFileResource, result2 error) {
	fake.GetApplicationFilesStub = nil
	fake.getApplicationFilesReturns = struct {
		result1 []resources.AppFileResource
		result2 error
	}{result1, result2}
}

func (fake *FakeApplicationBitsRepository) UploadBits(appGUID string, zipFile *os.File, presentFiles []resources.AppFileResource) (apiErr error) {
	var presentFilesCopy []resources.AppFileResource
	if presentFiles != nil {
		presentFilesCopy = make([]resources.AppFileResource, len(presentFiles))
		copy(presentFilesCopy, presentFiles)
	}
	fake.uploadBitsMutex.Lock()
	fake.uploadBitsArgsForCall = append(fake.uploadBitsArgsForCall, struct {
		appGUID      string
		zipFile      *os.File
		presentFiles []resources.AppFileResource
	}{appGUID, zipFile, presentFilesCopy})
	fake.uploadBitsMutex.Unlock()
	if fake.UploadBitsStub != nil {
		return fake.UploadBitsStub(appGUID, zipFile, presentFiles)
	} else {
		return fake.uploadBitsReturns.result1
	}
}

func (fake *FakeApplicationBitsRepository) UploadBitsCallCount() int {
	fake.uploadBitsMutex.RLock()
	defer fake.uploadBitsMutex.RUnlock()
	return len(fake.uploadBitsArgsForCall)
}

func (fake *FakeApplicationBitsRepository) UploadBitsArgsForCall(i int) (string, *os.File, []resources.AppFileResource) {
	fake.uploadBitsMutex.RLock()
	defer fake.uploadBitsMutex.RUnlock()
	return fake.uploadBitsArgsForCall[i].appGUID, fake.uploadBitsArgsForCall[i].zipFile, fake.uploadBitsArgsForCall[i].presentFiles
}

func (fake *FakeApplicationBitsRepository) UploadBitsReturns(result1 error) {
	fake.UploadBitsStub = nil
	fake.uploadBitsReturns = struct {
		result1 error
	}{result1}
}

var _ applicationbits.Repository = new(FakeApplicationBitsRepository)
