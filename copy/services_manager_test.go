package copy_test

import (
	"fmt"

	"code.cloudfoundry.org/cli/cf/api"
	"code.cloudfoundry.org/cli/cf/api/applications"
	"code.cloudfoundry.org/cli/cf/errors"
	"code.cloudfoundry.org/cli/cf/models"

	"github.com/mevansam/cf-cli-api/cli"
	. "github.com/mevansam/cf-cli-api/cli/mocks"
	"github.com/mevansam/cf-cli-api/copy"
	"github.com/mevansam/cf-cli-api/utils"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Service Manager Tests", func() {

	var (
		err    error
		logger *cli.Logger

		sm copy.ServicesManager
		sc copy.ServiceCollection
	)

	BeforeEach(func() {
		logger = cli.NewLogger(true, "true")

		sm = copy.NewCfCliServicesManager()
		sc = &mockServiceCollection{}

		serviceKeyFormat := "__%s_copy_for_" + fmt.Sprintf("/%s/%s/%s", "destTarget", "destOrg", "destSpace")
		err = sm.Init(srcSession, destSession, serviceKeyFormat, logger)
		if err != nil {
			Fail(err.Error())
		}
	})
	AfterEach(func() {
	})

	Context("Copy Services", func() {
		It("Should copy services from source to destination sessions.", func() {

			srcSession.MockUserProvidedServices = func() api.UserProvidedServiceInstanceRepository {
				return &FakeUserProvidedServiceInstanceRepository{
					GetSummariesStub: func() (models.UserProvidedServiceSummary, error) {
						return srcUPSSummary, nil
					},
				}
			}
			srcSession.MockServiceSummary = func() api.ServiceSummaryRepository {
				return &FakeServiceSummaryRepository{
					GetSummariesInCurrentSpaceStub: func() (services []models.ServiceInstance, err error) {
						for _, s := range srcServices {
							services = append(services, s)
						}
						return
					},
				}
			}
			srcSession.MockServices = func() api.ServiceRepository {
				return &FakeServiceRepository{
					FindInstanceByNameStub: func(name string) (instance models.ServiceInstance, err error) {
						for _, s := range srcServices {
							if name == s.Name {
								return s, nil
							}
						}
						err = &errors.ModelNotFoundError{}
						return
					},
				}
			}
			srcSession.MockServiceKeys = func() api.ServiceKeyRepository {
				return &FakeServiceKeyRepository{
					CreateServiceKeyStub: func(serviceGUID string, keyName string, params map[string]interface{}) (err error) {
						service, exists := srcServices[serviceGUID]
						Expect(exists).Should(BeTrue())

						serviceKey := models.ServiceKey{
							Fields: models.ServiceKeyFields{
								GUID: fmt.Sprintf("svc-bnd-%d", srcServiceKeyGUIDCounter),
								Name: keyName,
							},
						}
						srcServiceKeys[serviceKey.Fields.GUID] = serviceKey
						srcServiceKeyGUIDCounter++

						service.ServiceKeys = append(service.ServiceKeys, serviceKey.Fields)
						srcServices[serviceGUID] = service
						return
					},
					GetServiceKeyStub: func(serviceGUID string, keyName string) (serviceKey models.ServiceKey, err error) {
						service, exists := srcServices[serviceGUID]
						Expect(exists).Should(BeTrue())
						for _, sk := range service.ServiceKeys {
							if keyName == sk.Name {
								serviceKey.Fields = sk
								return
							}
						}
						err = &errors.ModelNotFoundError{}
						return
					},
					DeleteServiceKeyStub: func(serviceKeyGUID string) (err error) {
						serviceKey, exists := srcServiceKeys[serviceKeyGUID]
						Expect(exists).Should(BeTrue())
						Expect(serviceKey.Fields.Name).To(Equal("__svc3_copy_for_/destTarget/destOrg/destSpace"))
						delete(srcServiceKeys, serviceKeyGUID)
						for _, s := range srcServices {
							for i, sk := range s.ServiceKeys {
								if serviceKeyGUID == sk.GUID {
									s.ServiceKeys = append(s.ServiceKeys[:i], s.ServiceKeys[i+1:]...)
									srcServices[s.GUID] = s
									return
								}
							}
						}
						return
					},
				}
			}
			destSession.MockServiceSummary = func() api.ServiceSummaryRepository {
				return &FakeServiceSummaryRepository{
					GetSummariesInCurrentSpaceStub: func() (services []models.ServiceInstance, err error) {
						for _, s := range destServices {
							services = append(services, s)
						}
						return
					},
				}
			}
			destSession.MockServices = func() api.ServiceRepository {
				return &FakeServiceRepository{
					CreateServiceInstanceStub: func(name, planGUID string, params map[string]interface{}, tags []string) (err error) {
						serviceOffering := destServicePlans[planGUID]
						for _, servicePlan := range serviceOffering.Plans {
							if planGUID == servicePlan.GUID {
								service := models.ServiceInstance{
									ServiceInstanceFields: models.ServiceInstanceFields{
										GUID:             fmt.Sprintf("svc-%d", destServiceGUIDCounter),
										Name:             name,
										ApplicationNames: []string{},
									},
									ServiceBindings: []models.ServiceBindingFields{},
									ServiceKeys:     []models.ServiceKeyFields{},
									ServicePlan:     servicePlan,
									ServiceOffering: serviceOffering.ServiceOfferingFields,
								}
								destServiceGUIDCounter++
								destServices[service.GUID] = service
								return
							}
						}
						err = fmt.Errorf("Service plan %s was not found", planGUID)
						return
					},
					FindInstanceByNameStub: func(name string) (instance models.ServiceInstance, err error) {
						for _, s := range destServices {
							if name == s.Name {
								return s, nil
							}
						}
						err = &errors.ModelNotFoundError{}
						return
					},
					FindServiceOfferingsForSpaceByLabelStub: func(spaceGUID, name string) (offering models.ServiceOfferings, err error) {
						Expect(spaceGUID).To(Equal(destSpace.GUID))
						_, expectedOffering := utils.ContainsInStrings([]string{name}, []string{"MySQL", "RabbitMQ"})
						Expect(expectedOffering).Should(BeTrue())
						for _, o := range destServiceOfferings {
							if name == o.Label {
								offering = append(offering, o)
								return
							}
						}
						err = &errors.ModelNotFoundError{}
						return
					},
					DeleteServiceStub: func(instance models.ServiceInstance) (err error) {
						delete(destServices, instance.GUID)
						return
					},
				}
			}
			destSession.MockServiceBindings = func() api.ServiceBindingRepository {
				return &FakeServiceBindingRepository{
					CreateStub: func(instanceGUID string, appGUID string, paramsMap map[string]interface{}) error {
						service := destServices[instanceGUID]
						serviceBinding := models.ServiceBindingFields{
							GUID:    fmt.Sprintf("svc-bnd-%d", destServiceBindingGUIDCounter),
							AppGUID: appGUID,
						}
						destServiceBindingGUIDCounter++
						service.ServiceBindings = append(service.ServiceBindings, serviceBinding)
						destServiceBindings[serviceBinding.GUID] = serviceBinding
						return nil
					},
					DeleteStub: func(instance models.ServiceInstance, appGUID string) (bool, error) {
						for i, sb := range instance.ServiceBindings {
							if appGUID == sb.AppGUID {
								_, exists := destServiceBindings[sb.GUID]
								Expect(exists).Should(BeTrue())
								delete(destServiceBindings, sb.GUID)
								instance.ServiceBindings = append(instance.ServiceBindings[:i], instance.ServiceBindings[i+1:]...)
								destServices[instance.GUID] = instance
								return true, nil
							}
						}
						return false, nil
					},
				}
			}
			destSession.MockServiceKeys = func() api.ServiceKeyRepository {
				return &FakeServiceKeyRepository{
					DeleteServiceKeyStub: func(serviceKeyGUID string) (err error) {
						serviceKey, exists := destServiceKeys[serviceKeyGUID]
						Expect(exists).Should(BeTrue())
						Expect(serviceKey.Fields.Name).To(Equal("some-svc-key-for-svc3"))
						delete(destServiceKeys, serviceKeyGUID)
						for _, s := range destServices {
							for i, sk := range s.ServiceKeys {
								if serviceKeyGUID == sk.GUID {
									s.ServiceKeys = append(s.ServiceKeys[:i], s.ServiceKeys[i+1:]...)
									destServices[s.GUID] = s
									return
								}
							}
						}
						return
					},
				}
			}
			destSession.MockUserProvidedServices = func() api.UserProvidedServiceInstanceRepository {
				return &FakeUserProvidedServiceInstanceRepository{
					CreateStub: func(name, drainURL string, routeServiceURL string, params map[string]interface{}) (apiErr error) {
						userProvidedService := models.UserProvidedServiceEntity{
							UserProvidedService: models.UserProvidedService{
								Name:        name,
								Credentials: params,
								SpaceGUID:   destSpace.GUID,
							},
						}
						destUPSSummary.Resources = append(destUPSSummary.Resources, userProvidedService)
						service := models.ServiceInstance{
							ServiceInstanceFields: models.ServiceInstanceFields{
								GUID:             fmt.Sprintf("svc-%d", destServiceGUIDCounter),
								Name:             name,
								ApplicationNames: []string{},
							},
							ServiceBindings: []models.ServiceBindingFields{},
							ServiceKeys:     []models.ServiceKeyFields{},
						}
						destServiceGUIDCounter++
						destServices[service.GUID] = service
						return
					},
				}
			}
			destSession.MockServicePlans = func() api.ServicePlanRepository {
				return &FakeServicePlanRepository{
					SearchStub: func(searchParameters map[string]string) (servicePlanFields []models.ServicePlanFields, err error) {
						serviceOffering, exists := destServiceOfferings[searchParameters["service_guid"]]
						Expect(exists).Should(BeTrue())
						servicePlanFields = serviceOffering.Plans
						return
					},
				}
			}
			destSession.MockApplications = func() applications.Repository {
				return &FakeApplicationsRepository{
					CreateRestageRequestStub: func(guid string) (err error) {
						_, appExists := utils.ContainsInStrings([]string{guid}, []string{"app-2000", "app-2001"})
						Expect(appExists).Should(BeTrue())
						return
					},
				}
			}
			sc, err = sm.ServicesToBeCopied([]string{"app1", "app2"}, []string{"svc1"})
			if err != nil {
				Fail(err.Error())
			}
			err = sm.DoCopy(sc, true)
			if err != nil {
				Fail(err.Error())
			}

			expectUPSExists("svc1")
			ups := expectUPSExists("ups1")
			Expect(ups.Credentials["ups1-cred1"]).To(Equal("abcd"))
			Expect(ups.Credentials["ups1-cred2"]).To(Equal("wxyz"))
			ups = expectUPSExists("ups2")
			Expect(ups.Credentials["ups2-cred1"]).To(Equal("1234"))
			Expect(ups.Credentials["ups2-cred2"]).To(Equal("5678"))

			expectServiceExists("ups1")
			expectServiceExists("ups2")
			expectServiceExists("svc1")
			svc := expectServiceExists("svc3")
			Expect(svc.ServicePlan.Name).To(Equal("Small"))
			Expect(svc.ServiceOffering.Label).To(Equal("RabbitMQ"))
		})
	})
})

