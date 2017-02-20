package copy_test

import (
	. "github.com/mevansam/cf-cli-api/cli/mocks"
	. "github.com/onsi/ginkgo"
	// . "github.com/onsi/gomega"
)

var _ = Describe("Service Manager Tests", func() {

	var (
		srcSession  *MockSession
		destSession *MockSession
	)

	BeforeEach(func() {
		srcSession = &MockSession{}
		destSession = &MockSession{}
	})

	Context("Test AppBits", func() {
	})
	Context("Test AppDroplet", func() {
	})
})
