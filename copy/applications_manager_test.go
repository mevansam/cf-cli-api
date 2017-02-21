package copy_test

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"time"

	"code.cloudfoundry.org/cli/cf/api"
	"code.cloudfoundry.org/cli/cf/api/applicationbits"
	"code.cloudfoundry.org/cli/cf/api/applications"
	"code.cloudfoundry.org/cli/cf/api/resources"
	"code.cloudfoundry.org/cli/cf/errors"
	"code.cloudfoundry.org/cli/cf/models"
	"github.com/mevansam/cf-cli-api/cli"
	. "github.com/mevansam/cf-cli-api/cli/mocks"
	"github.com/mevansam/cf-cli-api/copy"
	"github.com/mevansam/cf-cli-api/utils"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Application Manager Tests", func() {

	var (
		err    error
		logger *cli.Logger

		am copy.ApplicationsManager
		sc copy.ServiceCollection
	)

	BeforeEach(func() {
		logger = cli.NewLogger(true, "true")

		am = copy.NewCfCliApplicationsManager()
		sc = &mockServiceCollection{}

		err = am.Init(srcSession, destSession, logger)
		if err != nil {
			Fail(err.Error())
		}
	})
	AfterEach(func() {
	})

	Context("Copy Applications", func() {
		It("Should copy applications from source to destination sessions.", func() {

			srcSession.MockAppSummary = func() api.AppSummaryRepository {
				return &FakeAppSummaryRepository{
					GetSummariesInCurrentSpaceStub: func() (apps []models.Application, apiErr error) {
						return srcApps, nil
					},
				}
			}
			srcSession.MockDownloadAppContent = func(appGUID string, outputFile *os.File, asDroplet bool) (err error) {
				Expect(asDroplet).Should(BeFalse())
				Expect(outputFile).ShouldNot(BeNil())
				_, validAppGUID := utils.ContainsInStrings([]string{appGUID}, []string{srcApps[0].GUID, srcApps[1].GUID})
				Expect(validAppGUID).Should(BeTrue())
				_, err = outputFile.WriteString(srcAppContent[appGUID])
				return
			}
			srcSession.MockDomains = func() api.DomainRepository {
				return &FakeDomainRepository{
					FirstOrDefaultStub: func(orgGUID string, name *string) (models.DomainFields, error) {
						Expect(orgGUID).To(Equal(srcOrg.GUID))
						return srcDefaultDomain, nil
					},
				}
			}
			destSession.MockApplications = func() applications.Repository {
				return &FakeApplicationsRepository{
					CreateStub: func(params models.AppParams) (app models.Application, apiErr error) {
						Expect(params.GUID).Should(BeNil())
						switch {
						case *params.Name == srcApps[0].Name:
							app.GUID = newDestAppGUIDs[0]
						case *params.Name == srcApps[1].Name:
							app.GUID = newDestAppGUIDs[1]
						default:
							Fail("Invalid app name received for download test")
						}
						app.Name = *params.Name
						app.Routes = []models.RouteSummary{}
						destApps[app.GUID] = app
						return
					},
					ReadStub: func(name string) (app models.Application, err error) {
						for _, a := range destApps {
							if name == a.Name {
								app = a
								return
							}
						}
						err = &errors.ModelNotFoundError{}
						return
					},
					UpdateStub: func(appGUID string, params models.AppParams) (app models.Application, err error) {
						_, validAppGUID := utils.ContainsInStrings([]string{appGUID}, newDestAppGUIDs)
						Expect(validAppGUID).Should(BeTrue())
						Expect(*params.State).To(Equal("started"))
						return
					},
					DeleteStub: func(appGUID string) (err error) {
						app := destApps[appGUID]
						Expect(app).ShouldNot(BeNil())
						Expect(app.GUID).To(Equal("app-2000"))
						delete(destApps, appGUID)
						return
					},
				}
			}
			destSession.MockRoutes = func() api.RouteRepository {
				return &FakeRouteRepository{
					FindStub: func(host string, domain models.DomainFields, path string, port int) (route models.Route, err error) {
						for _, r := range destRoutes {
							if host == r.Host && domain.Name == r.Domain.Name {
								route = models.Route{
									GUID:   r.GUID,
									Host:   r.Host,
									Domain: r.Domain,
								}
								return
							}
						}
						err = &errors.ModelNotFoundError{}
						return
					},
					CreateStub: func(host string, domain models.DomainFields, path string, port int, useRandomPort bool) (route models.Route, err error) {
						newRouteGUID := fmt.Sprintf("route-%d", destRouteGUIDCounter)
						destRouteGUIDCounter++
						destRoutes[newRouteGUID] = models.RouteSummary{
							GUID:   newRouteGUID,
							Host:   host,
							Domain: domain,
						}
						route = models.Route{
							GUID:   newRouteGUID,
							Host:   host,
							Domain: domain,
						}
						return
					},
					BindStub: func(routeGUID, appGUID string) (err error) {
						route, exists := destRoutes[routeGUID]
						Expect(exists).Should(BeTrue())
						app, exists := destApps[appGUID]
						Expect(exists).Should(BeTrue())
						app.Routes = append(app.Routes, route)
						destApps[appGUID] = app
						return
					},
					DeleteStub: func(routeGUID string) (err error) {
						_, exists := destRoutes[routeGUID]
						Expect(exists).Should(BeTrue())
						_, validRouteGUID := utils.ContainsInStrings([]string{routeGUID}, []string{"route-2000", "route-2001"})
						Expect(validRouteGUID).Should(BeTrue())
						delete(destRoutes, routeGUID)
						return
					},
				}
			}
			destSession.MockApplicationBits = func() applicationbits.Repository {
				return &FakeApplicationBitsRepository{
					UploadBitsStub: func(appGUID string, zipFile *os.File, presentFiles []resources.AppFileResource) (apiErr error) {
						b := new(bytes.Buffer)
						_, err := io.Copy(b, zipFile)
						if err != nil {
							return err
						}
						Expect(string(b.Bytes())).To(Equal(destAppContent[appGUID]))
						return
					},
				}
			}
			destSession.MockDomains = func() api.DomainRepository {
				return &FakeDomainRepository{
					FirstOrDefaultStub: func(orgGUID string, name *string) (models.DomainFields, error) {
						Expect(orgGUID).To(Equal(destOrg.GUID))
						return destDefaultDomain, nil
					},
					ListDomainsForOrgStub: func(orgGUID string, cb func(models.DomainFields) bool) error {
						for _, d := range destDomains {
							if !cb(d) {
								return nil
							}
						}
						return nil
					},
				}
			}
			destSession.MockServiceBindings = func() api.ServiceBindingRepository {
				return &FakeServiceBindingRepository{
					CreateStub: func(instanceGUID string, appGUID string, paramsMap map[string]interface{}) error {
						app, exists := destApps[appGUID]
						Expect(exists).Should(BeTrue())
						bindings, exists := destAppServiceBindings[app.Name]
						Expect(exists).Should(BeTrue())
						_, svcExists := utils.ContainsInStrings([]string{instanceGUID}, bindings)
						Expect(svcExists).Should(BeTrue())
						return nil
					},
				}
			}

			ac, err := am.ApplicationsToBeCopied([]string{"app1", "app2"}, false)
			if err != nil {
				Fail(err.Error())
			}
			err = am.DoCopy(ac, sc, "", "")
			if err != nil {
				Fail(err.Error())
			}

			app1, exists := destApps["app-2001"]
			Expect(exists).Should(BeTrue())
			Expect(app1.GUID).To(Equal("app-2001"))
			Expect(app1.Name).To(Equal("app1"))
			Expect(len(app1.Routes)).To(Equal(2))
			Expect(app1.Routes[0].GUID).To(Equal("route-2002"))
			Expect(app1.Routes[0].URL()).To(Equal("app1.acme-dest.com"))
			Expect(app1.Routes[1].GUID).To(Equal("route-2003"))
			Expect(app1.Routes[1].URL()).To(Equal("foo1.acme-test.com"))

			app2, exists := destApps["app-2002"]
			Expect(exists).Should(BeTrue())
			Expect(app2.GUID).To(Equal("app-2002"))
			Expect(app2.Name).To(Equal("app2"))
			Expect(len(app2.Routes)).To(Equal(1))
			Expect(app2.Routes[0].GUID).To(Equal("route-2004"))
			Expect(app2.Routes[0].URL()).To(Equal("app2.acme-dest.com"))
		})
	})
})

