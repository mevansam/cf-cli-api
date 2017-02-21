package copy_test

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"code.cloudfoundry.org/cli/cf/api/applicationbits"
	"code.cloudfoundry.org/cli/cf/api/applications"
	"code.cloudfoundry.org/cli/cf/api/resources"
	"code.cloudfoundry.org/cli/cf/models"

	. "github.com/mevansam/cf-cli-api/cfapi/mocks"
	. "github.com/mevansam/cf-cli-api/copy"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Application Tests", func() {

	var (
		err    error
		tmpDir string

		session *MockSession

		provider *ApplicationCopier
	)

	BeforeEach(func() {
		session = &MockSession{}
		provider = &ApplicationCopier{}

		tmpDir, err = ioutil.TempDir("", "")
		os.MkdirAll(tmpDir, os.ModePerm)
		if err != nil {
			panic(err.Error())
		}
	})

	AfterEach(func() {
		os.RemoveAll(tmpDir)
	})

	Context("Test AppBits", func() {

		var (
			appContent ApplicationContent
			appModel   *models.Application
		)

		appModel = &models.Application{}

		appModel.GUID = "1234"
		appModel.Name = "testApp1"

		It("Downloads and Uploads application bits", func() {

			session.MockApplications = func() applications.Repository {
				return &FakeApplicationsRepository{
					CreateStub: func(params models.AppParams) (createdApp models.Application, apiErr error) {
						Expect(params.GUID).Should(BeNil())
						Expect(*params.Name).To(Equal("testApp1"))
						createdApp.GUID = "abcd"
						createdApp.Name = *params.Name
						return
					},
				}
			}
			session.MockApplicationBits = func() applicationbits.Repository {
				return &FakeApplicationBitsRepository{
					UploadBitsStub: func(appGUID string, zipFile *os.File, presentFiles []resources.AppFileResource) (apiErr error) {
						Expect(appGUID).To(Equal("abcd"))
						b := new(bytes.Buffer)
						_, err := io.Copy(b, zipFile)
						if err != nil {
							return err
						}
						Expect(string(b.Bytes())).To(Equal("application bits content"))
						return
					},
				}
			}
			session.MockDownloadAppContent = func(appGUID string, outputFile *os.File, asDroplet bool) error {
				Expect(appGUID).To(Equal("1234"))
				Expect(asDroplet).Should(BeFalse())
				Expect(outputFile).ShouldNot(BeNil())
				_, err := outputFile.WriteString("application bits content")
				if err != nil {
					return err
				}
				return nil
			}

			appContent = provider.NewApplication(appModel, tmpDir, false)

			err := appContent.Download(session)
			if err != nil {
				fmt.Println(err.Error())
				panic(err)
			}

			params := appContent.App().ToParams()
			params.GUID = nil

			destApp, err := appContent.Upload(session, params)
			if err != nil {
				fmt.Println(err.Error())
				panic(err)
			}

			Expect(destApp.GUID).To(Equal("abcd"))
		})
	})
	Context("Test AppDroplet", func() {

		var (
			appContent ApplicationContent
			appModel   *models.Application
		)

		appModel = &models.Application{}

		appModel.GUID = "6789"
		appModel.Name = "testApp2"

		It("Downloads and Uploads application bits", func() {

			session.MockApplications = func() applications.Repository {
				return &FakeApplicationsRepository{
					CreateStub: func(params models.AppParams) (createdApp models.Application, apiErr error) {
						Expect(params.GUID).Should(BeNil())
						Expect(*params.Name).To(Equal("testApp2"))
						createdApp.GUID = "wxyz"
						createdApp.Name = *params.Name
						return
					},
				}
			}
			session.MockDownloadAppContent = func(appGUID string, outputFile *os.File, asDroplet bool) error {
				Expect(appGUID).To(Equal("6789"))
				Expect(asDroplet).Should(BeTrue())
				Expect(outputFile).ShouldNot(BeNil())
				_, err := outputFile.WriteString("application droplet contents")
				if err != nil {
					return err
				}
				return nil
			}
			session.MockUploadDroplet = func(appGUID string, contentType string, dropletUploadRequest *os.File) error {
				Expect(appGUID).To(Equal("wxyz"))
				Expect(contentType).Should(MatchRegexp("multipart/form-data; boundary=[a-f0-9]+"))
				return nil
			}

			appContent = provider.NewApplication(appModel, tmpDir, true)

			err := appContent.Download(session)
			if err != nil {
				fmt.Println(err.Error())
				panic(err)
			}

			params := appContent.App().ToParams()
			params.GUID = nil

			destApp, err := appContent.Upload(session, params)
			if err != nil {
				fmt.Println(err.Error())
				panic(err)
			}

			Expect(destApp.GUID).To(Equal("wxyz"))
		})
	})
})