func expectUPSExists(name string) (ups models.UserProvidedServiceEntity) {
	for _, ups = range destUPSSummary.Resources {
		if name == ups.Name {
			return
		}
	}
	Fail(fmt.Sprintf("Expected user provided service %s was not found", name))
	return
}

func expectServiceExists(name string) (service models.ServiceInstance) {
	for _, service = range destServices {
		if name == service.Name {
			return
		}
	}
	Fail(fmt.Sprintf("Expected service %s was not found", name))
	return
}

var srcUPSSummary = models.UserProvidedServiceSummary{
	Resources: []models.UserProvidedServiceEntity{
		models.UserProvidedServiceEntity{
			UserProvidedService: models.UserProvidedService{
				Name: "ups1",
				Credentials: map[string]interface{}{
					"ups1-cred1": "abcd",
					"ups1-cred2": "wxyz",
				},
				SpaceGUID: "space-1000",
			},
		},
		models.UserProvidedServiceEntity{
			UserProvidedService: models.UserProvidedService{
				Name: "ups2",
				Credentials: map[string]interface{}{
					"ups2-cred1": "1234",
					"ups2-cred2": "5678",
				},
				SpaceGUID: "space-1000",
			},
		},
		models.UserProvidedServiceEntity{
			UserProvidedService: models.UserProvidedService{
				Name: "ups3",
				Credentials: map[string]interface{}{
					"ups3-cred1": "qwerty",
					"ups3-cred2": "asdfgh",
				},
				SpaceGUID: "space-1001",
			},
		},
	},
}

