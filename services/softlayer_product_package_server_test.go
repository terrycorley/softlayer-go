package services_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"os"

	slclientfakes "github.com/maximilien/softlayer-go/client/fakes"
	softlayer "github.com/maximilien/softlayer-go/softlayer"
	testhelpers "github.com/maximilien/softlayer-go/test_helpers"
)

var _ = Describe("SoftLayer_Product_Package_Server_Service", func() {

	var (
		username, apiKey string
		fakeClient       *slclientfakes.FakeSoftLayerClient
		serverService    softlayer.SoftLayer_Product_Package_Server_Service
		err              error
	)

	BeforeEach(func() {
		username = os.Getenv("SL_USERNAME")
		Expect(username).ToNot(Equal(""))

		apiKey = os.Getenv("SL_API_KEY")
		Expect(apiKey).ToNot(Equal(""))

		fakeClient = slclientfakes.NewFakeSoftLayerClient(username, apiKey)
		Expect(fakeClient).ToNot(BeNil())

		serverService, err = fakeClient.GetSoftLayer_Product_Package_Server_Service()
		Expect(err).ToNot(HaveOccurred())
		Expect(serverService).ToNot(BeNil())
	})

	Context("#GetName", func() {
		It("returns the name for the service", func() {
			name := serverService.GetName()
			Expect(name).To(Equal("SoftLayer_Product_Package_Server"))
		})
	})

	Context("#GetAllObjects", func() {
		BeforeEach(func() {
			fakeClient.FakeHttpClient.DoRawHttpRequestResponse, err = testhelpers.ReadJsonTestFixtures("services", "SoftLayer_Product_Package_Server_Service_getAllObjects.json")
			Expect(err).ToNot(HaveOccurred())
		})

		It("returns an array of datatypes.SoftLayer_Product_Package_Server", func() {
			packages, err := serverService.GetAllObjects()
			Expect(err).ToNot(HaveOccurred())
			Expect(packages).ToNot(BeNil())
		})

		Context("when HTTP client returns error codes 40x or 50x", func() {
			It("fails for error code 40x", func() {
				errorCodes := []int{400, 401, 499}
				for _, errorCode := range errorCodes {
					fakeClient.FakeHttpClient.DoRawHttpRequestInt = errorCode

					_, err := serverService.GetAllObjects()
					Expect(err).To(HaveOccurred())
				}
			})

			It("fails for error code 50x", func() {
				errorCodes := []int{500, 501, 599}
				for _, errorCode := range errorCodes {
					fakeClient.FakeHttpClient.DoRawHttpRequestInt = errorCode

					_, err := serverService.GetAllObjects()
					Expect(err).To(HaveOccurred())
				}
			})
		})
	})

	Context("#GetObject", func() {
		BeforeEach(func() {
			fakeClient.FakeHttpClient.DoRawHttpRequestResponse, err = testhelpers.ReadJsonTestFixtures("services", "SoftLayer_Product_Package_Server_Service_getObject.json")
			Expect(err).ToNot(HaveOccurred())
		})

		It("returns an array of datatypes.SoftLayer_Product_Package_Server", func() {
			p, err := serverService.GetObject(4691147)
			Expect(err).ToNot(HaveOccurred())
			Expect(p).ToNot(BeNil())
		})

		Context("when HTTP client returns error codes 40x or 50x", func() {
			It("fails for error code 40x", func() {
				errorCodes := []int{400, 401, 499}
				for _, errorCode := range errorCodes {
					fakeClient.FakeHttpClient.DoRawHttpRequestInt = errorCode

					_, err := serverService.GetObject(4691147)
					Expect(err).To(HaveOccurred())
				}
			})

			It("fails for error code 50x", func() {
				errorCodes := []int{500, 501, 599}
				for _, errorCode := range errorCodes {
					fakeClient.FakeHttpClient.DoRawHttpRequestInt = errorCode

					_, err := serverService.GetObject(4691147)
					Expect(err).To(HaveOccurred())
				}
			})
		})
	})
})
