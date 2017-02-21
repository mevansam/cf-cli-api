package copy_test

import (
	"code.cloudfoundry.org/cli/cf/models"
	. "github.com/mevansam/cf-cli-api/cfapi/mocks"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestCmd(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Copy Managers Suite")
}

var srcSession = &MockSession{
	MockGetSessionOrg: func() models.OrganizationFields {
		return srcOrg
	},
	MockGetSessionSpace: func() models.SpaceFields {
		return srcSpace
	},
}

var destSession = &MockSession{

	MockGetSessionOrg: func() models.OrganizationFields {
		return destOrg
	},
	MockGetSessionSpace: func() models.SpaceFields {
		return destSpace
	},
}

var srcOrg = models.OrganizationFields{
	GUID: "org-1000",
	Name: "source-org",
}
var srcSpace = models.SpaceFields{
	GUID: "space-1000",
	Name: "source-space",
}
var destOrg = models.OrganizationFields{
	GUID: "org-2000",
	Name: "dest-org",
}
var destSpace = models.SpaceFields{
	GUID: "space-2000",
	Name: "dest-space",
}