var srcServices = map[string]models.ServiceInstance{
	"svc-1000": models.ServiceInstance{
		ServiceInstanceFields: models.ServiceInstanceFields{
			GUID:             "svc-1000",
			Name:             "svc1",
			ApplicationNames: []string{"app1", "app2", "app3"},
		},
		ServiceKeys:     []models.ServiceKeyFields{},
		ServicePlan:     models.ServicePlanFields{Name: "Large"},
		ServiceOffering: models.ServiceOfferingFields{Label: "MySQL"},
	},
	"svc-1001": models.ServiceInstance{
		ServiceInstanceFields: models.ServiceInstanceFields{
			GUID:             "svc-1001",
			Name:             "svc2",
			ApplicationNames: []string{"app3"},
		},
		ServiceKeys:     []models.ServiceKeyFields{},
		ServicePlan:     models.ServicePlanFields{Name: "Medium"},
		ServiceOffering: models.ServiceOfferingFields{Label: "Redis"},
	},
	"svc-1002": models.ServiceInstance{
		ServiceInstanceFields: models.ServiceInstanceFields{
			GUID:             "svc-1002",
			Name:             "svc3",
			ApplicationNames: []string{"app2"},
		},
		ServiceKeys: []models.ServiceKeyFields{
			srcServiceKeys["svc-key-1000"].Fields,
		},
		ServicePlan:     models.ServicePlanFields{Name: "Small"},
		ServiceOffering: models.ServiceOfferingFields{Label: "RabbitMQ"},
	},
	"svc-1003": models.ServiceInstance{
		ServiceInstanceFields: models.ServiceInstanceFields{
			GUID:             "svc-1003",
			Name:             "ups1",
			ApplicationNames: []string{"app1", "app2"},
		},
		ServiceKeys: []models.ServiceKeyFields{},
	},
	"svc-1004": models.ServiceInstance{
		ServiceInstanceFields: models.ServiceInstanceFields{
			GUID:             "svc-1004",
			Name:             "ups2",
			ApplicationNames: []string{"app1", "app3"},
		},
		ServiceKeys: []models.ServiceKeyFields{},
	},
	"svc-1005": models.ServiceInstance{
		ServiceInstanceFields: models.ServiceInstanceFields{
			GUID:             "svc-1005",
			Name:             "ups3",
			ApplicationNames: []string{"app3"},
		},
		ServiceKeys: []models.ServiceKeyFields{},
	},
}