// mockServiceCollection -
type mockServiceCollection struct {
}

// AppBindings -
func (sc mockServiceCollection) AppBindings(appName string) (bindings []string, ok bool) {
	bindings, ok = destAppServiceBindings[appName]
	return
}

// Test Model

var timestamp = time.Now().Format(time.RFC3339)

var srcApps = []models.Application{
	models.Application{
		ApplicationFields: models.ApplicationFields{
			Name: "app1",
			GUID: "app-1000",
		},
		Routes: []models.RouteSummary{
			srcRoutes["route-1000"],
			srcRoutes["route-1002"],
		},
	},
	models.Application{
		ApplicationFields: models.ApplicationFields{
			Name: "app2",
			GUID: "app-1001",
		},
		Routes: []models.RouteSummary{
			srcRoutes["route-1001"],
		},
	},
}
var srcRoutes = map[string]models.RouteSummary{
	"route-1000": models.RouteSummary{
		GUID:   "route-1000",
		Host:   "app1",
		Domain: srcDefaultDomain,
	},
	"route-1001": models.RouteSummary{
		GUID:   "route-1001",
		Host:   "app2",
		Domain: srcDefaultDomain,
	},
	"route-1002": models.RouteSummary{
		GUID:   "route-1002",
		Host:   "foo1",
		Domain: srcDomains["domain-1001"],
	},
}
var srcDomains = map[string]models.DomainFields{
	"domain-1000": models.DomainFields{
		GUID: "domain-1000",
		Name: "acme-src.com",
	},
	"domain-1001": models.DomainFields{
		GUID: "domain-1001",
		Name: "acme-test.com",
	},
}
var srcDefaultDomain = srcDomains["domain-1000"]

