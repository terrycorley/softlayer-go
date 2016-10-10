package services_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"os"

	slclientfakes "github.com/maximilien/softlayer-go/client/fakes"
	softlayer "github.com/maximilien/softlayer-go/softlayer"
	testhelpers "github.com/maximilien/softlayer-go/test_helpers"
)

var _ = Describe("SoftLayer_Location_Group_Regional_Service", func() {

	var (
		username, apiKey string
		fakeClient       *slclientfakes.FakeSoftLayerClient
		service          softlayer.SoftLayer_Location_Group_Regional_Service
		err              error
	)

	BeforeEach(func() {
		username = os.Getenv("SL_USERNAME")
		Expect(username).ToNot(Equal(""))

		apiKey = os.Getenv("SL_API_KEY")
		Expect(apiKey).ToNot(Equal(""))

		fakeClient = slclientfakes.NewFakeSoftLayerClient(username, apiKey)
		Expect(fakeClient).ToNot(BeNil())

		service, err = fakeClient.GetSoftLayer_Location_Group_Regional_Service()
		Expect(err).ToNot(HaveOccurred())
		Expect(service).ToNot(BeNil())
	})

	Context("#GetName", func() {
		It("returns the name for the service", func() {
			name := service.GetName()
			Expect(name).To(Equal("SoftLayer_Location_Group_Regional"))
		})
	})

	Context("#GetAllObjects", func() {
		BeforeEach(func() {
			fakeClient.FakeHttpClient.DoRawHttpRequestResponse, err = testhelpers.ReadJsonTestFixtures("services", "SoftLayer_Location_Group_Regional_Service_getAllObjects.json")
			Expect(err).ToNot(HaveOccurred())
		})

		It("successfully retrieves array of SoftLayer_Location_Group instances", func() {
			o, err := service.GetAllObjects()
			Expect(err).ToNot(HaveOccurred())
			Expect(o).ToNot(BeNil())
		})

		Context("when HTTP client returns error codes 40x or 50x", func() {
			It("fails for error code 40x", func() {
				errorCodes := []int{400, 401, 499}
				for _, errorCode := range errorCodes {
					fakeClient.FakeHttpClient.DoRawHttpRequestInt = errorCode

					_, err := service.GetAllObjects()
					Expect(err).To(HaveOccurred())
				}
			})

			It("fails for error code 50x", func() {
				errorCodes := []int{500, 501, 599}
				for _, errorCode := range errorCodes {
					fakeClient.FakeHttpClient.DoRawHttpRequestInt = errorCode

					_, err := service.GetAllObjects()
					Expect(err).To(HaveOccurred())
				}
			})
		})

	})

	Context("#GetDatacenters", func() {
		BeforeEach(func() {
			fakeClient.FakeHttpClient.DoRawHttpRequestResponse, err = testhelpers.ReadJsonTestFixtures("services", "SoftLayer_Location_Group_Regional_Service_getDatacenters.json")
			Expect(err).ToNot(HaveOccurred())
		})

		It("successfully retrieves array of SoftLayer_Location instances", func() {
			d, err := service.GetDatacenters(346)
			Expect(err).ToNot(HaveOccurred())
			Expect(d).ToNot(BeNil())
		})

		Context("when HTTP client returns error codes 40x or 50x", func() {
			It("fails for error code 40x", func() {
				errorCodes := []int{400, 401, 499}
				for _, errorCode := range errorCodes {
					fakeClient.FakeHttpClient.DoRawHttpRequestInt = errorCode

					_, err := service.GetDatacenters(346)
					Expect(err).To(HaveOccurred())
				}
			})

			It("fails for error code 50x", func() {
				errorCodes := []int{500, 501, 599}
				for _, errorCode := range errorCodes {
					fakeClient.FakeHttpClient.DoRawHttpRequestInt = errorCode

					_, err := service.GetDatacenters(346)
					Expect(err).To(HaveOccurred())
				}
			})
		})

	})

	Context("#GetLocations", func() {
		BeforeEach(func() {
			fakeClient.FakeHttpClient.DoRawHttpRequestResponse, err = testhelpers.ReadJsonTestFixtures("services", "SoftLayer_Location_Group_Regional_Service_getLocations.json")
			Expect(err).ToNot(HaveOccurred())
		})

		It("successfully retrieves array of SoftLayer_Location instances", func() {
			o, err := service.GetLocations(123456)
			Expect(err).ToNot(HaveOccurred())
			Expect(o).ToNot(BeNil())
		})

		Context("when HTTP client returns error codes 40x or 50x", func() {
			It("fails for error code 40x", func() {
				errorCodes := []int{400, 401, 499}
				for _, errorCode := range errorCodes {
					fakeClient.FakeHttpClient.DoRawHttpRequestInt = errorCode

					_, err := service.GetLocations(123456)
					Expect(err).To(HaveOccurred())
				}
			})

			It("fails for error code 50x", func() {
				errorCodes := []int{500, 501, 599}
				for _, errorCode := range errorCodes {
					fakeClient.FakeHttpClient.DoRawHttpRequestInt = errorCode

					_, err := service.GetLocations(123456)
					Expect(err).To(HaveOccurred())
				}
			})
		})

	})

	Context("#GetObject", func() {
		BeforeEach(func() {
			fakeClient.FakeHttpClient.DoRawHttpRequestResponse, err = testhelpers.ReadJsonTestFixtures("services", "SoftLayer_Location_Group_Regional_Service_getObject.json")
			Expect(err).ToNot(HaveOccurred())
		})

		It("successfully retrieves instance of SoftLayer_Location_Group_Regional object", func() {
			o, err := service.GetObject(66)
			Expect(err).ToNot(HaveOccurred())
			Expect(o).ToNot(BeNil())
			Expect(o.Description).To(Equal("na-usa-east-1"))
			Expect(o.Name).To(Equal("na-usa-east-1"))
			Expect(o.Id).To(Equal(66))
		})

		Context("when HTTP client returns error codes 40x or 50x", func() {
			It("fails for error code 40x", func() {
				errorCodes := []int{400, 401, 499}
				for _, errorCode := range errorCodes {
					fakeClient.FakeHttpClient.DoRawHttpRequestInt = errorCode

					_, err := service.GetObject(66)
					Expect(err).To(HaveOccurred())
				}
			})

			It("fails for error code 50x", func() {
				errorCodes := []int{500, 501, 599}
				for _, errorCode := range errorCodes {
					fakeClient.FakeHttpClient.DoRawHttpRequestInt = errorCode

					_, err := service.GetObject(66)
					Expect(err).To(HaveOccurred())
				}
			})
		})

	})

	Context("#GetPreferredDatacenter", func() {
		BeforeEach(func() {
			fakeClient.FakeHttpClient.DoRawHttpRequestResponse, err = testhelpers.ReadJsonTestFixtures("services", "SoftLayer_Location_Group_Regional_Service_getPreferredDatacenter.json")
			Expect(err).ToNot(HaveOccurred())
		})

		It("successfully retrieves preferred datacenter object of SoftLayer_Location_Group_Regional", func() {
			datacenter, err := service.GetPreferredDatacenter(66)
			Expect(err).ToNot(HaveOccurred())
			Expect(datacenter).ToNot(BeNil())
			Expect(datacenter.LongName).To(Equal("Washington 4"))
			Expect(datacenter.Name).To(Equal("wdc04"))
			Expect(datacenter.Id).To(Equal(957095))
		})

		Context("when HTTP client returns error codes 40x or 50x", func() {
			It("fails for error code 40x", func() {
				errorCodes := []int{400, 401, 499}
				for _, errorCode := range errorCodes {
					fakeClient.FakeHttpClient.DoRawHttpRequestInt = errorCode

					_, err := service.GetPreferredDatacenter(66)
					Expect(err).To(HaveOccurred())
				}
			})

			It("fails for error code 50x", func() {
				errorCodes := []int{500, 501, 599}
				for _, errorCode := range errorCodes {
					fakeClient.FakeHttpClient.DoRawHttpRequestInt = errorCode

					_, err := service.GetPreferredDatacenter(66)
					Expect(err).To(HaveOccurred())
				}
			})
		})

	})

	Context("#GetLocationGroupType", func() {
		BeforeEach(func() {
			fakeClient.FakeHttpClient.DoRawHttpRequestResponse, err = testhelpers.ReadJsonTestFixtures("services", "SoftLayer_Location_Group_Regional_Service_getLocationGroupType.json")
			Expect(err).ToNot(HaveOccurred())
		})

		It("successfully retrieves SoftLayer_Location_Group_type for SoftLayer_Location_Group_Regional", func() {
			lgt, err := service.GetLocationGroupType(66)
			Expect(err).ToNot(HaveOccurred())
			Expect(lgt).ToNot(BeNil())
			Expect(lgt.Name).To(Equal("REGIONAL"))
		})

		Context("when HTTP client returns error codes 40x or 50x", func() {
			It("fails for error code 40x", func() {
				errorCodes := []int{400, 401, 499}
				for _, errorCode := range errorCodes {
					fakeClient.FakeHttpClient.DoRawHttpRequestInt = errorCode

					_, err := service.GetLocationGroupType(66)
					Expect(err).To(HaveOccurred())
				}
			})

			It("fails for error code 50x", func() {
				errorCodes := []int{500, 501, 599}
				for _, errorCode := range errorCodes {
					fakeClient.FakeHttpClient.DoRawHttpRequestInt = errorCode

					_, err := service.GetLocationGroupType(66)
					Expect(err).To(HaveOccurred())
				}
			})
		})

	})

})