var srcServiceKeys = map[string]models.ServiceKey{
	"svc-key-1000": models.ServiceKey{
		Fields: models.ServiceKeyFields{
			GUID: "svc-key-1000",
			Name: "__svc3_copy_for_/destTarget/destOrg/destSpace",
		},
	},
}
var srcServiceKeyGUIDCounter = 1001

var destUPSSummary = models.UserProvidedServiceSummary{
	Resources: []models.UserProvidedServiceEntity{},
}

var destServices = map[string]models.ServiceInstance{
	"svc-2000": models.ServiceInstance{
		ServiceInstanceFields: models.ServiceInstanceFields{
			GUID:             "svc-2000",
			Name:             "svc3",
			ApplicationNames: []string{"app2"},
		},
		ServiceBindings: []models.ServiceBindingFields{
			destServiceBindings["svc-bnd-2000"],
		},
		ServiceKeys: []models.ServiceKeyFields{
			destServiceKeys["svc-key-2000"].Fields,
		},
	},
	"svc-2001": models.ServiceInstance{
		ServiceInstanceFields: models.ServiceInstanceFields{
			GUID:             "svc-2001",
			Name:             "ups1",
			ApplicationNames: []string{"app1", "app2"},
		},
		ServiceBindings: []models.ServiceBindingFields{
			destServiceBindings["svc-bnd-2001"],
			destServiceBindings["svc-bnd-2002"],
		},
		ServiceKeys: []models.ServiceKeyFields{},
	},
}
var destServiceGUIDCounter = 2002