var destApps = map[string]models.Application{
	"app-2000": models.Application{
		ApplicationFields: models.ApplicationFields{
			Name: "app1",
			GUID: "app-2000",
		},
		Routes: []models.RouteSummary{
			destRoutes["route-2000"],
		},
	},
}
var destRoutes = map[string]models.RouteSummary{
	"route-2000": models.RouteSummary{
		GUID:   "route-2000",
		Host:   "app1",
		Domain: destDefaultDomain,
	},
	"route-2001": models.RouteSummary{
		GUID:   "route-2001",
		Host:   "foo1",
		Domain: destDomains["domain-2001"],
	},
}
var destDomains = map[string]models.DomainFields{
	"domain-2000": models.DomainFields{
		GUID: "domain-2000",
		Name: "acme-dest.com",
	},
	"domain-2001": models.DomainFields{
		GUID: "domain-2001",
		Name: "acme-test.com",
	},
}
var destDefaultDomain = destDomains["domain-2000"]

var newDestAppGUIDs = []string{"app-2001", "app-2002"}
var destRouteGUIDCounter = 2002

var srcAppContent = map[string]string{
	"app-1000": fmt.Sprintf("application bits content for app1: %s", timestamp),
	"app-1001": fmt.Sprintf("application bits content for app3: %s", timestamp),
}
var destAppContent = map[string]string{
	newDestAppGUIDs[0]: srcAppContent["app-1000"],
	newDestAppGUIDs[1]: srcAppContent["app-1001"],
}

var destAppServiceBindings = map[string][]string{
	"app1": []string{"svc-2000", "svc-2001"},
	"app2": []string{"svc-2001", "svc-2002", "svc-2003"},
}
