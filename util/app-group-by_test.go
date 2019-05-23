package util_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/mike-carey/cfquery/util"

	"github.com/cloudfoundry-community/go-cfclient"
)

var _ = Describe("AppGroupBy", func() {

	var (
		apps []cfclient.App
	)

	BeforeEach(func() {
		apps = []cfclient.App{
			cfclient.App{
				Guid:      "app1",
				Name:      "app-1",
				SpaceGuid: "space1",
			},
			cfclient.App{
				Guid:      "app2",
				Name:      "app-2",
				SpaceGuid: "space1",
			},
			cfclient.App{
				Guid:      "app3",
				Name:      "app-3",
				SpaceGuid: "space2",
			},
			cfclient.App{
				Guid:      "app4",
				Name:      "app-4",
				SpaceGuid: "space2",
			},
		}
	})

	It("Should group by space", func() {
		expect := map[string][]cfclient.App{
			"space1": []cfclient.App{
				apps[0],
				apps[1],
			},
			"space2": []cfclient.App{
				apps[2],
				apps[3],
			},
		}

		actual, err := AppGroupBy(apps, func(app cfclient.App) (string, error) {
			return app.SpaceGuid, nil
		})

		Expect(err).To(BeNil())
		Expect(actual).To(Equal(expect))
	})

})