var destServiceKeys = map[string]models.ServiceKey{
	"svc-key-2000": models.ServiceKey{
		Fields: models.ServiceKeyFields{
			GUID: "svc-key-2000",
			Name: "some-svc-key-for-svc3",
		},
	},
}

var destServiceBindings = map[string]models.ServiceBindingFields{
	"svc-bnd-2000": models.ServiceBindingFields{
		GUID:    "svc-bnd-2000",
		AppGUID: "app-2001",
	},
	"svc-bnd-2001": models.ServiceBindingFields{
		GUID:    "svc-bnd-2001",
		AppGUID: "app-2000",
	},
	"svc-bnd-2002": models.ServiceBindingFields{
		GUID:    "svc-bnd-2002",
		AppGUID: "app-2001",
	},
}
var destServiceBindingGUIDCounter = 2003

var destServiceOfferings = map[string]models.ServiceOffering{
	"svc-offering-2000": models.ServiceOffering{
		ServiceOfferingFields: models.ServiceOfferingFields{
			GUID:  "svc-offering-2000",
			Label: "MySQL",
		},
		Plans: []models.ServicePlanFields{
			models.ServicePlanFields{
				GUID: "svc-offering-plan-2000",
				Name: "Small",
			},
			models.ServicePlanFields{
				GUID: "svc-offering-plan-2001",
				Name: "Medium",
			},
			models.ServicePlanFields{
				GUID: "svc-offering-plan-2002",
				Name: "Large",
			},
		},
	},
	"svc-offering-2001": models.ServiceOffering{
		ServiceOfferingFields: models.ServiceOfferingFields{
			GUID:  "svc-offering-2001",
			Label: "RabbitMQ",
		},
		Plans: []models.ServicePlanFields{
			models.ServicePlanFields{
				GUID: "svc-offering-plan-2003",
				Name: "Small",
			},
			models.ServicePlanFields{
				GUID: "svc-offering-plan-2004",
				Name: "Medium",
			},
			models.ServicePlanFields{
				GUID: "svc-offering-plan-2005",
				Name: "Large",
			},
		},
	},
}

var destServicePlans = map[string]models.ServiceOffering{
	"svc-offering-plan-2000": destServiceOfferings["svc-offering-2000"],
	"svc-offering-plan-2001": destServiceOfferings["svc-offering-2000"],
	"svc-offering-plan-2002": destServiceOfferings["svc-offering-2000"],
	"svc-offering-plan-2003": destServiceOfferings["svc-offering-2001"],
	"svc-offering-plan-2004": destServiceOfferings["svc-offering-2001"],
	"svc-offering-plan-2005": destServiceOfferings["svc-offering-2001"],
}
