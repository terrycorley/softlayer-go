package services_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"os"

	slclientfakes "github.com/maximilien/softlayer-go/client/fakes"
	softlayer "github.com/maximilien/softlayer-go/softlayer"
	testhelpers "github.com/maximilien/softlayer-go/test_helpers"
)

var _ = Describe("SoftLayer_Location_Service", func() {

	var (
		username, apiKey string
		fakeClient       *slclientfakes.FakeSoftLayerClient
		locationService  softlayer.SoftLayer_Location_Service
		err              error
	)

	BeforeEach(func() {
		username = os.Getenv("SL_USERNAME")
		Expect(username).ToNot(Equal(""))

		apiKey = os.Getenv("SL_API_KEY")
		Expect(apiKey).ToNot(Equal(""))

		fakeClient = slclientfakes.NewFakeSoftLayerClient(username, apiKey)
		Expect(fakeClient).ToNot(BeNil())

		locationService, err = fakeClient.GetSoftLayer_Location_Service()
		Expect(err).ToNot(HaveOccurred())
		Expect(locationService).ToNot(BeNil())
	})

	Context("#GetName", func() {
		It("returns the name for the service", func() {
			name := locationService.GetName()
			Expect(name).To(Equal("SoftLayer_Location"))
		})
	})

	Context("#GetLocation", func() {
		BeforeEach(func() {
			fakeClient.FakeHttpClient.DoRawHttpRequestResponse, err = testhelpers.ReadJsonTestFixtures("services", "SoftLayer_Location_Service_getObject.json")
			Expect(err).ToNot(HaveOccurred())
		})

		It("successfully retrieves SoftLayer_Location instance", func() {
			location, err := locationService.GetObject(449612)
			Expect(err).ToNot(HaveOccurred())
			Expect(location).ToNot(BeNil())
			Expect(location.LongName).To(Equal("Sydney 1"))
			Expect(location.Name).To(Equal("syd01"))
			Expect(location.Id).To(Equal(449612))
		})

		Context("when HTTP client returns error codes 40x or 50x", func() {
			It("fails for error code 40x", func() {
				errorCodes := []int{400, 401, 499}
				for _, errorCode := range errorCodes {
					fakeClient.FakeHttpClient.DoRawHttpRequestInt = errorCode

					_, err := locationService.GetObject(123456)
					Expect(err).To(HaveOccurred())
				}
			})

			It("fails for error code 50x", func() {
				errorCodes := []int{500, 501, 599}
				for _, errorCode := range errorCodes {
					fakeClient.FakeHttpClient.DoRawHttpRequestInt = errorCode

					_, err := locationService.GetObject(123456)
					Expect(err).To(HaveOccurred())
				}
			})
		})

	})

	Context("#GetAvailableObjectStorageDatacenters", func() {
		BeforeEach(func() {
			fakeClient.FakeHttpClient.DoRawHttpRequestResponse, err = testhelpers.ReadJsonTestFixtures("services", "SoftLayer_Location_Service_getAvailableObjectStorageDatacenters.json")
			Expect(err).ToNot(HaveOccurred())
		})

		It("returns an array of datatypes.SoftLayer_Location", func() {
			locations, err := locationService.GetAvailableObjectStorageDatacenters()
			Expect(err).ToNot(HaveOccurred())
			Expect(locations).ToNot(BeNil())
		})

		Context("when HTTP client returns error codes 40x or 50x", func() {
			It("fails for error code 40x", func() {
				errorCodes := []int{400, 401, 499}
				for _, errorCode := range errorCodes {
					fakeClient.FakeHttpClient.DoRawHttpRequestInt = errorCode

					_, err := locationService.GetAvailableObjectStorageDatacenters()
					Expect(err).To(HaveOccurred())
				}
			})

			It("fails for error code 50x", func() {
				errorCodes := []int{500, 501, 599}
				for _, errorCode := range errorCodes {
					fakeClient.FakeHttpClient.DoRawHttpRequestInt = errorCode

					_, err := locationService.GetAvailableObjectStorageDatacenters()
					Expect(err).To(HaveOccurred())
				}
			})
		})
	})

	Context("#GetDatacenters", func() {
		BeforeEach(func() {
			fakeClient.FakeHttpClient.DoRawHttpRequestResponse, err = testhelpers.ReadJsonTestFixtures("services", "SoftLayer_Location_Service_getDatacenters.json")
			Expect(err).ToNot(HaveOccurred())
		})

		It("returns an array of datatypes.SoftLayer_Location", func() {
			locations, err := locationService.GetDatacenters()
			Expect(err).ToNot(HaveOccurred())
			Expect(locations).ToNot(BeNil())
		})

		Context("when HTTP client returns error codes 40x or 50x", func() {
			It("fails for error code 40x", func() {
				errorCodes := []int{400, 401, 499}
				for _, errorCode := range errorCodes {
					fakeClient.FakeHttpClient.DoRawHttpRequestInt = errorCode

					_, err := locationService.GetDatacenters()
					Expect(err).To(HaveOccurred())
				}
			})

			It("fails for error code 50x", func() {
				errorCodes := []int{500, 501, 599}
				for _, errorCode := range errorCodes {
					fakeClient.FakeHttpClient.DoRawHttpRequestInt = errorCode

					_, err := locationService.GetDatacenters()
					Expect(err).To(HaveOccurred())
				}
			})
		})
	})

	Context("#GetLocationStatus", func() {
		BeforeEach(func() {
			fakeClient.FakeHttpClient.DoRawHttpRequestResponse, err = testhelpers.ReadJsonTestFixtures("services", "SoftLayer_Location_Service_getLocationStatus.json")
			Expect(err).ToNot(HaveOccurred())
		})

		It("retrieves instance of SoftLayer_Location_Status for SoftLayer_Location", func() {
			status, err := locationService.GetLocationStatus(123456)
			Expect(err).ToNot(HaveOccurred())
			Expect(status).ToNot(BeNil())
		})

		Context("when HTTP client returns error codes 40x or 50x", func() {
			It("fails for error code 40x", func() {
				errorCodes := []int{400, 401, 499}
				for _, errorCode := range errorCodes {
					fakeClient.FakeHttpClient.DoRawHttpRequestInt = errorCode

					_, err := locationService.GetLocationStatus(123456)
					Expect(err).To(HaveOccurred())
				}
			})

			It("fails for error code 50x", func() {
				errorCodes := []int{500, 501, 599}
				for _, errorCode := range errorCodes {
					fakeClient.FakeHttpClient.DoRawHttpRequestInt = errorCode

					_, err := locationService.GetLocationStatus(123456)
					Expect(err).To(HaveOccurred())
				}
			})
		})
	})

	Context("#GetPriceGroups", func() {
		BeforeEach(func() {
			fakeClient.FakeHttpClient.DoRawHttpRequestResponse, err = testhelpers.ReadJsonTestFixtures("services", "SoftLayer_Location_Service_getPriceGroups.json")
			Expect(err).ToNot(HaveOccurred())
		})

		It("retrieves array of SoftLayer_Location_Group for SoftLayer_Location prices", func() {
			priceGroups, err := locationService.GetPriceGroups(123456)
			Expect(err).ToNot(HaveOccurred())
			Expect(priceGroups).ToNot(BeNil())
		})

		Context("when HTTP client returns error codes 40x or 50x", func() {
			It("fails for error code 40x", func() {
				errorCodes := []int{400, 401, 499}
				for _, errorCode := range errorCodes {
					fakeClient.FakeHttpClient.DoRawHttpRequestInt = errorCode

					_, err := locationService.GetPriceGroups(123456)
					Expect(err).To(HaveOccurred())
				}
			})

			It("fails for error code 50x", func() {
				errorCodes := []int{500, 501, 599}
				for _, errorCode := range errorCodes {
					fakeClient.FakeHttpClient.DoRawHttpRequestInt = errorCode

					_, err := locationService.GetPriceGroups(123456)
					Expect(err).To(HaveOccurred())
				}
			})
		})
	})

	Context("#GetRegions", func() {
		BeforeEach(func() {
			fakeClient.FakeHttpClient.DoRawHttpRequestResponse, err = testhelpers.ReadJsonTestFixtures("services", "SoftLayer_Location_Service_getRegions.json")
			Expect(err).ToNot(HaveOccurred())
		})

		It("retrieves array of SoftLayer_Location_Region objects", func() {
			regions, err := locationService.GetRegions(123456)
			Expect(err).ToNot(HaveOccurred())
			Expect(regions).ToNot(BeNil())
		})

		Context("when HTTP client returns error codes 40x or 50x", func() {
			It("fails for error code 40x", func() {
				errorCodes := []int{400, 401, 499}
				for _, errorCode := range errorCodes {
					fakeClient.FakeHttpClient.DoRawHttpRequestInt = errorCode

					_, err := locationService.GetRegions(123456)
					Expect(err).To(HaveOccurred())
				}
			})

			It("fails for error code 50x", func() {
				errorCodes := []int{500, 501, 599}
				for _, errorCode := range errorCodes {
					fakeClient.FakeHttpClient.DoRawHttpRequestInt = errorCode

					_, err := locationService.GetRegions(123456)
					Expect(err).To(HaveOccurred())
				}
			})
		})
	})
})
